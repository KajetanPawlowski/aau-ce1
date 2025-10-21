package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler connects HTTP routes to the service.
type Handler struct {
	service *Service
}

// NewHandler creates a new user handler.
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// RegisterRoutes registers all user routes under /users.
func (h *Handler) RegisterRoutes(r *gin.Engine) {
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("", h.createUser)
		usersGroup.GET("/:id", h.getUser)
	}
}

// createUser handles POST /users
func (h *Handler) createUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user := h.service.CreateUser(req.Username)
	c.JSON(http.StatusCreated, user)
}

// getUser handles GET /users/:id
func (h *Handler) getUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, exists := h.service.GetUser(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
