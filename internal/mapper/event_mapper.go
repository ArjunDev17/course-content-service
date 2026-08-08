package mapper

import (
	"github.com/google/uuid"

	"github.com/ArjunDev17/course-content-service/domain"
	"github.com/ArjunDev17/course-content-service/events"
	"github.com/ArjunDev17/course-content-service/internal/constants"
)

func ToCourseCreatedEvent(
	course *domain.Course,
) events.CourseCreatedEvent {

	return events.CourseCreatedEvent{
		EventID:    uuid.NewString(),
		EventType:  constants.CourseCreatedEventType,
		OccurredAt: course.CreatedAt,

		CourseID:    course.ID,
		Title:       course.Title,
		Description: course.Description,
		Category:    course.Category,
		Instructor:  course.Instructor,
	}
}
