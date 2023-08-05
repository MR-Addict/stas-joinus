package configs

import (
	"flag"
	"log"
	"os"
	"server/models"

	"github.com/joho/godotenv"
)

var Config models.ConfigType

func LoadConfig() {
	var flagPassword, envPassowrd string
	flag.StringVar(&flagPassword, "p", "", "Provide a dashboard login password")
	flag.Parse()

	if flagPassword != "" {
		Config.PASSWORD = flagPassword
	} else {
		godotenv.Load()

		envPassowrd = os.Getenv("PASSWORD")
		if envPassowrd != "" {
			Config.PASSWORD = envPassowrd
		} else {
			log.Fatal("Provide a dashboard login password")
		}
	}
}
