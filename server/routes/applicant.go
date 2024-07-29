package routes

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ApplicantsRoute(app *fiber.App) {
	app.Post("/api/applicant", middlewares.Limitter, controllers.ApplicantCreate)
	app.Get("/api/applicants", middlewares.Auth, controllers.ApplicantQuery)
	app.Get("/api/applicants/stats", middlewares.Auth, controllers.ApplicantStats)
	app.Get("/api/applicants/download", middlewares.Auth, controllers.ApplicantDownload)
}
