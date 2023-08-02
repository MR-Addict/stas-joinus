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

	routes.AdminRoute(app)
	routes.ApplicantsRoute(app)

	app.Static("/", "./public", fiber.Static{MaxAge: 60 * 60 * 24})
	app.Listen(":4000")
}
