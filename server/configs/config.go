package configs

import (
	"log"
	"os"
	"server/models"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Config models.ConfigType

func LoadConfig() {
	godotenv.Load()

	Config = models.ConfigType{
		Port:          os.Getenv("PORT"),
		Cors:          os.Getenv("CORS"),
		AdminPassword: os.Getenv("ADMIN_PASS"),
	}

	if Config.Port == "" {
		Config.Port = "4000"
	}

	if err := validator.New().Struct(&Config); err != nil {
		log.Fatal(err)
	}
}

func SetupCors(app *fiber.App) {
	if Config.Cors != "" {
		app.Use(cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins:     Config.Cors,
		}))
	}
}
