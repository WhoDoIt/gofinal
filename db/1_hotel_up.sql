CREATE TABLE IF NOT EXISTS users (
	user_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY NOT NULL,
	telegram varchar(100) NOT NULL,
	email varchar(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS hotels
(
    hotel_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY NOT NULL,
    owner_id integer REFERENCES users(user_id) NOT NULL,
	name varchar(100) NOT NULL,
	location varchar(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms
(
	room_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY NOT NULL,
	hotel_id integer REFERENCES hotels(hotel_id) NOT NULL,
	type varchar(100),
	price real NOT NULL
)