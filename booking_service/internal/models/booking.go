package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Booking struct {
	BookingID    int       `json:"booking_id"`
	ClientID     int       `json:"client_id"`
	HotelID      int       `json:"hotel_id"`
	RoomID       int       `json:"room_id"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	Status       string    `json:"status"`
}

type BookingModel struct {
	DB *pgxpool.Pool
}

func (m *BookingModel) Insert(ctx context.Context, client_id int, hotel_id int, room_id int,
	check_in_date time.Time, check_out_date time.Time, status string) (int, error) {
	stmt_check := `SELECT booking_id FROM bookings WHERE room_id=$1 AND max(check_in_date, $1) < min(check_out_date, $2) LIMIT 1`

	var booking_id int
	err := m.DB.QueryRow(ctx, stmt_check, check_in_date, check_out_date).Scan(&booking_id)
	if err != nil {
		return 0, err
	}

	stmt := `INSERT INTO bookings(client_id, hotel_id, room_id, checkin_date, checkout_date, status)
			 VALUES($1, $2, $3, $4, $5, $6) RETURNING booking_id`
	err = m.DB.QueryRow(ctx, stmt, client_id, hotel_id, room_id, check_in_date, check_out_date, status).Scan(&booking_id)
	if err != nil {
		return 0, err
	}
	return booking_id, nil
}

func (m *BookingModel) GetBooking(ctx context.Context, booking_id int) (*Booking, error) {
	stmt := `SELECT * FROM bookings WHERE booking_id=$1`
	booking := &Booking{}
	err := m.DB.QueryRow(ctx, stmt, booking_id).Scan(&booking.BookingID, &booking.ClientID, &booking.HotelID, &booking.RoomID, &booking.CheckInDate, &booking.CheckOutDate, &booking.Status)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (m *BookingModel) GetClientBookings(ctx context.Context, client_id int) ([]*Booking, error) {
	stmt := `SELECT * FROM bookings WHERE client_id=$1`
	rows, err := m.DB.Query(ctx, stmt, client_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookings := []*Booking{}
	for rows.Next() {
		booking := &Booking{}
		if err := rows.Scan(&booking.BookingID, &booking.ClientID, &booking.HotelID, &booking.RoomID, &booking.CheckInDate, &booking.CheckOutDate, &booking.Status); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (m *BookingModel) GetHotelBookings(ctx context.Context, hotel_id int) ([]*Booking, error) {
	stmt := `SELECT * FROM bookings WHERE hotel_id=$1`
	rows, err := m.DB.Query(ctx, stmt, hotel_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookings := []*Booking{}
	for rows.Next() {
		booking := &Booking{}
		if err := rows.Scan(&booking.BookingID, &booking.ClientID, &booking.HotelID, &booking.RoomID, &booking.CheckInDate, &booking.CheckOutDate, &booking.Status); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (m *BookingModel) GetNotAvailableRooms(ctx context.Context, hotel_id int) ([]int, error) {
	stmt := `SELECT room_id FROM bookings WHERE hotel_id=$1 AND checkout_date >= NOW()::DATE AND checkin_date <= NOW()::DATE`
	rows, err := m.DB.Query(ctx, stmt, hotel_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []int{}
	for rows.Next() {
		var room_id int
		if err := rows.Scan(&room_id); err != nil {
			return nil, err
		}
		rooms = append(rooms, room_id)
	}

	return rooms, nil
}
