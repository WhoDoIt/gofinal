package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/WhoDoIt/gofinal/booking_service/internal/models"
	"github.com/WhoDoIt/gofinal/booking_service/protos/protos"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (app *Application) BookingCreate(w http.ResponseWriter, r *http.Request) {
	booking := models.Booking{}
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	is, err := app.GRPCClient.IsValidHotelID(r.Context(), wrapperspb.Int32(int32(booking.HotelID)))
	if !is.GetValue() || err != nil {
		app.badRequest(w, err)
		return
	}

	is, err = app.GRPCClient.IsValidPersonID(r.Context(), wrapperspb.Int32(int32(booking.ClientID)))
	if !is.GetValue() || err != nil {
		app.badRequest(w, err)
		return
	}

	is, err = app.GRPCClient.IsValidRoomID(r.Context(), wrapperspb.Int32(int32(booking.RoomID)))
	if !is.GetValue() || err != nil {
		app.badRequest(w, err)
		return
	}

	blocked, err := app.BookingModel.GetNotAvailableRooms(r.Context(), booking.HotelID)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	for _, room_id := range blocked {
		if booking.RoomID == room_id {
			app.badRequest(w, errors.New("room "+strconv.Itoa(room_id)+" is not available"))
		}
	}

	res, err := json.Marshal(booking)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	user, err := app.GRPCClient.GetContact(r.Context(), wrapperspb.Int32(int32(booking.ClientID)))
	if err != nil {
		app.badRequest(w, err)
		return
	}

	err = app.KafkaWriter.Write(r.Context(), []byte(user.Telegram), res)
	if err != nil {
		app.badRequest(w, err)
	}

}

func (app *Application) BookingGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.badRequest(w, err)
		return
	}

	booking, err := app.BookingModel.GetBooking(r.Context(), id)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	app.returnJSON(w, booking, http.StatusOK)
}

func (app *Application) BookingGetByClient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.badRequest(w, err)
		return
	}

	booking, err := app.BookingModel.GetClientBookings(r.Context(), id)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	app.returnJSON(w, booking, http.StatusOK)
}

func (app *Application) BookingGetInHotel(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.badRequest(w, err)
		return
	}

	booking, err := app.BookingModel.GetHotelBookings(r.Context(), id)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	app.returnJSON(w, booking, http.StatusOK)
}

func (app *Application) BookingAvailable(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.badRequest(w, err)
		return
	}
	rooms, err := app.GRPCClient.GetAllRoomsInHotel(r.Context(), wrapperspb.Int32(int32(id)))
	if err != nil {
		app.badRequest(w, err)
		return
	}

	blocked, err := app.BookingModel.GetNotAvailableRooms(r.Context(), id)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	blocked_map := make(map[int]bool)
	for _, id := range blocked {
		blocked_map[id] = true
	}

	res_rooms := make([]*protos.SingleRoom, 0)
	for _, room := range rooms.GetRooms() {
		if !blocked_map[int(room.RoomId)] {
			res_rooms = append(res_rooms, room)
		}
	}

	app.returnJSON(w, res_rooms, http.StatusNotImplemented)
}
