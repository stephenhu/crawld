package main

import (
	//"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func crawler(link string) {

	lnk := SanitizeURL(link)

	res, err := http.Get(lnk)

	if err != nil {
		appLog(err.Error(), "crawler")
	} else {

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			appLog(err.Error(), "crawler")
		} else {

			doc.Find(HTML_A).Each(func(index int, item *goquery.Selection) {

				l, _ := item.Attr(ATTR_HREF)

				if strings.Contains(l, LNK_JDCOM) || strings.Contains(l, LNK_JDCOM_CCC_X) {
					
					s := strings.TrimSuffix(l, "#comment")

					s = SanitizeURL(s)
					
					if !productStore.Exists(s) {

						productStore.Put(s, map[string] interface{}{
							"created": time.Now(),
						})
	
						rediss.LPush(PRODUCTQ, s)

					}

					/*
					_, ok := m[s]

					if !ok {

						m[s] = CrawldEntity{
							Referral: SanitizeURL(s),
							Created: time.Now(),
							State: STATE_OPEN,
						}

					}

					*/

				}

			})

		}
	}

} // crawler
