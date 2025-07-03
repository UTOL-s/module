package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/UTOL-s/module/fxEcho/example/models"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userService *models.UserService
	logger      *zap.Logger
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *models.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

// ListUsers handles GET /api/users
func (h *UserHandler) ListUsers(c echo.Context) error {
	users := h.userService.GetUsers()
	h.logger.Info("listing users", zap.Int("count", len(users)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
		"count": len(users),
	})
}

// GetUser handles GET /api/users/:id
func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, exists := h.userService.GetUserByID(id)
	if !exists {
		h.logger.Warn("user not found", zap.String("id", id))
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "User not found",
		})
	}

	h.logger.Info("retrieved user", zap.String("id", id))
	return c.JSON(http.StatusOK, user)
}

// CreateUser handles POST /api/users
func (h *UserHandler) CreateUser(c echo.Context) error {
	var request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.Bind(&request); err != nil {
		h.logger.Error("failed to bind request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if request.Name == "" || request.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Name and email are required",
		})
	}

	user := h.userService.CreateUser(request.Name, request.Email)
	return c.JSON(http.StatusCreated, user)
}
