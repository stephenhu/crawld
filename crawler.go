package main

import (
	//"log"
	"net/http"
	"strings"
	//"time"

	"github.com/PuerkitoBio/goquery"
)

func crawler(link string) {

	lnk := SanitizeURL(link)

	appLogInfo(lnk)
	
	res, err := http.Get(lnk)

	if err != nil {
		appLogError(err, "crawler")
	} else {

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			appLogError(err, "crawler")
		} else {

			doc.Find(HTML_A).Each(func(index int, item *goquery.Selection) {

				l, _ := item.Attr(ATTR_HREF)

				//TODO: filter all
				if strings.Contains(l, LNK_JDCOM) || strings.Contains(l, LNK_JDCOM_CCC_X) {
					
					s := strings.TrimSuffix(l, "#comment")

					s = SanitizeURL(s)

					r, err := rediss.SAdd(PRODUCTS, s).Result()

					if err != nil {
						appLogError(err, "crawler")
					} else {

						if r > 0 {

							err = rediss.LPush(PRODUCTSQ, s).Err()
	
							if err != nil {
								appLogError(err, "crawler")
							}
		
						}
	
					}

				}

			})

		}
	}

} // crawler
