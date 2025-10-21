package users

// Service defines the user business logic layer.
type Service struct{}

// NewService creates a new user service.
func NewService() *Service {
	return &Service{}
}

// CreateUser wraps model.CreateUser.
func (s *Service) CreateUser(username string) User {
	return CreateUser(username)
}

// GetUser wraps model.GetUser.
func (s *Service) GetUser(id int) (User, bool) {
	return GetUser(id)
}
