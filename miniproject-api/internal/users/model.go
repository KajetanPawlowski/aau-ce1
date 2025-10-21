package users

import "sync"

// User represents a user entity.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// In-memory "database"
var (
	users  = make(map[int]User)
	nextID = 1
	mu     sync.Mutex
)

// CreateUser stores a new user and returns it.
func CreateUser(username string) User {
	mu.Lock()
	defer mu.Unlock()
	user := User{
		ID:       nextID,
		Username: username,
	}
	users[nextID] = user
	nextID++
	return user
}

// GetUser retrieves a user by ID.
func GetUser(id int) (User, bool) {
	mu.Lock()
	defer mu.Unlock()
	user, exists := users[id]
	return user, exists
}
