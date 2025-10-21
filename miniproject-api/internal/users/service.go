package users

import "miniproject-api/pkg/database"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// CreateUser saves a new user in the DB
func (s *Service) CreateUser(username string) User {
	user := User{Username: username}
	database.DB.Create(&user)
	return user
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(id uint) (User, bool) {
	var user User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return User{}, false
	}
	return user, true
}
