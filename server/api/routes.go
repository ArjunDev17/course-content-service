package api

import (
	httpHandler "github.com/ArjunDev17/course-content-service/handler/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	courseHandler *httpHandler.CourseHandler,
	healthHandler *httpHandler.HealthHandler,
) *gin.Engine {

	router := gin.Default()

	api := router.Group("/api/v1")

	courseHandler.Register(api)

	router.GET("/health", healthHandler.Health)

	return router
}