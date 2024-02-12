package models

import (
	"github.com/jinzhu/gorm"
)

// DB ...
var DB *gorm.DB

// ConnectDatabase ...
func ConnectDatabase() {
	database, err := gorm.Open("mysql", "jay:jay123@(localhost:3306)/gin")
	if err != nil {
		panic("Failed to connect to the database.")
	}

	database.AutoMigrate(&Book{})
	DB = database
}
