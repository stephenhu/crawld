package main

import (
	"flag"
	"fmt"
	"log"
	
	"github.com/gomodule/redigo/redis"
)

const (
	APP_NAME				= "crawld"
	VERSION         = "0.1"
)

var (

	conf				= flag.String("conf", "config.json", "configuration file")
	redisAddr   = flag.String("redis", ":6379", "redis server")
	query    		= flag.String("query", "", "query parameter")
	
)

var cnx redis.Conn
var m = make(map[string] bool)


func appLog(msg string, fname string) {
	log.Printf("%s v%s: %s(): %s", APP_NAME, VERSION, fname, msg)
} // appLog


func initRedis() {

	c, err := redis.Dial("tcp", *redisAddr)

	if err != nil {

		appLog(err.Error(), "initRedis")
		log.Fatal("No redis connection found")
	
	}

	cnx = c

} // initRedis


func main() {

	flag.Parse()

	//log.Printf("Initiating connection to redis at address %s", *redisAddr)

	//initRedis()

	//defer cnx.Close()

	if *query == "" {
		log.Fatal("Please add location to crawl with the -page option")
	}

	p := 20

	for i := 0; i < p; i++ {

		crawler(fmt.Sprintf(SRC_JDCOM, *query, *query, i))

	}

	log.Println(m)
	log.Println(len(m))

} // main
