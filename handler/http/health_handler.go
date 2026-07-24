package http_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "UP",
		"service":   "course-content-service",
		"timestamp": time.Now().UTC(),
	})
}
