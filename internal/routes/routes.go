package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minetaro12/dns-tools/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/whois", handlers.PostWhois)
	app.Post("/api/lookup", handlers.PostLookup)
}
