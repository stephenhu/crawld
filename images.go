package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	//"github.com/gomodule/redigo/redis"
	//"github.com/go-redis/redis"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)


func storeImageList(link string) {

	l := SanitizeURL(link)

	res, err := http.Get(l)

	if err != nil {
		appLog(err.Error(), "storeImageList")
	} else {

		defer res.Body.Close()

		var j map[string] interface{}

		err := json.NewDecoder(res.Body).Decode(&j)

		if err != nil {
			appLog(err.Error(), "storeImageList")
		} else {

			doc := html.NewTokenizer(strings.NewReader(j["content"].(string)))

			for {

				e := doc.Next()

				if e == html.ErrorToken {
					break
				} else {

					name, _ := doc.TagName()
					 
					if string(name) == HTML_IMG {

						k, v, _ := doc.TagAttr()

						if string(k) == ATTR_DATA_LAZYLOAD {
							
							/*
							g[string(v)] = CrawldImage {
								Referral: l,
								Created: time.Now(),
								Valid: true,
							}
							*/

							
							imageStore.Put(SanitizeURL(string(v)),map[string] interface{}{
								"referral": l,
								"created": time.Now(),
							})


						}

					}

				}

			}

		}

	}

} // storeImageList


func getImageList(link string) {

	re := regexp.MustCompile(
		`cd.jd.com/description/channel\?skuId=[0-9]+&mainSkuId=[0-9]+&cdn=[0-9]+`)

	res, err := http.Get(link)

	if err != nil {
		appLog(err.Error(), "getImageList")
	} else {

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			appLog(err.Error(), "getImageList")
		} else {

			doc.Find(HTML_SCRIPT).Each(func(index int, item *goquery.Selection) {

				t := item.Text()

				if strings.Contains(t, JS_DESC) {

					match := re.FindString(t)

					storeImageList(match)

				}

			})

		}

	}

} // getImageList


func productParser() {

	r, err := rediss.BLPop(0, PRODUCTQ).Result()

	if err != nil {
		appLog(err.Error(), "productParser")
	} else {
		log.Println(r)
		getImageList(r[1])
	}

} // productParser
