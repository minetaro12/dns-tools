package main

import (
	"dns-tools/handlers"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed web/build/*
var static embed.FS

func main() {
	public, err := fs.Sub(static, "web/build")
	if err != nil {
		log.Fatal(err)
	}

	httpListen := fmt.Sprintf(":%v", getEnv("PORT", "8000"))

	http.Handle("/", http.FileServer(http.FS(public)))

	http.HandleFunc("/whois", handlers.Whois)
	http.HandleFunc("/lookup", handlers.Lookup)

	log.Println("Server Listening on", httpListen)
	log.Fatal(http.ListenAndServe(httpListen, logRequest(http.DefaultServeMux)))
}

func getEnv(key, fallback string) string {
	if val, isFound := os.LookupEnv(key); isFound {
		return val
	}
	return fallback
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
