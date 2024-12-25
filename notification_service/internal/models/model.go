package models

import "time"

type Booking struct {
	BookingID    int       `json:"booking_id"`
	ClientID     int       `json:"client_id"`
	HotelID      int       `json:"hotel_id"`
	RoomID       int       `json:"room_id"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	Status       string    `json:"status"`
}
