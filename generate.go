package main

import (
	"html/template"
	"os"

	"github.com/stephenhu/go-nude"
)


func generateHtml() {

	t := template.Must(template.New("").Parse(`<html><body>{{range .}}<img src={{.}}></img>{{end}}</body></html>`))
	
	links := []string{}

	for k, _ := range g {

		isNude, err := nude.IsURLNude("http:" + k)

		if err != nil {
			appLog(err.Error(), "generateHtml")
		} else {

			if isNude {
				links = append(links, "http:" + k)
			}

		}

	}

	f, err := os.Create("bmodel.html")

	if err != nil {
		appLog(err.Error(), "generateHtml")
	} else {

		err := t.Execute(f, links)

		if err != nil {
			appLog(err.Error(), "generateHtml")
		}
	
	}

} // generateHtml
