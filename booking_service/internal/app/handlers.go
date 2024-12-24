package app

import (
	"net/http"
	"strconv"

	grpcclient "github.com/WhoDoIt/gofinal/booking_service/internal/grpc_client"
)

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
	rooms, err := app.GRPCClient.GetAllRoomsInHotel(r.Context(), id)
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

	res_rooms := make([]*grpcclient.Room, 0)
	for _, room := range rooms {
		if !blocked_map[room.RoomID] {
			res_rooms = append(res_rooms, room)
		}
	}

	app.returnJSON(w, res_rooms, http.StatusNotImplemented)
}
