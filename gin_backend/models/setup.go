package models

import (
	"os"

	"github.com/jinzhu/gorm"
)

// DB ...
var DB *gorm.DB

// ConnectDatabase ...
func ConnectDatabase() {
	gin_mysql_dsn := os.Getenv("GIN_MYSQL_DSN")
	database, err := gorm.Open("mysql", gin_mysql_dsn)
	if err != nil {
		panic("Failed to connect to the database.")
	}

	database.AutoMigrate(&Book{})
	DB = database
}
