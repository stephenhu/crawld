package main

import (
	"time"

)

const (

	SRC_JDCOM = "http://search.jd.com/Search?keyword=%s&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&wq=%s&page=%d&s=1&click=0"

)

const (

	PRODUCTS        = "crawld.products"
	IMAGES          = "crawld.images"
	PRODUCTQ        = "queue.product"
	IMAGEQ          = "queue.image"

)

const (

	FIELD_CREATED				= "field.created"
	FIELD_REFERRAL			= "field.referral"
	FIELD_META					= "field.meta"
	FIELD_RATING				= "field.rating"
	FIELD_TAGS					= "field.tags"

)

const (

	LNK_JDCOM 				= "item.jd.com"
	LNK_JDCOM_CCC_X 	= "ccc-x.jd.com"

)

const (

	STATE_OPEN						= 0
	STATE_CLOSED					= 1

)

const (

	HTML_A          = "a"
	HTML_IMG        = "img"
	HTML_SCRIPT     = "script"

)

const (

	ATTR_HREF 						= "href"
	ATTR_SRC        			= "src"
	ATTR_DATA_LAZYLOAD		= "data-lazyload"

)

const (

	JS_DESC		= "cd.jd.com/description"

)


const (
	REDIS_LPUSH					= "LPUSH"
	REDIS_BLPOP         = "BLPOP"
	REDIS_HMSET         = "HMSET"
	REDIS_HMGET         = "HMGET"
)


type ServiceConfig struct {
	Host				string							`json:"host"`
	Port				string							`json:"port"`
	User				string							`json:"user"`
	Password		string							`json:"password"`
	Meta        map[string] string	`json:"meta"`
}

type CrawldConfig struct {
	Depth			int								`json:"depth"`
	Queries   []string          `json:"queries"`
	Redis     ServiceConfig			`json:"redis"`
	RabbitMQ  ServiceConfig			`json:"rabbitmq"`
}

type CrawldEntity struct {
	Referral			string								`json:"referral"`
	Meta          map[string]string			`json:"meta"`
	Created     	time.Time      				`json:"created"`
	State         int                   `json:"state"`
	Valid         bool                  `json:"valid"`
	Tags          []string              `json:"tags"`
	Rating        int                		`json:"rating"`
}
