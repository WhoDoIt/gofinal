INSERT INTO users(telegram, email) VALUES('411984714', '');
INSERT INTO users(telegram, email) VALUES('', 'xd@gmail.com');
INSERT INTO users(telegram, email) VALUES('411984714', '');
INSERT INTO users(telegram, email) VALUES('411984714', '');

INSERT INTO hotels(owner_id, name, location) VALUES (1, 'Sirius', 'Sochi');
INSERT INTO hotels(owner_id, name, location) VALUES (2, 'Motel', 'MiddleOfNowhere');
INSERT INTO hotels(owner_id, name, location) VALUES (2, 'Empty', 'Kanash');

INSERT INTO rooms(hotel_id, type, price) VALUES (1, 'Poor', 100);
INSERT INTO rooms(hotel_id, type, price) VALUES (1, 'Middle', 150);
INSERT INTO rooms(hotel_id, type, price) VALUES (1, 'Luxury', 299.99);

INSERT INTO rooms(hotel_id, type, price) VALUES (2, 'Ultrapoor', 9.99);
INSERT INTO rooms(hotel_id, type, price) VALUES (2, 'Middle', 140);
INSERT INTO rooms(hotel_id, type, price) VALUES (2, 'SSSLuxury', 19999.99);