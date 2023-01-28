package main

import (
	"net/http"

	"github.com/likexian/whois"
)

func whoisHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r, err := whois.Whois(r.FormValue("domain"))
		if err != nil {
			errorResponse(w)
			return
		}
		w.Write([]byte(string(r)))
	default:
		errorResponse(w)
	}
}
