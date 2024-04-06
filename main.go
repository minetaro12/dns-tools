package main

import (
	"dns-tools/handlers"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	httpListen := fmt.Sprintf(":%v", getEnv("PORT", "8000"))

	app := fiber.New()

	app.Static("/", "./web/build")
	app.Post("/whois", handlers.Whois)
	app.Post("/lookup", handlers.Lookup)

	err := app.Listen(httpListen)
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
