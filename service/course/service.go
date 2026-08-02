package course

import (
	"context"

	"github.com/ArjunDev17/course-content-service/domain"
)

type CourseService struct {
	repository Repository
}

func New(
	repository Repository,
) *CourseService {
	return &CourseService{
		repository: repository,
	}
}

func (s *CourseService) Create(
	ctx context.Context,
	course *domain.Course,
) (*domain.Course, error) {

	if err := s.repository.Create(ctx, course); err != nil {
		return nil, err
	}

	return course, nil
}