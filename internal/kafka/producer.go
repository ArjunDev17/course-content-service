package kafka

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

// Producer is responsible for communicating with the Kafka broker.
//
// It wraps kafka-go's Writer so that the rest of the application
// does not directly depend on the kafka-go library.
type Producer struct {
	writer *kafkago.Writer
}

// NewProducer creates a reusable Kafka producer.
//
// brokers:
//   - List of Kafka broker addresses.
//   - Example: []string{"localhost:29092"}
//
// NOTE:
// This Producer is created once when the application starts
// and reused for every publish request.
func NewProducer(brokers []string) *Producer {

	writer := &kafkago.Writer{

		// Kafka Broker Address
		Addr: kafkago.TCP(brokers...),

		// Hash Balancer ensures that messages having the
		// same Key always go to the same partition.
		//
		// Example:
		// Key = Course ID
		//
		// Course Created
		// Course Updated
		// Course Deleted
		//
		// All events for the same Course will be stored
		// in the same partition preserving ordering.
		Balancer: &kafkago.Hash{},

		// Required acknowledgements from Kafka.
		//
		// RequireOne = Leader Broker acknowledges.
		//
		// We'll discuss other options later:
		//   RequireNone
		//   RequireOne
		//   RequireAll
		RequiredAcks: kafkago.RequireOne,

		// Automatically batch multiple messages.
		//
		// Kafka performs much better when messages
		// are sent in batches instead of one by one.
		BatchSize: 100,

		// Flush batch after this duration
		// even if BatchSize is not reached.
		BatchTimeout: 10,
	}

	return &Producer{
		writer: writer,
	}
}

// Publish writes a single message to Kafka.
func (p *Producer) Publish(
	ctx context.Context,
	message kafkago.Message,
) error {

	return p.writer.WriteMessages(
		ctx,
		message,
	)
}

// Close releases Kafka resources.
//
// Should be called during application shutdown.
func (p *Producer) Close() error {

	return p.writer.Close()
}