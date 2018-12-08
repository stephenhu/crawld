package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
		
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	_ "github.com/golang-migrate/migrate/database/sqlite3"
)

const (
	APP_NAME				= "crawld"
	VERSION         = "0.1"
)

var (

	conf					= flag.String("conf", "crawld.json", "Configuration file")
	databaseAddr  = flag.String("database", "./db/crawld.db", "Database address")
	depth       	= flag.Int("depth", 1, "Number of pages to crawl")
	query    			= flag.String("query", "", "Query parameter")
	redisAddr   	= flag.String("redis", "", "Redis server")
	serverAddr    = flag.String("server", ":8883", "Server address")
	
)

var rediss 		*redis.Client 	= nil
var data      *sql.DB 				= nil

var config 		= CrawldConfig{}


func redisStr() string {
  return fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
} // redisStr

	
func normalizeConfig() {

	if *redisAddr != "" {
		
		host, port, err := net.SplitHostPort(*redisAddr)

		if err != nil {
			appLogError(err, "normalizeConfig")
		} else {

			config.Redis.Host = host
			config.Redis.Port = port

		}

	}

	/*
	if *rabbitMQAddr != "" {
		
		host, port, err := net.SplitHostPort(*rabbitMQAddr)

		if err != nil {
			appLog(err.Error(), "normalizeConfig")
		} else {

			config.RabbitMQ.Host = host
			config.RabbitMQ.Port = port

		}

	}
	*/

	if *query != "" {
		config.Queries = append(config.Queries, *query)
	}




} // normalizeConfig


func parseConfig() {

	_, err := os.Stat(*conf)
	
	if err != nil || os.IsNotExist(err) {
		appLogError(err, "parseConfig")
	} else {

		buf, err := ioutil.ReadFile(*conf)

		if err != nil {
			appLogError(err, "parseConfig")
		} else {

			err := json.Unmarshal(buf, &config)

			if err != nil {
				appLogError(err, "parseConfig")
			}

			normalizeConfig()

		}

	}

} // parseConfig


func appLogError(err error, fname string) {
	log.Printf("[%s v%s] %s(): %s", APP_NAME, VERSION, fname, err.Error())
} // appLogError


func appLogInfo(msg string) {
	log.Printf("[%s v%s] %s", APP_NAME, VERSION, msg)
} // appLogInfo


func initRedis() {

	c := redis.NewClient(&redis.Options{
		Addr:					redisStr(),
		DialTimeout: 	10 * time.Second,
		ReadTimeout: 	30 * time.Second,
		WriteTimeout:	30 * time.Second,
		PoolSize:			10,
		PoolTimeout:	30 * time.Second,
	})

	err := c.Ping().Err()

	if err != nil {

		appLogError(err, "initRedis")
		log.Fatal("Unable to connect to Redis")
	
	}

	rediss = c

	appLogInfo(fmt.Sprintf("Redis connection established at %s", redisStr()))

} // initRedis


func initDatabase() {

	_, err := os.Stat(*databaseAddr)

	if err != nil || os.IsNotExist(err) {
		appLogError(err, "connectDatabase")
		log.Fatal("Database not found, please initialize database.")
	} else {

		db, err := sql.Open("sqlite3", *databaseAddr)

		if err != nil {
			appLogError(err, "connectDatabase")
		}

		data = db

		appLogInfo(fmt.Sprintf("Database connection established at %s", *databaseAddr))

	}

} // initDatabase


func initRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/crawlers", crawlerAPIHandler)
	router.HandleFunc("/api/images", imageAPIHandler)
	router.HandleFunc("/api/models", modelAPIHandler)
	
	return router

} // initRoutes


func main() {
	
	flag.Parse()

	parseConfig()

	initRedis()

	defer rediss.Close()

	initDatabase()

	go productParser()

	appLogInfo(fmt.Sprintf("Listening on port %s", *serverAddr))
	log.Fatal(http.ListenAndServe(*serverAddr, initRoutes()))
	
} // main
