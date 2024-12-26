package dummy

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/WhoDoIt/gofinal/payment_service/internal/models"
)

type Dummy struct {
	Client  *http.Client
	InfoLog *log.Logger
}

func (d *Dummy) ProcessPayment(ctx context.Context, req *models.PaymentRequest) {
	go func() {
		time.Sleep(5 * time.Second)
		res := models.PaymentResponse{
			UniqueID: req.UniqueID,
			Status:   `confirmed`,
		}

		randres := rand.Int() % 2
		if randres == 0 {
			res.Status = `cancelled`
		}

		res_data, err := json.Marshal(res)
		if err != nil {
			d.InfoLog.Println(err.Error())
			return
		}

		request, err := http.NewRequestWithContext(ctx, "POST", req.WebhookURL, bytes.NewBuffer(res_data))
		if err != nil {
			d.InfoLog.Println(err.Error())
			return
		}

		response, err := d.Client.Do(request)
		if err != nil {
			d.InfoLog.Println(err.Error())
			return
		}
		defer response.Body.Close()
	}()
}
