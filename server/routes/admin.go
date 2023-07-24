package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(app fiber.Router) {
	app.Post("/login", controllers.Login)
}
