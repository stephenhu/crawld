package main

import (
	"html/template"
	"os"

	"gocv.io/x/gocv"
)


func generateHtml() {

	t := template.Must(template.New("").Parse(`<html><body>{{range .}}<img src={{.}}></img>{{end}}</body></html>`))
	
	links := []string{}

	for k, _ := range g {

		img := gocv.IMRead("http:" + k, gocv.IMReadColor)

		if img.Empty() {
			appLog(err.Error(), "generateHtml")
		} else {

			hsv := gocv.NewMat()

			gocv.CvtColor(img, &hsv, gocv.ColorBGRToHSV)

			lower := gocv.NewMatFromScalar(gocv.NewScalar.(0.0, 48.0, 80.0, 0.0), gocv.MatTypeCV8UC3)
			upper := gocv.NewMatFromScalar(gocv.NewScalar.(20.0, 255.0, 255.0, 0.0), gocv.MatTypeCV8UC3)
			
			mask := gocv.NewMat()

			skin := gocv.InRange(hsv, lower, upper, &mask)

			log.Println(skin)

			//links = append(links, "http:" + k)


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
