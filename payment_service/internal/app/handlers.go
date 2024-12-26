package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/WhoDoIt/gofinal/payment_service/internal/models"
)

func (app *Application) paymentHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.PaymentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	app.Dummy.ProcessPayment(context.Background(), req)
	w.WriteHeader(http.StatusOK)
}
