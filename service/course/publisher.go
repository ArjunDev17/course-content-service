package course

import (
	"context"

	"github.com/ArjunDev17/course-content-service/events"
)

type EventPublisher interface {
	PublishCourseCreated(
		ctx context.Context,
		event events.CourseCreatedEvent,
	) error
}
