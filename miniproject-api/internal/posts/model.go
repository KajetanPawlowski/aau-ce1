package posts

import (
	"miniproject-api/pkg/database"
	"miniproject-api/internal/threads"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content  string `json:"content" gorm:"not null"`
	ThreadID uint   `json:"thread_id"`
	Thread   threads.Thread
}

// AutoMigrate runs migrations for posts
func AutoMigrate() {
	database.DB.AutoMigrate(&Post{})
}
