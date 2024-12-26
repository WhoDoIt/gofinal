package app

import (
	"context"
	"log"
	"net/http"

	"github.com/WhoDoIt/gofinal/hotel_service/internal/metrics"
	"github.com/WhoDoIt/gofinal/hotel_service/internal/models"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Application struct {
	ErrorLog   *log.Logger
	InfoLog    *log.Logger
	HotelModel models.HotelModelInterface
	RoomModel  models.RoomModelInterface
	Metrics    *metrics.Metrics
}

func (app *Application) Start(ctx context.Context) error {
	return nil
}

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hotels", app.hotelsAll)
	mux.HandleFunc("POST /hotels", app.hotelCreate)
	mux.HandleFunc("GET /hotels/get", app.hotelGet)
	mux.HandleFunc("GET /hotels/owner/get", app.hotelGetByOwner)
	mux.HandleFunc("POST /hotels/room", app.roomCreate)
	mux.HandleFunc("DELETE /hotels/room", app.roomDelete)
	mux.HandleFunc("PUT /hotels/room", app.roomPut)
	mux.Handle("GET /hotels/metrics", promhttp.Handler())
	return app.collectMetrics(app.logRequest(mux))
}
