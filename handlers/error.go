package handlers

import "net/http"

func Error(w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Write([]byte(string("Invalid Request")))
}
