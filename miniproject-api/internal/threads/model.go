package threads

import (
	"miniproject-api/pkg/database"
	"miniproject-api/internal/users"
	"gorm.io/gorm"
)

// Thread represents a discussion thread owned by a user
type Thread struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	UserID uint   `json:"user_id"`
	User   users.User
}

// AutoMigrate runs migrations for threads
func AutoMigrate() {
	database.DB.AutoMigrate(&Thread{})
}
