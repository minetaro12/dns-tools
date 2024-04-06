package main

import (
	"dns-tools/handlers"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed web/build/*
var embedFs embed.FS

func main() {
	httpListen := fmt.Sprintf(":%v", getEnv("PORT", "8000"))

	app := fiber.New()

	// app.Static("/", "./web/build")
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedFs),
		PathPrefix: "web/build",
		Browse:     false,
	}))
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
