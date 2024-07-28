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
	app := fiber.New()

	configs.LoadConfig()
	configs.ConnectDb()
	configs.SetupCors(app)

	routes.AdminRoute(app)
	routes.ApplicantsRoute(app)
	routes.ServeStatic(app, staticFiles)

	app.Listen(":" + configs.Config.Port)
}
