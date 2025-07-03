package models

import (
	"time"

	"go.uber.org/zap"
)

// User represents a user in our system
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserService handles user-related business logic
type UserService struct {
	logger *zap.Logger
	users  map[string]User
}

// NewUserService creates a new user service
func NewUserService(logger *zap.Logger) *UserService {
	users := map[string]User{
		"1": {ID: "1", Name: "Alice Johnson", Email: "alice@example.com"},
		"2": {ID: "2", Name: "Bob Smith", Email: "bob@example.com"},
		"3": {ID: "3", Name: "Charlie Brown", Email: "charlie@example.com"},
	}

	return &UserService{
		logger: logger,
		users:  users,
	}
}

// GetUsers returns all users
func (s *UserService) GetUsers() []User {
	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// GetUserByID returns a user by ID
func (s *UserService) GetUserByID(id string) (*User, bool) {
	user, exists := s.users[id]
	return &user, exists
}

// CreateUser creates a new user
func (s *UserService) CreateUser(name, email string) User {
	id := time.Now().Format("20060102150405") // Simple ID generation
	user := User{ID: id, Name: name, Email: email}
	s.users[id] = user
	s.logger.Info("created new user", zap.String("id", id), zap.String("name", name))
	return user
}
