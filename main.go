package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
	
	//"github.com/gomodule/redigo/redis"
	"github.com/go-redis/redis"
	//"github.com/streadway/amqp"
)

const (
	APP_NAME				= "crawld"
	VERSION         = "0.1"
)

var (

	conf					= flag.String("conf", "crawld.json", "Configuration file")
	depth       	= flag.Int("depth", 1, "Number of pages to crawl")
	mem         	= flag.Bool("mem", false, "Utilize memory instead of Redis to store URLs")
	query    			= flag.String("query", "", "Query parameter")
	redisAddr   	= flag.String("redis", "", "Redis server")
	//rabbitMQAddr  = flag.String("rabbitmq", "", "Rabbitmq server")
	
)

var rediss 				*redis.Client
//var rabbits 			*amqp.Connection
//var productsch    *amqp.Channel
//var productsq     amqp.Queue

var config = CrawldConfig{}

var productStore 		Store
var imageStore 			Store


func redisStr() string {
  return fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
} // redisStr

/*
func rabbitMQStr() string {
  return fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.RabbitMQ.User,
		config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port,
	  config.RabbitMQ.Meta["vhost"])
} // rabbitMQStr
*/
	
func normalizeConfig() {

	if *redisAddr != "" {
		
		host, port, err := net.SplitHostPort(*redisAddr)

		if err != nil {
			appLog(err.Error(), "normalizeConfig")
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
		appLog(err.Error(), "parseConfig")
	} else {

		buf, err := ioutil.ReadFile(*conf)

		if err != nil {
			appLog(err.Error(), "parseConfig")
		} else {

			err := json.Unmarshal(buf, &config)

			if err != nil {
				appLog(err.Error(), "parseConfig")
			}

			normalizeConfig()

		}

	}

} // parseConfig


func appLog(msg string, fname string) {
	log.Printf("[%s v%s] %s(): %s", APP_NAME, VERSION, fname, msg)
} // appLog

/*
func initRabbitMQ() {

	c, err := amqp.Dial(rabbitMQStr())

	if err != nil {
		appLog(err.Error(), "initRabbitMQ")
		log.Println(rabbitMQStr())
		log.Fatal("Unable to connect to RabbitMQ")
	} else {

		ch, err := c.Channel()

		if err != nil {
			appLog(err.Error(), "initRabbitMQ")
			log.Fatal("Unable to connect to RabbitMQ channel")
		} else {

			q, err := ch.QueueDeclare(
				PRODUCTS, true, false, false, false, nil,
			)

			if err != nil {
				appLog(err.Error(), "initRabbitMQ")
				log.Fatal("Unable to declare a queue in RabbitMQ")
			}

			productsq 	= q
			productsch 	= ch

		}

	}

	rabbits 		= c

} // initRabbitMQ
*/

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

		appLog(err.Error(), "initRedis")
		log.Fatal("Unable to connect to Redis")
	
	}

	rediss = c

} // initRedis


func main() {

	flag.Parse()

	parseConfig()

	if *mem {

		productStore 		= MemStore{
			Entities: make(map[string] map[string] interface{}),
		}

		imageStore			= MemStore{
			Entities: make(map[string] map[string] interface{}),
		}

	} else {

		appLog(fmt.Sprintf("Initiating connection to Redis at %s",
		  redisStr()), "main")

		initRedis()

		defer rediss.Close()
	
		productStore 		= RedisStore{}
		imageStore			= RedisStore{}

	
	}

	//initRabbitMQ()

	//defer rabbits.Close()

	//defer productsch.Close()

	go productParser()
	
	if len(config.Queries) == 0 {
		log.Fatal("Please add a query for crawld with the -query option")
	}

	for i := 0; i < len(config.Queries); i++ {

		for j := 1; j <= *depth; j++ {
			crawler(fmt.Sprintf(SRC_JDCOM, config.Queries[i], config.Queries[i], j))
		}

	}

	all, err := imageStore.GetAll()

	if err != nil {
		appLog(err.Error(), "main")
	} else {
		//log.Println(all)
		log.Println(len(all))	
	}
	

	/*
	for _, k := range m {
		getImageList(k.Referral)
	}

	//log.Println(len(g))
	
	generateHtml()
	*/
	
} // main
