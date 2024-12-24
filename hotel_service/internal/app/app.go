package app

import (
	"context"
	"log"
	"net/http"

	"github.com/WhoDoIt/gofinal/hotel_service/internal/models"
)

type Application struct {
	ErrorLog   *log.Logger
	InfoLog    *log.Logger
	HotelModel *models.HotelModel
	RoomModel  *models.RoomModel
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
	return app.logRequest(mux)
}
