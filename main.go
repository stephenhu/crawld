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

	conf				= flag.String("conf", "config.json", "Configuration file")
	depth       = flag.Int("depth", 1, "Number of pages to crawl")
	mem         = flag.Bool("mem", false, "Utilize memory instead of Redis to store URLs")
	query    		= flag.String("query", "", "Query parameter")
	redisAddr   = flag.String("redis", ":6379", "Redis server")
	
)

var cnx redis.Conn

var m map[string] CrawldPage
var g map[string] CrawldImage


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

	if *mem {

		m = make(map[string] CrawldPage)
		g = make(map[string] CrawldImage)

	} else {

		appLog(fmt.Sprintf(
			"Initiating connection to Redis at %s", *redisAddr),
			"main")

		initRedis()
	
		defer cnx.Close()
	
	}

	if *query == "" {
		log.Fatal("Please add a query for crawld with the -query option")
	}

	for i := 1; i <= *depth; i++ {

		crawler(fmt.Sprintf(SRC_JDCOM, *query, *query, i))

	}

	//log.Println(m)
	log.Println(len(m))
	log.Println(*depth)

	for _, k := range m {
		getImageList(k.Referral)
	}

	log.Println(len(g))
	
	generateHtml()

} // main
