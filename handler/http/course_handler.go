package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ArjunDev17/course-content-service/domain"
	requestdto "github.com/ArjunDev17/course-content-service/handler/http/request"
	responsedto "github.com/ArjunDev17/course-content-service/handler/http/response"
	courseservice "github.com/ArjunDev17/course-content-service/service/course"
)

type CourseHandler struct {
	service *courseservice.CourseService
}

func NewCourseHandler(
	service *courseservice.CourseService,
) *CourseHandler {

	return &CourseHandler{
		service: service,
	}
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {

	var req requestdto.CreateCourseRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
			},
		)

		return
	}

	course := &domain.Course{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Instructor:  req.Instructor,
	}

	savedCourse, err := h.service.Create(
		c.Request.Context(),
		course,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": err.Error(),
			},
		)

		return
	}

	response := responsedto.CourseResponse{
		ID:          savedCourse.ID,
		Title:       savedCourse.Title,
		Description: savedCourse.Description,
		Category:    savedCourse.Category,
		Instructor:  savedCourse.Instructor,
	}

	c.JSON(
		http.StatusCreated,
		response,
	)
}