package configs

import (
	"log"
	"server/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Applicant{})

	Db = db
}
