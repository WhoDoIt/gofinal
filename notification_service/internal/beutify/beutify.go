package beutify

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/WhoDoIt/gofinal/notification_service/internal/delivery"
	"github.com/WhoDoIt/gofinal/notification_service/internal/models"
)

func PrepareMessage(ctx context.Context, deliverer *delivery.Deliverer, key []byte, value []byte) error {
	chat_id, err := strconv.ParseInt(string(key), 10, 64)
	if err != nil {
		return err
	}

	booking := &models.Booking{}
	err = json.Unmarshal(value, &booking)
	if err != nil {
		return err
	}
	test := "NEW BOOKING\n"
	test += "CHECKIN DATE: " + booking.CheckInDate.String() + "\n"
	test += "CHECKOUT DATE: " + booking.CheckOutDate.String() + "\n"
	deliverer.SendMessage(ctx, chat_id, test)
	return nil
}
