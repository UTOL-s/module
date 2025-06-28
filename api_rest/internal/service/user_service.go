// User service for business logic operations.
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/UTOL-s/module/api_rest/internal/model"
	"github.com/UTOL-s/module/api_rest/internal/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// UserService handles user business logic
type UserService struct {
	repo   *repository.UserRepository
	logger *zap.Logger
}

// NewUserService creates a new user service
func NewUserService(repo *repository.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

// CreateUser creates a new user with validation
func (s *UserService) CreateUser(ctx context.Context, email, username, password, firstName, lastName string) (*model.User, error) {
	s.logger.Info("creating new user", zap.String("email", email), zap.String("username", username))

	// Check if user already exists
	if _, err := s.repo.GetByEmail(ctx, email); err == nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	if _, err := s.repo.GetByUsername(ctx, username); err == nil {
		return nil, fmt.Errorf("user with username %s already exists", username)
	}

	// Create new user
	user := &model.User{
		Email:     email,
		Username:  username,
		Password:  password, // In real app, this should be hashed
		FirstName: firstName,
		LastName:  lastName,
		Role:      "user",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to repository
	if err := s.repo.Create(ctx, user); err != nil {
		s.logger.Error("failed to save user to repository", zap.Error(err))
		return nil, err
	}

	s.logger.Info("user created successfully", zap.Uint("id", user.ID))
	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	s.logger.Debug("getting user by ID", zap.Uint("id", id))
	return s.repo.GetByID(ctx, id)
}

// GetUserByEmail retrieves a user by email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	s.logger.Debug("getting user by email", zap.String("email", email))
	return s.repo.GetByEmail(ctx, email)
}

// GetUserByUsername retrieves a user by username
func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	s.logger.Debug("getting user by username", zap.String("username", username))
	return s.repo.GetByUsername(ctx, username)
}

// UpdateUser updates user information
func (s *UserService) UpdateUser(ctx context.Context, id uint, firstName, lastName string) (*model.User, error) {
	s.logger.Info("updating user", zap.Uint("id", id))

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.FirstName = firstName
	user.LastName = lastName
	user.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, user); err != nil {
		s.logger.Error("failed to save updated user", zap.Error(err))
		return nil, err
	}

	s.logger.Info("user updated successfully", zap.Uint("id", user.ID))
	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	s.logger.Info("deleting user", zap.Uint("id", id))

	if err := s.repo.Delete(ctx, id); err != nil {
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}

	s.logger.Info("user deleted successfully", zap.Uint("id", id))
	return nil
}

// ListUsers retrieves a list of users with pagination
func (s *UserService) ListUsers(ctx context.Context, offset, limit int) ([]*model.User, error) {
	s.logger.Debug("listing users", zap.Int("offset", offset), zap.Int("limit", limit))
	return s.repo.List(ctx, offset, limit)
}

// SearchUsers searches users by query
func (s *UserService) SearchUsers(ctx context.Context, query string, offset, limit int) ([]*model.User, error) {
	s.logger.Debug("searching users", zap.String("query", query), zap.Int("offset", offset), zap.Int("limit", limit))
	return s.repo.Search(ctx, query, offset, limit)
}

// GetUsersCount returns the total number of users
func (s *UserService) GetUsersCount(ctx context.Context) (int64, error) {
	s.logger.Debug("getting users count")
	return s.repo.Count(ctx)
}

// ValidateUserCredentials validates user login credentials
func (s *UserService) ValidateUserCredentials(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if !user.IsActive {
		return nil, fmt.Errorf("user account is deactivated")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}
