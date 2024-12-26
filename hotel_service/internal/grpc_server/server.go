package grpc_server

import (
	"context"

	"github.com/WhoDoIt/gofinal/hotel_service/internal/models"
	"github.com/WhoDoIt/gofinal/hotel_service/protos/protos"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Server struct {
	protos.UnimplementedHotelServiceServer
	HotelModel *models.HotelModel
	RoomModel  *models.RoomModel
	UserModel  *models.UserModel
}

func (s *Server) IsValidPersonID(ctx context.Context, id *wrappers.Int32Value) (*wrappers.BoolValue, error) {
	_, err := s.UserModel.Get(ctx, int(id.GetValue()))
	return wrapperspb.Bool(err == nil), nil
}
func (s *Server) IsValidHotelID(ctx context.Context, id *wrappers.Int32Value) (*wrappers.BoolValue, error) {
	_, err := s.HotelModel.Get(ctx, int(id.GetValue()))
	return wrapperspb.Bool(err == nil), nil
}
func (s *Server) IsValidRoomID(ctx context.Context, id *wrappers.Int32Value) (*wrappers.BoolValue, error) {
	_, err := s.RoomModel.GetById(ctx, int(id.GetValue()))
	return wrapperspb.Bool(err == nil), nil
}
func (s *Server) GetAllRoomsInHotel(ctx context.Context, id *wrappers.Int32Value) (*protos.Rooms, error) {
	hotels, err := s.HotelModel.Get(ctx, int(id.GetValue()))
	if err != nil {
		return nil, err
	}
	rooms := protos.Rooms{
		Rooms: make([]*protos.SingleRoom, 0),
	}

	for _, room := range hotels.Rooms {
		room_proto := protos.SingleRoom{
			Price:  room.Price,
			Type:   room.Type,
			RoomId: int32(room.RoomID),
		}
		rooms.Rooms = append(rooms.Rooms, &room_proto)
	}
	return &rooms, nil
}

func (s *Server) GetRoom(ctx context.Context, id *wrappers.Int32Value) (*protos.SingleRoom, error) {
	room, err := s.RoomModel.GetById(ctx, int(id.GetValue()))
	if err != nil {
		return nil, err
	}
	return &protos.SingleRoom{
		RoomId: int32(room.RoomID),
		Type:   room.Type,
		Price:  room.Price,
	}, nil
}
func (s *Server) GetHotel(ctx context.Context, id *wrappers.Int32Value) (*protos.Hotel, error) {
	hotel, err := s.HotelModel.Get(ctx, int(id.GetValue()))
	if err != nil {
		return nil, err
	}
	return &protos.Hotel{
		HotelId:  int32(hotel.HotelID),
		OwnerId:  int32(hotel.OwnerID),
		Name:     hotel.Name,
		Location: hotel.Location,
	}, nil
}
func (s *Server) GetContact(ctx context.Context, id *wrappers.Int32Value) (*protos.Contact, error) {
	user, err := s.UserModel.Get(ctx, int(id.GetValue()))
	if err != nil {
		return nil, err
	}
	return &protos.Contact{
		UserId:   int32(user.UserID),
		Telegram: user.Telegram,
		Email:    user.Email,
	}, nil
}
