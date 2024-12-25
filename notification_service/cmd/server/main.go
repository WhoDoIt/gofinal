//go:build unix

package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/WhoDoIt/gofinal/notification_service/internal/kafka"
	"github.com/go-telegram/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(os.Getenv("TG_TOKEN"))
	if err != nil {
		panic(err)
	}

	kafka_port := os.Getenv("KAFKA_PORT")
	kafka_topic := os.Getenv("KAFKA_TOPIC")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	consumer := kafka.NewKafkaReader("kafka:"+kafka_port, kafka_topic, 0, infoLog, errorLog, b)
	consumer.Read(context.Background())

	b.Start(ctx)
}
