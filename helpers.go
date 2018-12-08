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

		if !strings.Contains(s, HTTP) && !strings.Contains(s, HTTPS) {
			return HTTP + s
		}

		return s

	}

} // SanitizeURL
