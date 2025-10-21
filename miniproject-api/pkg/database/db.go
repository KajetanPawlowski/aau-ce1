package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// Init initializes the database connection
func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("miniproject.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	log.Println("Database connected")
}
