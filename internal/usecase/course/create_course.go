package course

import (
	"context"

	"github.com/ArjunDev17/course-content-service/domain"
	"github.com/ArjunDev17/course-content-service/internal/mapper"
	courseservice "github.com/ArjunDev17/course-content-service/service/course"
)

type CreateCourseUseCase struct {
	repository courseservice.Repository
	publisher  courseservice.Publisher
}

// NewCreateCourseUseCase creates a new instance of CreateCourseUseCase.
func NewCreateCourseUseCase(
	repository courseservice.Repository,
	publisher courseservice.Publisher,
) *CreateCourseUseCase {

	return &CreateCourseUseCase{
		repository: repository,
		publisher:  publisher,
	}
}

// Execute handles the complete business flow for creating a course.
//
// Flow:
//
//	1. Save course into PostgreSQL
//	2. Convert Course -> CourseCreatedEvent
//	3. Publish event to Kafka
//	4. Return created course
func (u *CreateCourseUseCase) Execute(
	ctx context.Context,
	course *domain.Course,
) (*domain.Course, error) {

	// Persist the course.
	if err := u.repository.Create(
		ctx,
		course,
	); err != nil {
		return nil, err
	}

	// Convert the domain model into an event.
	event := mapper.ToCourseCreatedEvent(course)

	// Publish the event to Kafka.
	if err := u.publisher.PublishCourseCreated(
		ctx,
		event,
	); err != nil {
		return nil, err
	}

	return course, nil
}