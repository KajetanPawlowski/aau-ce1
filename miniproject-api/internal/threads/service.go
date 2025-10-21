package threads

import "miniproject-api/pkg/database"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// CreateThread creates a new thread
func (s *Service) CreateThread(title string, userID uint) Thread {
	thread := Thread{
		Title:  title,
		UserID: userID,
	}
	database.DB.Create(&thread)
	return thread
}

// GetThread retrieves a thread by ID
func (s *Service) GetThread(id uint) (Thread, bool) {
	var thread Thread
	result := database.DB.Preload("User").First(&thread, id)
	if result.Error != nil {
		return Thread{}, false
	}
	return thread, true
}

// ListThreads returns all threads
func (s *Service) ListThreads() []Thread {
	var threads []Thread
	database.DB.Preload("User").Find(&threads)
	return threads
}
