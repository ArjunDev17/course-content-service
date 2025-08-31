package course_service

import (
	"context"
	"errors"
	"time"

	"github.com/ArjunDev17/course-content-service/model"
	mongorepo "github.com/ArjunDev17/course-content-service/repository/mongo"
)

type Service interface {
	CreateCourse(ctx context.Context, c *model.Course) (*model.Course, error)
	GetCourse(ctx context.Context, id string) (*model.Course, error)
	ListCourses(ctx context.Context, filters map[string]interface{}, page, limit int64) ([]*model.Course, int64, error)
	UpdateCourse(ctx context.Context, id string, update map[string]interface{}) (*model.Course, error)
	DeleteCourse(ctx context.Context, id string) error
}

type courseService struct {
	repo mongorepo.CourseRepository
}

func NewCourseService(repo mongorepo.CourseRepository) Service {
	return &courseService{repo: repo}
}

func (s *courseService) CreateCourse(ctx context.Context, c *model.Course) (*model.Course, error) {
	if c.Title == "" {
		return nil, errors.New("title required")
	}
	// set created times for nested structures if needed
	now := time.Now().UTC()
	c.CreatedAt = now
	c.UpdatedAt = now
	return s.repo.Create(ctx, c)
}

func (s *courseService) GetCourse(ctx context.Context, id string) (*model.Course, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *courseService) ListCourses(ctx context.Context, filters map[string]interface{}, page, limit int64) ([]*model.Course, int64, error) {
	return s.repo.GetAll(ctx, filters, page, limit)
}

func (s *courseService) UpdateCourse(ctx context.Context, id string, update map[string]interface{}) (*model.Course, error) {
	update["updated_at"] = time.Now().UTC()
	return s.repo.Update(ctx, id, update)
}

func (s *courseService) DeleteCourse(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
