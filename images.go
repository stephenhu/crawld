package main

import (
	"encoding/json"
	//"log"
	"net/http"
	"regexp"
	"strings"
	//"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)


func storeImageList(link string) {

	l := SanitizeURL(link)

	res, err := http.Get(l)

	if err != nil {
		appLogError(err, "storeImageList")
	} else {

		defer res.Body.Close()

		var j map[string] interface{}

		err := json.NewDecoder(res.Body).Decode(&j)

		if err != nil {
			appLogError(err, "storeImageList")
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

							cleanUrl := SanitizeURL(string(v))

							r, err := rediss.SAdd(IMAGES, cleanUrl).Result()

							if err != nil {
								appLogError(err, "storeImageList")
							} else {

								if r > 0 {

									err := rediss.LPush(IMAGESQ, cleanUrl).Err()

									if err != nil {
										appLogError(err, "storeImageList")
									}

								}

							}


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
		appLogError(err, "getImageList")
	} else {

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			appLogError(err, "getImageList")
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

	for {

		r, err := rediss.LPop(PRODUCTSQ).Result()
	
		if err != nil {
			//appLogError(err, "productParser")
		} else {
			getImageList(r)
		}
	
	}

} // productParser
