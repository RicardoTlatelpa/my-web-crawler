package frontier

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Frontier struct {
	writer *kafka.Writer
}

func New(brokerAddress, topic string) * Frontier {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &Frontier{
		writer: writer,
	}
}

func (f* Frontier) Enqueue(url string) {
	err := f.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key: []byte("url"),
			Value: []byte(url),
		},
	)
	if err != nil {
		log.Printf("Failed to enqueue URL to Kafka: %v", err)
	}
}