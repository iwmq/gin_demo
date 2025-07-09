package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ...
var DB *gorm.DB

// ConnectDatabase ...
func ConnectDatabase() {
	gin_postgres_dsn := os.Getenv("GIN_POSTGRES_DSN")
	database, err := gorm.Open(postgres.Open(gin_postgres_dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database.")
	} else {
		println("Connected to the database successfully.")
	}

	database.AutoMigrate(&Book{})
	DB = database
}
