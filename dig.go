package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/lixiangzhong/dnsutil"
	"github.com/miekg/dns"
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

func execDig(digDns, digType, domain string) (*dns.Msg, error) {
	var d dnsutil.Dig
	if digDns == "" {
		d.SetDNS("8.8.8.8")
	} else {
		d.SetDNS(digDns)
	}

	var r *dns.Msg
	var err error
	switch {
	case matchString("A", digType):
		r, err = d.GetMsg(dns.TypeA, domain)
	case matchString("AAAA", digType):
		r, err = d.GetMsg(dns.TypeAAAA, domain)
	case matchString("CNAME", digType):
		r, err = d.GetMsg(dns.TypeCNAME, domain)
	case matchString("MX", digType):
		r, err = d.GetMsg(dns.TypeMX, domain)
	case matchString("TXT", digType):
		r, err = d.GetMsg(dns.TypeTXT, domain)
	case matchString("NS", digType):
		r, err = d.GetMsg(dns.TypeNS, domain)
	default:
		r, err = d.GetMsg(dns.TypeANY, domain)
	}

	if err != nil {
		return nil, err
	}
	return r, nil
}

func matchString(str string, target string) bool {
	r := regexp.MustCompile(fmt.Sprintf(`^(?i)%v$`, str))
	return r.MatchString(target)
}
