package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/WhoDoIt/gofinal/booking_service/internal/app"
	"github.com/WhoDoIt/gofinal/booking_service/internal/kafka"
	"github.com/WhoDoIt/gofinal/booking_service/internal/models"
	"github.com/WhoDoIt/gofinal/booking_service/protos/protos"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	db_login := os.Getenv("BOOKINGSVC_DB_LOGIN")
	db_password := os.Getenv("BOOKINGSVC_DB_PASSWORD")
	db_ip := os.Getenv("BOOKINGSVC_DB_IP")
	db_port := os.Getenv("BOOKINGSVC_DB_PORT")
	db_database := os.Getenv("BOOKINGSVC_DB_DATABASE")

	svc_port := os.Getenv("BOOKINGSVC_PORT")
	grpc_port := os.Getenv("GRPC_PORT")
	kafka_port := os.Getenv("KAFKA_PORT")
	kafka_topic := os.Getenv("KAFKA_TOPIC")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_login, db_password, db_ip, db_port, db_database)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	conn, err := pgxpool.New(ctx, dsn)

	if err != nil {
		fmt.Fprint(os.Stdout, err)
		return
	}

	if err = conn.Ping(ctx); err != nil {
		fmt.Fprint(os.Stdout, err)
		return
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	grpc_client, err := grpc.NewClient(":"+grpc_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorLog.Fatalln(err)
		return
	}
	app := &app.Application{
		InfoLog:      infoLog,
		ErrorLog:     errorLog,
		BookingModel: &models.BookingModel{DB: conn},
		GRPCClient:   protos.NewHotelServiceClient(grpc_client),
		KafkaWriter:  kafka.NewKafkaWriter("kafka:"+kafka_port, kafka_topic),
	}

	app.InfoLog.Printf("HTTP server starting listening on :%s\n", svc_port)
	if err := http.ListenAndServe(":"+svc_port, app.Routes()); err != nil {
		app.ErrorLog.Fatalln(err)
	}
}
