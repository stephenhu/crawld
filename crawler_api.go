package main

import (
	"fmt"
	"net/http"
)


func crawlerAPIHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		q := r.FormValue("q")

		if q == "" {
			w.WriteHeader(http.StatusNotFound)
		} else {

			for i := 1; i <= *depth; i++ {
				crawler(fmt.Sprintf(SRC_JDCOM, q, q, i))
			}

		}


	case http.MethodGet:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	
} // crawlerAPIHandler
