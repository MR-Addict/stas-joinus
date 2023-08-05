package configs

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var Env EnvType

type EnvType struct {
	JWT_SECRET string `validate:"required"`
	PASSWORD   string `validate:"required"`
}

func LoadEnv() {
	godotenv.Load()
	env := EnvType{
		PASSWORD:   os.Getenv("PASSWORD"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	var validate = validator.New()
	if err := validate.Struct(&env); err != nil {
		log.Fatal(err)
	}

	Env = env
}
