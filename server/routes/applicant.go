package routes

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ApplicantsRoute(app fiber.Router) {
	app.Post("/api/applicant", middlewares.Limitter, controllers.ApplicantCreate)
	app.Get("/api/applicants", middlewares.Auth, controllers.ApplicantGetAll)
}
