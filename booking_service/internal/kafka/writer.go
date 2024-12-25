package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	kafkago "github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	Writer *kafkago.Writer
}

func NewKafkaWriter(addr string, topic string) *KafkaWriter {
	return &KafkaWriter{
		Writer: &kafkago.Writer{
			Addr:                   kafkago.TCP(addr),
			Topic:                  topic,
			AllowAutoTopicCreation: true,
			Balancer:               &kafka.LeastBytes{},
		},
	}
}

func (k *KafkaWriter) Write(ctx context.Context, key []byte, value []byte) error {
	err := k.Writer.WriteMessages(ctx, kafkago.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return err
	}
	return nil
}
