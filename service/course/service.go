package course

import (
	"context"

	"github.com/ArjunDev17/course-content-service/domain"
	"github.com/ArjunDev17/course-content-service/internal/mapper"
)

type CourseService struct {
	repository Repository
	publisher  Publisher
}

func New(
	repository Repository,
	publisher Publisher,
) *CourseService {

	return &CourseService{
		repository: repository,
		publisher:  publisher,
	}
}

func (s *CourseService) Create(
	ctx context.Context,
	course *domain.Course,
) (*domain.Course, error) {

	// Step 1: Save course into PostgreSQL
	if err := s.repository.Create(ctx, course); err != nil {
		return nil, err
	}

	// Step 2: Convert domain object to Kafka event
	event := mapper.ToCourseCreatedEvent(course)

	// Step 3: Publish event
	if err := s.publisher.PublishCourseCreated(ctx, event); err != nil {
		return nil, err
	}

	return course, nil
}