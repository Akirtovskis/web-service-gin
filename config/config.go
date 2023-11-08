package config

import (
	"log"
	"web-service-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// The hostname 'db' matches the service name in the 'docker-compose.yml' file
	dsn := "postgres://postgres:postgres@db:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	// Perform your migrations and other database setup here
	err = db.AutoMigrate(&models.Album{})
	if err != nil {
		log.Fatal("Database migration error:", err)
	}

	DB = db
}
