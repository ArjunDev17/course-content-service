package course

import (
	"context"

	"github.com/ArjunDev17/course-content-service/events"
)

type Publisher interface {
	PublishCourseCreated(
		ctx context.Context,
		event events.CourseCreatedEvent,
	) error
}