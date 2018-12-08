package main

import (
	"encoding/json"
	"net/http"
)


func imageAPIHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		r, err := rediss.LPop(IMAGESQ).Result()

		if err != nil {
			appLogError(err, "imageAPIHandler.Get")
			w.WriteHeader(http.StatusNotFound)
		} else {

			j, err := json.Marshal(map[string] string{
				"imageUrl": r,
			})

			if err != nil {
				appLogError(err, "iamgeAPIHandler.Get")
			} else {
				w.Write(j)
			}

		}

	case http.MethodPost:		
	case http.MethodPut:
	case http.MethodDelete:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	
} // imageAPIHandler
