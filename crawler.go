package main

import (
	//"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	HTML_A          = "a"
	HTML_IMG        = "img"
	ATTR_HREF 			= "href"
	ATTR_SRC        = "src"
)

func crawler(link string) {

	res, err := http.Get(link)

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
					//log.Println(l)
					
					s := strings.TrimSuffix(l, "#comment")

					_, ok := m[s]

					if !ok {
						m[s] = false
					}

				}

			})

		}
	}

} // crawler
