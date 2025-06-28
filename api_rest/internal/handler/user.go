// User HTTP handlers for the API.
package handler

import (
	"net/http"
	"strconv"

	"github.com/UTOL-s/module/api_rest/internal/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userService *service.UserService
	logger      *zap.Logger
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{userService: userService, logger: logger}
}

// CreateUser handles user creation
func (h *UserHandler) CreateUser(c echo.Context) error {
	var request struct {
		Email     string `json:"email"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	if err := c.Bind(&request); err != nil {
		h.logger.Error("failed to bind user data", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	user, err := h.userService.CreateUser(c.Request().Context(), request.Email, request.Username, request.Password, request.FirstName, request.LastName)
	if err != nil {
		h.logger.Error("failed to create user", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	h.logger.Info("user created successfully", zap.String("email", user.Email))
	return c.JSON(http.StatusCreated, user)
}

// GetUser handles user retrieval by ID
func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user, err := h.userService.GetUserByID(c.Request().Context(), uint(id))
	if err != nil {
		h.logger.Error("failed to get user", zap.Error(err), zap.Uint64("id", id))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser handles user updates
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	if err := c.Bind(&request); err != nil {
		h.logger.Error("failed to bind user data", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	user, err := h.userService.UpdateUser(c.Request().Context(), uint(id), request.FirstName, request.LastName)
	if err != nil {
		h.logger.Error("failed to update user", zap.Error(err), zap.Uint64("id", id))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	h.logger.Info("user updated successfully", zap.Uint64("id", id))
	return c.JSON(http.StatusOK, user)
}

// DeleteUser handles user deletion
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	if err := h.userService.DeleteUser(c.Request().Context(), uint(id)); err != nil {
		h.logger.Error("failed to delete user", zap.Error(err), zap.Uint64("id", id))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	h.logger.Info("user deleted successfully", zap.Uint64("id", id))
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

// ListUsers handles user listing with pagination
func (h *UserHandler) ListUsers(c echo.Context) error {
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}

	users, err := h.userService.ListUsers(c.Request().Context(), offset, limit)
	if err != nil {
		h.logger.Error("failed to list users", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve users"})
	}

	return c.JSON(http.StatusOK, users)
}

// SearchUsers handles user search
func (h *UserHandler) SearchUsers(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Search query is required"})
	}

	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}

	users, err := h.userService.SearchUsers(c.Request().Context(), query, offset, limit)
	if err != nil {
		h.logger.Error("failed to search users", zap.Error(err), zap.String("query", query))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search users"})
	}

	return c.JSON(http.StatusOK, users)
}
