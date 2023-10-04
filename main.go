package main

import (
	"dns-tools/handlers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	httpListen := fmt.Sprintf(":%v", getEnv("PORT", "8000"))

	gin.SetMode("release")
	r := gin.Default()
	r.Static("/", "./web/build")
	r.POST("/whois", handlers.Whois)
	r.POST("/lookup", handlers.Lookup)

	log.Println("Server Listening on", httpListen)
	err := r.Run(httpListen)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if val, isFound := os.LookupEnv(key); isFound {
		return val
	}
	return fallback
}
