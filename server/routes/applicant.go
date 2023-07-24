package routes

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ApplicantsRoute(app fiber.Router) {
	app.Post("/applicant", controllers.CreateApplicant)
	app.Get("/applicants", middlewares.Auth, controllers.GetAllApplicants)
}
