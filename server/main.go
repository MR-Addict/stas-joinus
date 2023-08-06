package main

import (
	"embed"
	"server/configs"
	"server/routes"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//go:embed all:build
var staticFiles embed.FS

func main() {
	configs.LoadConfig()
	configs.ConnectDb()

	app := fiber.New()

	routes.AdminRoute(app)
	routes.ApplicantsRoute(app)
	routes.ServeStatic(app, staticFiles)

	app.Listen(":" + strconv.Itoa(configs.Config.Port))
}
