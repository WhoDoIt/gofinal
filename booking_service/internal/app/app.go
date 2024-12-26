package app

import (
	"context"
	"log"
	"net/http"

	"github.com/WhoDoIt/gofinal/booking_service/internal/kafka"
	"github.com/WhoDoIt/gofinal/booking_service/internal/metrics"
	"github.com/WhoDoIt/gofinal/booking_service/internal/models"
	"github.com/WhoDoIt/gofinal/booking_service/protos/protos"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Application struct {
	ErrorLog       *log.Logger
	InfoLog        *log.Logger
	BookingModel   *models.BookingModel
	GRPCClient     protos.HotelServiceClient
	KafkaWriter    *kafka.KafkaWriter
	Metrics        *metrics.Metrics
	PaymentURL     string
	PaymentWebhook string
}

func (app *Application) Start(ctx context.Context) error {
	return nil
}

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /bookings", app.BookingCreate)
	mux.HandleFunc("POST /bookings/webhook", app.BookingWebhook)
	mux.HandleFunc("GET /bookings/get", app.BookingGet)
	mux.HandleFunc("GET /bookings/hotel/get", app.BookingGetInHotel)
	mux.HandleFunc("GET /bookings/client/get", app.BookingGetByClient)
	mux.HandleFunc("GET /bookings/available/get", app.BookingAvailable)
	mux.Handle("GET /bookings/metrics", promhttp.Handler())

	return app.collectMetrics(app.logRequest(mux))
}
