gen:
	protoc -I=proto --go_out=booking_service/protos/ --go-grpc_out=booking_service/protos/ ./proto/hotel.proto 
	protoc -I=proto --go_out=hotel_service/protos/ --go-grpc_out=hotel_service/protos/ ./proto/hotel.proto 

compose:
	docker compose up --build;docker compose down -v

crun: gen compose