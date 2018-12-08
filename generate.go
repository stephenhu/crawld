package main

import (
    //"bytes"
    //"html/template"
    //"image"
    //"log"
    //"net/http"
    //"os"
    //"strings"

    //"gocv.io/x/gocv"
)


func generateHtml() {
/*
    t := template.Must(template.New("").Parse(`<html><body>{{range .}}<img src={{.}}></img>{{end}}</body></html>`))
    
    links := []string{}

    
    all, err := imageStore.GetAll()

    if err != nil {
        appLogError(err, "generateHtml")
    } else {

        for k, _ := range all {

            res, err := http.Get(k)
    
            if err != nil {
              appLogError(err, "generateHtml")
            } else {
    
                buf := new(bytes.Buffer)
    
                buf.ReadFrom(res.Body)
    
                defer res.Body.Close()
    
                img, err := gocv.IMDecode(buf.Bytes(), gocv.IMReadColor)
    
                if err != nil {
                    appLogError(err, "generateHtml")
                } else if img.Empty() {
                    //appLogInfo("Unable to get image, skipping", "generateHtml") 
                } else {
    
                    hsv := gocv.NewMat()
    
                    gocv.CvtColor(img, &hsv, gocv.ColorBGRToHSV)
    
                    lower := gocv.NewMatFromScalar(gocv.NewScalar(0.0, 48.0, 80.0, 0.0), gocv.MatTypeCV8UC3)
                    upper := gocv.NewMatFromScalar(gocv.NewScalar(20.0, 255.0, 255.0, 0.0), gocv.MatTypeCV8UC3)
    
                    mask := gocv.NewMat()
    
                    gocv.InRange(hsv, lower, upper, &mask)
    
                    kernel := gocv.GetStructuringElement(gocv.MorphEllipse, image.Point{11, 11})
    
                    gocv.Erode(mask, &mask, kernel)
    
                    gocv.Dilate(mask, &mask, kernel)
    
                    blur := gocv.NewMat()
    
                    defer blur.Close()
    
                    gocv.GaussianBlur(mask, &blur, image.Pt(3, 3),
                      0, 0, gocv.BorderDefault)
    
                    count := gocv.CountNonZero(blur)
    
                    log.Println(k)
                    log.Println(count)
                    log.Println(mask.Total())
    
    
                    if count > 0 {
                        links = append(links, u)
                     }
    
    
                    links = append(links, k)
    
            }

            f, err := os.Create("bmodel.html")

            if err != nil {
                appLogError(err, "generateHtml")
            } else {

                err := t.Execute(f, links)

                if err != nil {
                    appLogError(err, "generateHtml")
                }
            
            }
        }
            
    }

    //}
    
*/
} // generateHtml
