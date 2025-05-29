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
		App: models.AppConfig{
			StartTime: os.Getenv("START_TIME"),
			EndTime:   os.Getenv("END_TIME"),
		},

		Server: models.ServerConfig{
			Port:          os.Getenv("PORT"),
			Cors:          os.Getenv("CORS"),
			AdminPassword: os.Getenv("ADMIN_PASS"),
		},
	}

	if Config.Server.Port == "" {
		Config.Server.Port = "4000"
	}

	if err := validator.New().Struct(&Config); err != nil {
		log.Fatal(err)
	}
}

func SetupCors(app *fiber.App) {
	if Config.Server.Cors == "*"{
		log.Fatal("CORS cannot be set to '*' for security reasons. Please specify allowed origins.")
	}

	if Config.Server.Cors != "" {
		app.Use(cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins:     Config.Server.Cors,
		}))
	}
}
