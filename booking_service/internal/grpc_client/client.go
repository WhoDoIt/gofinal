package grpcclient

import (
	"context"

	"github.com/WhoDoIt/gofinal/booking_service/protos/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCClient struct {
	conn *grpc.ClientConn
}

func NewClient(addr string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GRPCClient{conn: conn}, err
}

func (g *GRPCClient) IsValidPersonID(ctx context.Context, id int) (bool, error) {
	client := protos.NewHotelServiceClient(g.conn)
	res, err := client.IsValidPersonID(ctx, wrapperspb.Int32(int32(id)))
	return res.Value, err
}

func (g *GRPCClient) IsValidHotelID(ctx context.Context, id int) (bool, error) {
	client := protos.NewHotelServiceClient(g.conn)
	res, err := client.IsValidHotelID(ctx, wrapperspb.Int32(int32(id)))
	return res.Value, err
}

func (g *GRPCClient) IsValidRoomID(ctx context.Context, id int) (bool, error) {
	client := protos.NewHotelServiceClient(g.conn)
	res, err := client.IsValidRoomID(ctx, wrapperspb.Int32(int32(id)))
	return res.Value, err
}

type Room struct {
	RoomID int     `json:"room_id"`
	Type   string  `json:"type"`
	Price  float32 `json:"price"`
}

func (g *GRPCClient) GetAllRoomsInHotel(ctx context.Context, id int) ([]*Room, error) {
	client := protos.NewHotelServiceClient(g.conn)
	res, err := client.GetAllRoomsInHotel(ctx, wrapperspb.Int32(int32(id)))
	if err != nil {
		return nil, err
	}
	rooms := make([]*Room, 0)
	for _, room := range res.Rooms {
		rooms = append(rooms, &Room{RoomID: int(room.RoomId), Type: room.Type, Price: room.Price})
	}
	return rooms, nil
}

func (g *GRPCClient) GetRoomPrice(ctx context.Context, id int) (float32, error) {
	client := protos.NewHotelServiceClient(g.conn)
	res, err := client.GetRoomPrice(ctx, wrapperspb.Int32(int32(id)))
	return res.Value, err
}
