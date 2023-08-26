package lib

import (
	"errors"
	"net"
	"strings"

	"github.com/miekg/dns"
)

func Resolve(fqdn, dnsAddr, recordType string) (string, error) {
	// recordTypeの値によって、dns.TypeAなどを設定する
	var t uint16
	if strings.EqualFold(recordType, "A") {
		t = dns.TypeA
	} else if strings.EqualFold(recordType, "AAAA") {
		t = dns.TypeAAAA
	} else if strings.EqualFold(recordType, "CNAME") {
		t = dns.TypeCNAME
	} else if strings.EqualFold(recordType, "MX") {
		t = dns.TypeMX
	} else if strings.EqualFold(recordType, "NS") {
		t = dns.TypeNS
	} else if strings.EqualFold(recordType, "TXT") {
		t = dns.TypeTXT
	} else {
		return "", errors.New("unsupported record type")
	}

	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(fqdn), t)

	r, _, err := c.Exchange(m, net.JoinHostPort(dnsAddr, "53"))
	if err != nil {
		return "", err
	}

	return r.String(), nil
}
