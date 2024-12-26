package models

import (
	"fmt"
	"strings"
	"time"
)

type Booking struct {
	BookingID    int      `json:"booking_id"`
	ClientID     int      `json:"client_id"`
	HotelID      int      `json:"hotel_id"`
	RoomID       int      `json:"room_id"`
	CheckInDate  JSONTime `json:"checkin_date"`
	CheckOutDate JSONTime `json:"checkout_date"`
	Status       string   `json:"status"`
}

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	res := fmt.Sprintf("\"%s\"", t.Format(time.DateOnly))
	return []byte(res), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(time.DateOnly, s)
	return
}
