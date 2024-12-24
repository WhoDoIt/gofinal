CREATE TABLE IF NOT EXISTS bookings
(
    booking_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY NOT NULL,
    client_id integer NOT NULL,
	hotel_id integer NOT NULL,
	room_id integer NOT NULL,
	checkin_date date NOT NULL,
	checkout_date date NOT NULL,
	status varchar(10) CHECK (status IN ('pending', 'confirmed', 'cancelled')) NOT NULL
);