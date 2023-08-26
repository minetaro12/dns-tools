package handlers

import (
	"dns-tools/lib"
	"net/http"
	"os/exec"
)

type nslookupReq struct {
	Fqdn string `json:"fqdn"`
	Dns  string `json:"dns"`
	Type string `json:"type"`
}

func Nslookup(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("cache-control", "public, max-age=0, must-revalidate")

	switch r.Method {
	case "POST":
		if !lib.IsValidRequest(r) {
			http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
			return
		}

		var t nslookupReq
		if err := lib.JsonDecode(r.Body, &t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r, err := execNslookup(t); err != nil {
			http.Error(w, r, http.StatusInternalServerError)
		} else {
			w.Write([]byte(string(r)))
		}

	default:
		// POSTメソッド以外は受け付けない
		w.Header().Add("allow", "POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
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
		return string(r), err
	}
	return string(r), nil
}
