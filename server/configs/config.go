package configs

import (
	"flag"
	"os"
	"server/models"
	"strconv"

	"github.com/joho/godotenv"
)

var Config models.ConfigType

func LoadConfig() {
	flagHelp := flag.Bool("help", false, "Print help message")
	flagPort := flag.Int("port", 4000, "Provide a port number")
	flagPassword := flag.String("pass", "", "Provide a login password")
	flag.Parse()
	godotenv.Load()

	// print help message
	if *flagHelp {
		flag.PrintDefaults()
		os.Exit(1)
	}

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
		flag.PrintDefaults()
		os.Exit(1)
	}
}
