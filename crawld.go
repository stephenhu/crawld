package main

import (
	"time"

)

const (

	SRC_JDCOM = "http://search.jd.com/Search?keyword=%s&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&wq=%s&page=%d&s=1&click=0"

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

type CrawldPage struct {
	Referral			string								`json:"referral"`
	Meta          map[string]string			`json:"meta"`
	Created     	time.Time      				`json:"created"`
	State         int                   `json:"state"`
	Tags          []string              `json:"tags"`
}

type CrawldImage struct {
	Referral			string								`json:"referral"`
	Meta          map[string]string			`json:"meta"`
	Created     	time.Time      				`json:"created"`
	Valid         bool                  `json:"valid"`
	Tags          []string              `json:"tags"`
	Rating        int                   `json:"rating"`
}
