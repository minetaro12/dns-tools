package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func digHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r, err := execDig(r.FormValue("dns"), r.FormValue("type"), r.FormValue("domain"))
		if err != nil {
			errorResponse(w)
			return
		}
		fmt.Fprint(w, r)
	default:
		errorResponse(w)
	}
}

func execDig(reqDns, reqType, domain string) (string, error) {
	var digDns, digType string
	if reqDns == "" {
		digDns = "8.8.8.8"
	} else {
		digDns = reqDns
	}

	if reqType == "" {
		digType = "A"
	} else {
		digType = reqType
	}

	r, err := exec.Command("dig", domain, "@"+digDns, digType).Output()
	if err != nil {
		return "", err
	}
	return string(r), nil
}
