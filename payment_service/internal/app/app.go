package app

import (
	"log"
	"net/http"

	"github.com/WhoDoIt/gofinal/payment_service/internal/dummy"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Dummy    *dummy.Dummy
}

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /payment", app.paymentHandler)

	return app.logRequest(mux)
}
