package models

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Room struct {
	RoomID  int     `json:"room_id"`
	HotelID int     `json:"hotel_id"`
	Type    string  `json:"type"`
	Price   float32 `json:"price"`
}

type RoomModel struct {
	DB *pgxpool.Pool
}

func (m *RoomModel) Insert(ctx context.Context, hotel_id int, room_type string, room_price float32) (int, error) {
	stmt := `INSERT INTO rooms(hotel_id, type, price)
				VALUES($1, $2, $3) RETURNING room_id`

	var room_id int
	err := m.DB.QueryRow(ctx, stmt, hotel_id, room_type, room_price).Scan(&room_id)
	if err != nil {
		return 0, err
	}
	return room_id, nil
}

func (m *RoomModel) GetById(ctx context.Context, room_id int) (*Room, error) {
	stmt := `SELECT * FROM rooms WHERE room_id=$1 LIMIT 1`

	room := &Room{}
	err := m.DB.QueryRow(ctx, stmt, room_id).Scan(&room_id, &room.HotelID, &room.Type, &room.Price)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (m *RoomModel) GetAllInHotel(ctx context.Context, hotel_id int) ([]*Room, error) {
	stmt := `SELECT * FROM rooms WHERE hotel_id = $1`
	rows, err := m.DB.Query(ctx, stmt, hotel_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []*Room{}
	for rows.Next() {
		room := &Room{}
		if err := rows.Scan(&room.RoomID, &room.HotelID, &room.Type, &room.Price); err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (m *RoomModel) UpdateRoom(ctx context.Context, hotel_id int, room_id int, room_type string, room_price float32) error {
	stmt := `UPDATE rooms SET type=$1, price=$2 WHERE hotel_id=$3 AND room_id=$4`
	res, err := m.DB.Exec(ctx, stmt, room_type, room_price, hotel_id, room_id)
	if err != nil {
		return err
	}
	if res.RowsAffected() < 1 {
		return errors.New("couldn't update room that doesn't exist")
	}
	return nil
}

func (m *RoomModel) DeleteRoom(ctx context.Context, hotel_id int, room_id int) error {
	stmt := `DELETE FROM rooms WHERE hotel_id=$1 AND room_id=$2`
	res, err := m.DB.Exec(ctx, stmt, hotel_id, room_id)
	if err != nil {
		return err
	}
	if res.RowsAffected() < 1 {
		return errors.New("couldn't delete room that doesn't exist")
	}
	return nil
}
