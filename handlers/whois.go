package handlers

import (
	"net/http"

	"dns-tools/lib"

	"github.com/likexian/whois"
)

type whoisReq struct {
	Domain string `json:"domain"`
}

func Whois(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("cache-control", "public, max-age=0, must-revalidate")

	switch r.Method {
	case "POST":
		if !lib.IsValidRequest(r) {
			http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
			return
		}

		var t whoisReq
		if err := lib.JsonDecode(r.Body, &t); err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		r, err := whois.Whois(t.Domain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(string(r)))
	default:
		// POSTメソッド以外は受け付けない
		w.Header().Add("allow", "POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
