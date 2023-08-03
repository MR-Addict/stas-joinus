package main

import (
	"embed"
	"server/configs"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

//go:embed all:build
var staticFiles embed.FS

func main() {
	configs.ConnectDB()
	defer configs.DisconnectDB()

	app := fiber.New()

	routes.AdminRoute(app)
	routes.ApplicantsRoute(app)
	routes.ServeStatic(app, staticFiles)

	app.Listen(":4000")
}