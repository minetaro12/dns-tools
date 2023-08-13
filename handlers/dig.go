package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
)

func Dig(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r, err := execDig(r.FormValue("dns"), r.FormValue("type"), r.FormValue("fqdn"))
		if err != nil {
			Error(w)
			return
		}
		fmt.Fprint(w, r)
	default:
		Error(w)
	}
}

func execDig(reqDns, reqType, fqdn string) (string, error) {
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

	r, err := exec.Command("dig", fqdn, "@"+digDns, digType).Output()
	if err != nil {
		return "", err
	}
	return string(r), nil
}
