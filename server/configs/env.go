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
		MONGODB_URI: os.Getenv("MONGODB_URI"),
		USERNAME:    os.Getenv("USERNAME"),
		PASSWORD:    os.Getenv("PASSWORD"),
	}

	var validate = validator.New()
	if err := validate.Struct(&env); err != nil {
		log.Fatal(err)
	}

	return env
}
