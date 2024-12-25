package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/WhoDoIt/gofinal/hotel_service/internal/models"
)

func (app *Application) hotelsAll(w http.ResponseWriter, r *http.Request) {
	models, err := app.HotelModel.GetAll(r.Context())
	if err != nil {
		app.internalError(w, err)
		return
	}

	app.returnJSON(w, models, http.StatusOK)
}

type HotelCreateResponse struct {
	HotelID int `json:"hotel_id"`
}

func (app *Application) hotelCreate(w http.ResponseWriter, r *http.Request) {
	hotel := models.Hotel{}
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	hotel.HotelID, err = app.HotelModel.Insert(r.Context(), hotel.OwnerID, hotel.Name, hotel.Location)
	if err != nil {
		app.internalError(w, err)
		return
	}
	response := &HotelCreateResponse{}
	response.HotelID = hotel.HotelID

	app.returnJSON(w, response, http.StatusCreated)
}

func (app *Application) hotelGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.badRequest(w, err)
		return
	}

	model, err := app.HotelModel.Get(r.Context(), id)
	if err != nil {
		app.internalError(w, err)
		return
	}

	app.returnJSON(w, model, http.StatusOK)
}

func (app *Application) hotelGetByOwner(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.badRequest(w, err)
		return
	}

	model, err := app.HotelModel.GetByOwner(r.Context(), id)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	app.returnJSON(w, model, http.StatusOK)
}

type RoomCreateResponse struct {
	RoomID int `json:"room_id"`
}

func (app *Application) roomCreate(w http.ResponseWriter, r *http.Request) {
	room := models.Room{}
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		app.badRequest(w, err)
		return
	}

	roomID, err := app.RoomModel.Insert(r.Context(), room.HotelID, room.Type, room.Price)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	response := RoomCreateResponse{RoomID: roomID}
	app.returnJSON(w, response, http.StatusCreated)
}

func (app *Application) roomPut(w http.ResponseWriter, r *http.Request) {
	room := models.Room{}
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		app.badRequest(w, err)
		return
	}

	err := app.RoomModel.UpdateRoom(r.Context(), room.HotelID, room.RoomID, room.Type, room.Price)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *Application) roomDelete(w http.ResponseWriter, r *http.Request) {
	room := models.Room{}
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		app.badRequest(w, err)
		return
	}

	err := app.RoomModel.DeleteRoom(r.Context(), room.HotelID, room.RoomID)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
