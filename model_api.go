package main

import (
	//"encoding/json"
	//"log"
	"net/http"
	"strings"
)


const (

	CREATE_MODEL 	= "INSERT into images(" +
		"imageurl) " +
		"VALUES(?)"

	GET_MODEL			= "SELECT id " +
		"FROM images " +
		"WHERE imageurl=?"

)


func modelExists(m string) bool {

	row := data.QueryRow(
		GET_MODEL, m,
	)

	id := 0

	err := row.Scan(&id)

	if err != nil {
		appLogError(err, "modelExists")
		return false
	} else {
		return true
	}

} // modelExists


func createModel(m string) {

	//TODO: parse tags, referral links, dimensions, meta data

	if !modelExists(m) {

		_, err := data.Exec(
			CREATE_MODEL, m,
		)
	
		if err != nil {
			appLogError(err, "createModel")
		} else {
			appLogInfo("Model added " + m)
		}
	
	}

} // createModel


func batchRatings(ratings []string) {

	for _, r := range ratings {

		s := strings.Trim(r, " ")
	
		if !modelExists(s) {

			createModel(s)

		}

	}

} // batchRatings


func modelAPIHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		imageUrl := r.FormValue("imageUrl")

		createModel(imageUrl)
		
	case http.MethodGet:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	
} // modelAPIHandler
