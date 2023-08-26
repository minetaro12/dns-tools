package handlers

import (
	"dns-tools/lib"
	"net/http"
)

type LookupReq struct {
	Fqdn string `json:"fqdn"`
	Dns  string `json:"dns"`
	Type string `json:"type"`
}

func Lookup(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("cache-control", "public, max-age=0, must-revalidate")

	switch r.Method {
	case "POST":
		if !lib.IsValidRequest(r) {
			http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
			return
		}

		var t LookupReq
		if err := lib.JsonDecode(r.Body, &t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r, err := execNslookup(t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write([]byte(string(r)))
		}

	default:
		// POSTメソッド以外は受け付けない
		w.Header().Add("allow", "POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func execNslookup(t LookupReq) (string, error) {
	// 空ならば8.8.8.8を設定する
	if t.Dns == "" {
		t.Dns = "8.8.8.8"
	}

	// 空ならばAレコードに設定する
	if t.Type == "" {
		t.Type = "A"
	}

	r, err := lib.Resolve(t.Fqdn, t.Dns, t.Type)
	if err != nil {
		return "", err
	}
	return r, nil
}
