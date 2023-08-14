package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/likexian/whois"
)

type whoisReq struct {
	Domain string `json:"domain"`
}

func Whois(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	// w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

	switch r.Method {
	case "POST":
		if !isValidReq(r) {
			http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
			return
		}

		var t whoisReq
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		r, err := whois.Whois(t.Domain)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(string(r)))
	// case "OPTIONS":
	// 	w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}
}
