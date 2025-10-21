package users

import (
	"miniproject-api/pkg/database"
	"gorm.io/gorm"
)

// User represents a user entity in the DB
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
}

// AutoMigrate runs GORM migrations
func AutoMigrate() {
	database.DB.AutoMigrate(&User{})
}
