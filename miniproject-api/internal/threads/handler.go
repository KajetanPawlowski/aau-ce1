package threads

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	threadGroup := r.Group("/threads")
	{
		threadGroup.POST("", h.CreateThread)
		threadGroup.GET("/:id", h.GetThread)
		threadGroup.GET("", h.ListThreads)
	}
}
// @Summary Create a new thread
// @Tags threads
// @Accept json
// @Produce json
// @Param thread body Thread true "Thread data"
// @Success 201 {object} Thread
// @Router /threads [post]
func (h *Handler) CreateThread(c *gin.Context) {
	var req struct {
		Title  string `json:"title" binding:"required"`
		UserID uint   `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	thread := h.service.CreateThread(req.Title, req.UserID)
	c.JSON(http.StatusCreated, thread)
}

// @Summary Get a thread by ID
// @Tags threads
// @Produce json
// @Param id path int true "Thread ID"
// @Success 200 {object} Thread
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /threads/{id} [get]
func (h *Handler) GetThread(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(idInt)

	thread, exists := h.service.GetThread(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "thread not found"})
		return
	}
	c.JSON(http.StatusOK, thread)
}

// @Summary List all threads
// @Tags threads
// @Produce json
// @Success 200 {array} Thread
// @Router /threads [get]
func (h *Handler) ListThreads(c *gin.Context) {
	threads := h.service.ListThreads()
	c.JSON(http.StatusOK, threads)
}
