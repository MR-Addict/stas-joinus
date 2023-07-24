package routes

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(app fiber.Router) {
	app.Get("/user", middlewares.Auth, controllers.UserPing)
	app.Post("/user/login", controllers.UserLogin)
}
