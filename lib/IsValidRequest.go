package lib

import (
	"net/http"
	"strings"
)

func IsValidRequest(r *http.Request) bool {
	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		return true
	} else {
		return false
	}
}
