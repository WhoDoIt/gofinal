package app

import (
	"context"
	"log"
	"net/http"

	"github.com/WhoDoIt/gofinal/booking_service/internal/kafka"
	"github.com/WhoDoIt/gofinal/booking_service/internal/models"
	"github.com/WhoDoIt/gofinal/booking_service/protos/protos"
)

type Application struct {
	ErrorLog     *log.Logger
	InfoLog      *log.Logger
	BookingModel *models.BookingModel
	GRPCClient   protos.HotelServiceClient
	KafkaWriter  *kafka.KafkaWriter
}

func (app *Application) Start(ctx context.Context) error {
	return nil
}

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /bookings", app.BookingCreate)
	mux.HandleFunc("GET /bookings/get", app.BookingGet)
	mux.HandleFunc("GET /bookings/hotel/get", app.BookingGetInHotel)
	mux.HandleFunc("GET /bookings/client/get", app.BookingGetByClient)
	mux.HandleFunc("GET /bookings/available/get", app.BookingAvailable)

	return app.logRequest(mux)
}
