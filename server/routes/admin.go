package routes

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(app *fiber.App) {
	app.Get("/api/user", middlewares.Auth, controllers.UserPing)
	app.Post("/api/user/login", controllers.UserLogin)
	app.Get("/api/user/logout", controllers.UserLogout)
}
