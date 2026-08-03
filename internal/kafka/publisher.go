package kafka

import (
	"context"
	"encoding/json"

	"github.com/ArjunDev17/course-content-service/events"
	"github.com/ArjunDev17/course-content-service/internal/constants"

	kafkago "github.com/segmentio/kafka-go"
)

// Publisher publishes business events to Kafka.
type Publisher struct {
	producer *Producer
}

// NewPublisher creates a new Kafka publisher.
func NewPublisher(
	producer *Producer,
) *Publisher {

	return &Publisher{
		producer: producer,
	}
}

// PublishCourseCreated publishes the CourseCreatedEvent
// to the "course.created" Kafka topic.
func (p *Publisher) PublishCourseCreated(
	ctx context.Context,
	event events.CourseCreatedEvent,
) error {

	// Convert event to JSON
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Build Kafka message
	message := kafkago.Message{

		// Kafka Topic
		Topic: constants.CourseCreatedTopic,

		// Message Key
		//
		// We use CourseID as the key so that
		// all events related to the same course
		// always go to the same partition.
		Key: []byte(event.CourseID),

		// Actual Event Payload
		Value: payload,
	}

	// Publish message to Kafka
	return p.producer.Publish(
		ctx,
		message,
	)
}