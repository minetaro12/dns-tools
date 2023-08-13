package handlers

import (
	"net/http"

	"github.com/likexian/whois"
)

func Whois(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r, err := whois.Whois(r.FormValue("domain"))
		if err != nil {
			Error(w)
			return
		}
		w.Write([]byte(string(r)))
	default:
		Error(w)
	}
}
