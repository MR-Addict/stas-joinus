package configs

import (
	"flag"
	"fmt"
	"os"
	"server/models"
	"strconv"

	"github.com/joho/godotenv"
)

var Config models.ConfigType

func LoadConfig() {
	flagPort := flag.Int("port", 4000, "Provide a port number")
	flagPassword := flag.String("pass", "", "Provide a login password")
	flag.Parse()
	godotenv.Load()

	// parse port
	if envPort, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		Config.Port = envPort
	} else {
		Config.Port = *flagPort
	}

	// parse password
	if envPassowrd := os.Getenv("PASSWORD"); envPassowrd != "" {
		Config.Password = envPassowrd
	} else if *flagPassword != "" {
		Config.Password = *flagPassword
	} else {
		fmt.Println("You need to provide a login password")
		os.Exit(1)
	}
}
