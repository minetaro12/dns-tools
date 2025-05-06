package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/minetaro12/dns-tools/internal/routes"
)

//go:embed web/dist/*
var embedFs embed.FS

func main() {
	httpListen := fmt.Sprintf(":%v", getEnv("PORT", "8080"))

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
	}))
	app.Use(logger.New())
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedFs),
		PathPrefix: "web/dist",
		Browse:     false,
	}))
	routes.SetupRoutes(app)
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
