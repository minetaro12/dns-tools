package handlers

import "net/http"

func Error(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(string(message)))
}
