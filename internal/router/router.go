package router

import (
	"github.com/gin-gonic/gin"

	httphandler "github.com/ArjunDev17/course-content-service/handler/http"
)

func NewRouter(
	courseHandler *httphandler.CourseHandler,
) *gin.Engine {

	router := gin.Default()

	api := router.Group("/api/v1")

	{
		api.POST("/courses", courseHandler.CreateCourse)
	}

	return router
}