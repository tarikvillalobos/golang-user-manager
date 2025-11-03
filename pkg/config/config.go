package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	DB = db
	log.Println("✅ database connected")
}
