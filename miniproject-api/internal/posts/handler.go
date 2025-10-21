package posts

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

// CreatePost godoc
// @Summary Create a post in a thread
// @Tags posts
// @Accept json
// @Produce json
// @Param thread_id path int true "Thread ID"
// @Param post body Post true "Post content"
// @Success 201 {object} Post
// @Failure 400 {object} map[string]string
// @Router /threads/{thread_id}/posts [post]
func (h *Handler) CreatePost(c *gin.Context) {
	threadIDStr := c.Param("thread_id")
	threadIDInt, err := strconv.Atoi(threadIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thread id"})
		return
	}
	threadID := uint(threadIDInt)

	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	post := h.service.CreatePost(req.Content, threadID)
	c.JSON(http.StatusCreated, post)
}

// ListPosts godoc
// @Summary List posts in a thread
// @Tags posts
// @Accept json
// @Produce json
// @Param thread_id path int true "Thread ID"
// @Success 200 {array} Post
// @Router /threads/{thread_id}/posts [get]
func (h *Handler) ListPosts(c *gin.Context) {
	threadIDStr := c.Param("thread_id")
	threadIDInt, _ := strconv.Atoi(threadIDStr)
	threadID := uint(threadIDInt)

	posts := h.service.ListPosts(threadID)
	c.JSON(http.StatusOK, posts)
}
