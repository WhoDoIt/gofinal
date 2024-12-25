package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Hotel struct {
	HotelID  int     `json:"hotel_id"`
	OwnerID  int     `json:"owner_id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rooms    []*Room `json:"rooms"`
}

type HotelModel struct {
	DB *pgxpool.Pool
}

func (m *HotelModel) Insert(ctx context.Context, owner_id int, name string, location string) (int, error) {
	stmt := `INSERT INTO hotels(owner_id, name, location)
						 VALUES($1, $2, $3) RETURNING hotel_id`
	var id int
	err := m.DB.QueryRow(ctx, stmt, owner_id, name, location).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *HotelModel) Get(ctx context.Context, hotel_id int) (*Hotel, error) {
	stmt := `
		SELECT hotels.*, TO_JSON(ARRAY_REMOVE(ARRAY_AGG(rooms), NULL)) AS "userIds"
		FROM hotels
		LEFT JOIN rooms ON hotels.hotel_id=rooms.hotel_id
		WHERE hotels.hotel_id=$1
		GROUP BY hotels.hotel_id;
	`
	row := m.DB.QueryRow(ctx, stmt, hotel_id)

	hotel := &Hotel{}
	if err := row.Scan(&hotel.HotelID, &hotel.OwnerID, &hotel.Name, &hotel.Location, &hotel.Rooms); err != nil {
		return nil, err
	}

	return hotel, nil
}

func (m *HotelModel) GetByOwner(ctx context.Context, owner_id int) ([]*Hotel, error) {
	stmt := `
		SELECT hotels.*, TO_JSON(ARRAY_REMOVE(ARRAY_AGG(rooms), NULL)) AS "userIds"
		FROM hotels
		LEFT JOIN rooms ON hotels.hotel_id=rooms.hotel_id
		WHERE hotels.owner_id=$1
		GROUP BY hotels.hotel_id;
	`
	rows, err := m.DB.Query(ctx, stmt, owner_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hotels := []*Hotel{}
	for rows.Next() {
		hotel := &Hotel{}
		if err := rows.Scan(&hotel.HotelID, &hotel.OwnerID, &hotel.Name, &hotel.Location, &hotel.Rooms); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}

func (m *HotelModel) GetAll(ctx context.Context) ([]*Hotel, error) {
	stmt := `
		SELECT hotels.*, TO_JSON(ARRAY_REMOVE(ARRAY_AGG(rooms), NULL)) AS "userIds"
		FROM hotels
		LEFT JOIN rooms ON hotels.hotel_id=rooms.hotel_id
		GROUP BY hotels.hotel_id;
	`
	rows, err := m.DB.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hotels := []*Hotel{}
	for rows.Next() {
		hotel := &Hotel{}
		if err := rows.Scan(&hotel.HotelID, &hotel.OwnerID, &hotel.Name, &hotel.Location, &hotel.Rooms); err != nil {
			return nil, err
		}

		hotels = append(hotels, hotel)
	}

	return hotels, nil
}
