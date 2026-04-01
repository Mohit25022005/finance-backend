package config

import (
	"log"

	"finance-backend/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("finance.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed")
	}

	db.AutoMigrate(&models.User{}, &models.Record{})

	DB = db
}