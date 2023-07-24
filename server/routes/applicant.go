package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func ApplicantsRoute(app *fiber.App) {
	app.Post("/api/applicant", controllers.CreateApplicant)
	app.Get("/api/applicants", controllers.GetAllApplicants)
}
