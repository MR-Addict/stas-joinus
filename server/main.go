package main

import (
	"server/configs"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.ConnectDB()
	defer configs.DisconnectDB()

	app := fiber.New()
	api := app.Group("/api")

	routes.AdminRoute(api)
	routes.ApplicantsRoute(api)

	app.Static("/", "./public")
	app.Listen(":4000")
}
