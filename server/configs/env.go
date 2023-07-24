package configs

import (
	"log"
	"os"

	"server/models"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var Env = loadEnv()

func loadEnv() models.Env {
	godotenv.Load()
	env := models.Env{
		JWT_SECRET:     os.Getenv("JWT_SECRET"),
		MONGODB_URI:    os.Getenv("MONGODB_URI"),
		ADMIN_USERNAME: os.Getenv("ADMIN_USERNAME"),
		ADMIN_PASSWORD: os.Getenv("ADMIN_PASSWORD"),
	}

	var validate = validator.New()
	if err := validate.Struct(&env); err != nil {
		log.Fatal(err)
	}

	return env
}
