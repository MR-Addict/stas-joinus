package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func AppRoute(app *fiber.App) {
	app.Get("/api/app/config", controllers.AppConfig)
}
