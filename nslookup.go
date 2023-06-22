package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func nslookupHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r, err := execNslookup(r.FormValue("dns"), r.FormValue("type"), r.FormValue("fqdn"))
		if err != nil {
			errorResponse(w)
			return
		}
		fmt.Fprint(w, r)
	default:
		errorResponse(w)
	}
}

func execNslookup(reqDns, reqType, fqdn string) (string, error) {
	var nslDns, nslType string
	if reqDns == "" {
		nslDns = "8.8.8.8"
	} else {
		nslDns = reqDns
	}

	if reqType == "" {
		nslType = "A"
	} else {
		nslType = reqType
	}

	r, err := exec.Command("nslookup", "-type="+nslType, fqdn, nslDns).Output()
	if err != nil {
		return "", err
	}
	return string(r), nil
}
