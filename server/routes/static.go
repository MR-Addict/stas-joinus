package routes

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func ServeStatic(app *fiber.App, staticFiles embed.FS) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(staticFiles),
		PathPrefix: "build",
		Browse:     true,
		MaxAge:     60 * 60 * 24 * 30,
	}))
}
