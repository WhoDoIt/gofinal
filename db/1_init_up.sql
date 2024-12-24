CREATE TABLE IF NOT EXISTS users (
	user_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY NOT NULL,
	telegram varchar(100),
	email varchar(100)
);

CREATE TABLE IF NOT EXISTS hotels
(
    hotel_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY NOT NULL,
    owner_id integer REFERENCES users(user_id) NOT NULL,
	name varchar(100) NOT NULL,
	location varchar(100)
);

CREATE TABLE IF NOT EXISTS rooms
(
	room_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY NOT NULL,
	hotel_id integer REFERENCES hotels(hotel_id) NOT NULL,
	type varchar(100),
	price real
)

CREATE TABLE IF NOT EXISTS bookings
(
    booking_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY NOT NULL,
    client_id integer REFERENCES users(user_id) NOT NULL,
	hotel_id integer REFERENCES hotels(hotel_id) NOT NULL,
	room_id integer REFERENCES rooms(room_id) NOT NULL,
	checkin_date date NOT NULL,
	checkout_date date NOT NULL,
	status varchar(10) CHECK (status IN ('pending', 'confirmed', 'cancelled')) NOT NULL
);