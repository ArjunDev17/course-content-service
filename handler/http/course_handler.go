package http_handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/ArjunDev17/course-content-service/model"
	mongo_repo "github.com/ArjunDev17/course-content-service/repository/mongo"
	course_svc "github.com/ArjunDev17/course-content-service/service/course"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type CourseHandler struct {
	service course_svc.Service
}

func NewCourseHandler() *CourseHandler {
	repo := mongo_repo.NewCourseRepository()
	svc := course_svc.NewCourseService(repo)
	return &CourseHandler{service: svc}
}

// Register routes in router
func (h *CourseHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/courses", h.CreateCourse)
	rg.GET("/courses", h.ListCourses)
	rg.GET("/courses/:id", h.GetCourse)
	rg.PUT("/courses/:id", h.UpdateCourse)
	rg.DELETE("/courses/:id", h.DeleteCourse)
}

// CreateCourse POST /courses
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req model.Course
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	created, err := h.service.CreateCourse(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// ListCourses GET /courses
func (h *CourseHandler) ListCourses(c *gin.Context) {
	// filters
	filters := map[string]interface{}{}
	if category := c.Query("category"); category != "" {
		filters["category"] = category
	}
	if level := c.Query("level"); level != "" {
		filters["level"] = level
	}
	if tag := c.Query("tag"); tag != "" {
		filters["tags"] = tag // note: simple equality; you can use $in in repo if needed
	}
	// price filters
	if minP := c.Query("min_price"); minP != "" {
		if v, err := strconv.ParseFloat(minP, 64); err == nil {
			filters["price"] = bson.M{"$gte": v}
		}
	}
	// pagination
	page := int64(1)
	limit := int64(20)
	if p := c.Query("page"); p != "" {
		if v, err := strconv.ParseInt(p, 10, 64); err == nil && v > 0 {
			page = v
		}
	}
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.ParseInt(l, 10, 64); err == nil && v > 0 {
			limit = v
		}
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	courses, total, err := h.service.ListCourses(ctx, filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": courses, "total": total, "page": page, "limit": limit})
}

// GetCourse GET /courses/:id
func (h *CourseHandler) GetCourse(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	course, err := h.service.GetCourse(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "course not found or invalid id"})
		return
	}
	c.JSON(http.StatusOK, course)
}

// UpdateCourse PUT /courses/:id
func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	updated, err := h.service.UpdateCourse(ctx, id, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteCourse DELETE /courses/:id
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	if err := h.service.DeleteCourse(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
