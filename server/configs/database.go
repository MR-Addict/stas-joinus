package configs

import (
	"log"
	"os"
	"server/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		if err := os.MkdirAll("data", 0755); err != nil {
			log.Fatal(err)
		}
	}
	db, err := gorm.Open(sqlite.Open("data/data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Applicant{})

	Db = db
}
