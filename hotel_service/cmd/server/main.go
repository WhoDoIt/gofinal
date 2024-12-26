package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/WhoDoIt/gofinal/hotel_service/internal/app"
	"github.com/WhoDoIt/gofinal/hotel_service/internal/grpc_server"
	"github.com/WhoDoIt/gofinal/hotel_service/internal/metrics"
	"github.com/WhoDoIt/gofinal/hotel_service/internal/models"
	"github.com/WhoDoIt/gofinal/hotel_service/protos/protos"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

func main() {

	db_login := os.Getenv("HOTELSVC_DB_LOGIN")
	db_password := os.Getenv("HOTELSVC_DB_PASSWORD")
	db_ip := os.Getenv("HOTELSVC_DB_IP")
	db_port := os.Getenv("HOTELSVC_DB_PORT")
	db_database := os.Getenv("HOTELSVC_DB_DATABASE")

	svc_port := os.Getenv("HOTELSVC_PORT")
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
	app := &app.Application{
		InfoLog:    infoLog,
		ErrorLog:   errorLog,
		HotelModel: &models.HotelModel{DB: conn},
		RoomModel:  &models.RoomModel{DB: conn},
		Metrics:    metrics.NewMetrics(),
	}

	go func() {
		app.InfoLog.Printf("HTTP server starting listening on :%s\n", svc_port)
		if err := http.ListenAndServe(":"+svc_port, app.Routes()); err != nil {
			app.ErrorLog.Fatalln(err)
		}
	}()

	ip, err := net.Listen("tcp", "hotel_svc:"+grpc_port)
	if err != nil {
		errorLog.Fatalln(err)
		return
	}

	grpc := grpc.NewServer()
	protos.RegisterHotelServiceServer(grpc, &grpc_server.Server{
		HotelModel: &models.HotelModel{DB: conn},
		RoomModel:  &models.RoomModel{DB: conn},
		UserModel:  &models.UserModel{DB: conn},
	})

	app.InfoLog.Printf("GRPC server starting listening on hotel_svc:%s\n", grpc_port)

	if err := grpc.Serve(ip); err != nil {
		errorLog.Fatalln(err)
		return
	}
}
