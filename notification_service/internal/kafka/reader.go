package kafka

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/WhoDoIt/gofinal/notification_service/internal/delivery"
	"github.com/WhoDoIt/gofinal/notification_service/internal/models"
	"github.com/go-telegram/bot"
	kafkago "github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	reader    *kafkago.Reader
	deliverer *delivery.Deliverer
	infoLog   *log.Logger
	errorLog  *log.Logger
}

func NewKafkaReader(broker string, topic string, handler any, infoLog *log.Logger, errorLog *log.Logger, bot *bot.Bot) *KafkaReader {
	r := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
	return &KafkaReader{
		reader:    r,
		deliverer: &delivery.Deliverer{Bot: bot},
		infoLog:   infoLog,
		errorLog:  errorLog,
	}
}

func (c *KafkaReader) Read(ctx context.Context) {
	go func(r *kafkago.Reader) {
		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				c.errorLog.Println(err)
				continue
			}
			booking := &models.Booking{}
			err = json.Unmarshal(m.Value, &booking)
			if err != nil {
				c.errorLog.Println(err)
				continue
			}

			chat_id, err := strconv.ParseInt(string(m.Key), 10, 64)
			if err != nil {
				c.errorLog.Println(err)
				continue
			}
			c.deliverer.SendMessage(ctx, chat_id, booking.CheckOutDate.String())
			c.infoLog.Printf("got message: %s\n", string(m.Value))
		}
	}(c.reader)
}
