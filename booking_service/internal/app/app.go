package app

import (
	"context"
	"log"
	"net/http"

	grpcclient "github.com/WhoDoIt/gofinal/booking_service/internal/grpc_client"
	"github.com/WhoDoIt/gofinal/booking_service/internal/models"
)

type Application struct {
	ErrorLog     *log.Logger
	InfoLog      *log.Logger
	BookingModel *models.BookingModel
	GRPCClient   *grpcclient.GRPCClient
}

func (app *Application) Start(ctx context.Context) error {
	return nil
}

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("POST /bookings", app.Bo)
	mux.HandleFunc("GET /bookings/get", app.BookingGet)
	mux.HandleFunc("GET /bookings/hotel/get", app.BookingGetInHotel)
	mux.HandleFunc("GET /bookings/client/get", app.BookingGetByClient)
	mux.HandleFunc("GET /bookings/available/get", app.BookingAvailable)

	return app.logRequest(mux)
}
