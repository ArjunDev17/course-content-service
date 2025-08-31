package api

import (
	httpHandler "github.com/ArjunDev17/course-content-service/handler/http"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// versioned API group
	v1 := router.Group("/api/v1")

	courseHandler := httpHandler.NewCourseHandler()
	courseHandler.Register(v1)

	// add middleware, health checks, metrics etc here
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return router
}
