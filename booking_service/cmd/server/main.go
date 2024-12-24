package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/WhoDoIt/gofinal/booking_service/internal/app"
	grpcclient "github.com/WhoDoIt/gofinal/booking_service/internal/grpc_client"
	"github.com/WhoDoIt/gofinal/booking_service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	db_login := os.Getenv("BOOKINGSVC_DB_LOGIN")
	db_password := os.Getenv("BOOKINGSVC_DB_PASSWORD")
	db_ip := os.Getenv("BOOKINGSVC_DB_IP")
	db_port := os.Getenv("BOOKINGSVC_DB_PORT")
	db_database := os.Getenv("BOOKINGSVC_DB_DATABASE")

	svc_port := os.Getenv("BOOKINGSVC_PORT")
	grpc_port := os.Getenv("GRPC_PORT")

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
	client, err := grpcclient.NewClient("hotel_svc:" + grpc_port)
	if err != nil {
		errorLog.Fatalln(err)
		return
	}
	app := &app.Application{
		InfoLog:      infoLog,
		ErrorLog:     errorLog,
		BookingModel: &models.BookingModel{DB: conn},
		GRPCClient:   client,
	}

	app.InfoLog.Printf("HTTP server starting listening on :%s\n", svc_port)
	if err := http.ListenAndServe(":"+svc_port, app.Routes()); err != nil {
		app.ErrorLog.Fatalln(err)
	}
}
