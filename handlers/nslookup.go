package handlers

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

type nslookupReq struct {
	Fqdn string `json:"fqdn"`
	Dns  string `json:"dns"`
	Type string `json:"type"`
}

func Nslookup(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Methods", "POST")

	switch r.Method {
	case "POST":
		if !isValidReq(r) {
			http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
			return
		}

		var t nslookupReq
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		if r, err := execNslookup(t); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			w.Write([]byte(string(r)))
		}

	default:
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}
}

func execNslookup(t nslookupReq) (string, error) {
	var nslDns, nslType string
	if t.Dns == "" {
		nslDns = "8.8.8.8"
	} else {
		nslDns = t.Dns
	}

	if t.Type == "" {
		nslType = "A"
	} else {
		nslType = t.Type
	}

	r, err := exec.Command("nslookup", "-type="+nslType, t.Fqdn, nslDns).Output()
	if err != nil {
		return "", err
	}
	return string(r), nil
}
