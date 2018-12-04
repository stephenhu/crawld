package main

import (
	"strings"
)


const (

	DOUBLE_SLASH		= "//"
	HTTP						= "http://"
	HTTPS						= "https://"

)

func SanitizeURL(s string) string {

	if strings.HasPrefix(s, DOUBLE_SLASH) {
		return strings.Replace(s, DOUBLE_SLASH, HTTP, 1)
	} else {
		return s
	}

} // SanitizeURL
