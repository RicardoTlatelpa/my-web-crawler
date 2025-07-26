package frontier

import (
	"context"
	"log"
	"strconv"

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

func (f *Frontier) Enqueue(url string, priority int) {
    msg := kafka.Message{
        Key:   []byte("url"),
        Value: []byte(url),
        Headers: []kafka.Header{
            {
                Key:   "priority",
                Value: []byte(strconv.Itoa(priority)),
            },
        },
    }

    err := f.writer.WriteMessages(context.Background(), msg)
    if err != nil {
        log.Printf("Kafka write failed: %v", err)
    }
}