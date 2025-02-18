package lib

import (
	"errors"
	"net"
	"strings"

	"github.com/miekg/dns"
	"github.com/minetaro12/dns-tools/internal/models"
)

func Resolve(req models.Lookup) (string, error) {
	// req.Typeの値によって、dns.TypeAなどを設定する
	var t uint16
	if strings.EqualFold(req.Type, "A") {
		t = dns.TypeA
	} else if strings.EqualFold(req.Type, "AAAA") {
		t = dns.TypeAAAA
	} else if strings.EqualFold(req.Type, "CNAME") {
		t = dns.TypeCNAME
	} else if strings.EqualFold(req.Type, "MX") {
		t = dns.TypeMX
	} else if strings.EqualFold(req.Type, "NS") {
		t = dns.TypeNS
	} else if strings.EqualFold(req.Type, "TXT") {
		t = dns.TypeTXT
	} else {
		return "", errors.New("unsupported record type")
	}

	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(req.Fqdn), t)

	r, _, err := c.Exchange(m, net.JoinHostPort(req.Dns, "53"))
	if err != nil {
		return "", err
	}

	return r.String(), nil
}
