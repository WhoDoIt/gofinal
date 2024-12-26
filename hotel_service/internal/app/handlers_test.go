package app

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WhoDoIt/gofinal/hotel_service/internal/models"
	mock_models "github.com/WhoDoIt/gofinal/hotel_service/internal/models/mocks"
	"go.uber.org/mock/gomock"
)

func toJsonString(t *testing.T, v any) string {
	json, err := json.Marshal(v)
	if err != nil {
		t.Errorf("couldn't marshal test data: %s", err.Error())
		return ""
	}
	return string(json)
}

func TestHotelsAll(t *testing.T) {
	t.Run("list of hotels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotels := []*models.Hotel{
			{
				HotelID:  2,
				OwnerID:  3,
				Name:     "123",
				Location: "Rooms",
				Rooms:    []*models.Room{},
			},
			{
				HotelID:  1,
				OwnerID:  1,
				Name:     "123",
				Location: "Rooms",
				Rooms: []*models.Room{
					{
						RoomID: 1,
						Type:   "1",
						Price:  2,
					},
					{
						RoomID: 2,
						Type:   "1",
						Price:  2,
					},
				},
			},
		}

		hotels_json, err := json.Marshal(hotels)
		if err != nil {
			t.Errorf("couldn't marshal test data: %s", err.Error())
			return
		}

		hotel.EXPECT().GetAll(gomock.Any()).Return(hotels, nil)

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels", nil)
		res := httptest.NewRecorder()

		app.hotelsAll(res, req)

		if res.Body.String() != string(hotels_json) {
			t.Errorf("expected body: %q but got %q", string(hotels_json), res.Body.String())
			return
		}

		if res.Code != http.StatusOK {
			t.Errorf("expected code: %d but got %d", http.StatusOK, res.Code)
			return
		}
	})
	t.Run("empty list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotels := []*models.Hotel{}

		hotels_json, err := json.Marshal(hotels)
		if err != nil {
			t.Errorf("couldn't marshal test data: %s", err.Error())
			return
		}

		hotel.EXPECT().GetAll(gomock.Any()).Return(hotels, nil)

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels", nil)
		res := httptest.NewRecorder()

		app.hotelsAll(res, req)

		if res.Body.String() != "[]" {
			t.Errorf("expected body: %q but got %q", string(hotels_json), res.Body.String())
			return
		}

		if res.Code != http.StatusOK {
			t.Errorf("expected code: %d but got %d", http.StatusOK, res.Code)
			return
		}
	})
	t.Run("should return 500 on error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().GetAll(gomock.Any()).Return(nil, errors.New("123"))

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels", nil)
		res := httptest.NewRecorder()

		app.hotelsAll(res, req)

		if res.Code != http.StatusInternalServerError {
			t.Errorf("expected code: %d but got %d", http.StatusInternalServerError, res.Code)
			return
		}
	})
}

func TestHotelCreate(t *testing.T) {
	t.Run("empty body", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(0, nil).AnyTimes()

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodPost, "/hotels", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.hotelCreate(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected code: %d but got %d", http.StatusBadRequest, res.Code)
			return
		}
	})
	t.Run("error in database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(0, errors.New("postgres failed"))

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		request := models.Hotel{
			HotelID:  1,
			OwnerID:  2,
			Name:     "3",
			Location: "4",
			Rooms:    []*models.Room{},
		}

		req := httptest.NewRequest(http.MethodPost, "/hotels", strings.NewReader(toJsonString(t, request)))
		res := httptest.NewRecorder()

		app.hotelCreate(res, req)

		if res.Code != http.StatusInternalServerError {
			t.Errorf("expected code: %d but got %d", http.StatusInternalServerError, res.Code)
			return
		}
	})

	t.Run("succesfully created", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		request := models.Hotel{
			HotelID:  1,
			OwnerID:  2,
			Name:     "3",
			Location: "4",
			Rooms:    []*models.Room{},
		}

		response := HotelCreateResponse{
			HotelID: 5,
		}

		hotel.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(5, nil)

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodPost, "/hotels", strings.NewReader(toJsonString(t, request)))
		res := httptest.NewRecorder()

		app.hotelCreate(res, req)

		if res.Code != http.StatusCreated {
			t.Errorf("expected code: %d but got %d", http.StatusCreated, res.Code)
			return
		}

		if res.Body.String() != toJsonString(t, response) {
			t.Errorf("expected body: %q but got %q", toJsonString(t, response), res.Body.String())
			return
		}
	})
}

func TestHotelGet(t *testing.T) {
	request := &models.Hotel{
		HotelID:  1,
		OwnerID:  2,
		Name:     "3",
		Location: "4",
		Rooms: []*models.Room{
			{
				RoomID: 1,
				Type:   "xd",
				Price:  123,
			},
			{
				RoomID: 99,
				Type:   "xd",
				Price:  123,
			},
		},
	}
	t.Run("empty id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().Get(gomock.Any(), gomock.Any()).Return(request, nil).MaxTimes(0)
		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels/get", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.hotelGet(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected code: %d but got %d", http.StatusBadRequest, res.Code)
			return
		}
	})
	t.Run("non int id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().Get(gomock.Any(), gomock.Any()).Return(request, nil).MaxTimes(0)

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels/get?id=abcd", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.hotelGet(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected code: %d but got %d", http.StatusBadRequest, res.Code)
			return
		}
	})
	t.Run("error in database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, errors.New("postgres error"))

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels/get?id=1", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.hotelGet(res, req)

		if res.Code != http.StatusInternalServerError {
			t.Errorf("expected code: %d but got %d", http.StatusInternalServerError, res.Code)
			return
		}
	})

	t.Run("successful get", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().Get(gomock.Any(), gomock.Any()).Return(request, nil)

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotel/get?id=1", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.hotelGet(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("expected code: %d but got %d", http.StatusOK, res.Code)
			return
		}

		if res.Body.String() != toJsonString(t, request) {
			t.Errorf("expected body: %q but got %q", toJsonString(t, request), res.Body.String())
			return
		}
	})
}

func TestOwnerGet(t *testing.T) {
	response := []*models.Hotel{{
		HotelID:  1,
		OwnerID:  2,
		Name:     "3",
		Location: "4",
		Rooms: []*models.Room{
			{
				RoomID: 1,
				Type:   "xd",
				Price:  123,
			},
			{
				RoomID: 99,
				Type:   "xd",
				Price:  123,
			},
		},
	}, {
		HotelID:  1,
		OwnerID:  2,
		Name:     "3",
		Location: "4",
		Rooms: []*models.Room{
			{
				RoomID: 1,
				Type:   "xd",
				Price:  123,
			},
			{
				RoomID: 99,
				Type:   "xd",
				Price:  123,
			},
		},
	},
	}
	t.Run("empty id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().GetByOwner(gomock.Any(), gomock.Any()).Return(response, nil).MaxTimes(0)
		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels/owner/get", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.hotelGetByOwner(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected code: %d but got %d", http.StatusBadRequest, res.Code)
			return
		}
	})
	t.Run("non int id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().GetByOwner(gomock.Any(), gomock.Any()).Return(response, nil).MaxTimes(0)

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels/owner/get?id=abcd", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.hotelGetByOwner(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected code: %d but got %d", http.StatusBadRequest, res.Code)
			return
		}
	})
	t.Run("error in database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().GetByOwner(gomock.Any(), gomock.Any()).Return(nil, errors.New("postgres error"))

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels/owner/get?id=1", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.hotelGetByOwner(res, req)

		if res.Code != http.StatusInternalServerError {
			t.Errorf("expected code: %d but got %d", http.StatusInternalServerError, res.Code)
			return
		}
	})

	t.Run("successful get", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		hotel.EXPECT().GetByOwner(gomock.Any(), gomock.Any()).Return(response, nil)

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodGet, "/hotels/owner/get?id=1", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.hotelGetByOwner(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("expected code: %d but got %d", http.StatusOK, res.Code)
			return
		}

		if res.Body.String() != toJsonString(t, response) {
			t.Errorf("expected body: %q but got %q", toJsonString(t, response), res.Body.String())
			return
		}
	})
}

func TestRoomCreate(t *testing.T) {
	request := models.Room{
		HotelID: 123,
		Type:    "123",
		Price:   123,
	}
	response := &RoomCreateResponse{
		RoomID: 1230,
	}
	t.Run("empty body", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		room.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(5, nil).MaxTimes(0)
		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodPost, "/hotels/room", strings.NewReader(""))
		res := httptest.NewRecorder()

		app.roomCreate(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected code: %d but got %d", http.StatusBadRequest, res.Code)
			return
		}
	})
	t.Run("error in database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		room.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(0, errors.New("postgres error"))

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodPost, "/hotels/room", strings.NewReader(toJsonString(t, request)))
		res := httptest.NewRecorder()

		app.roomCreate(res, req)

		if res.Code != http.StatusInternalServerError {
			t.Errorf("expected code: %d but got %d", http.StatusInternalServerError, res.Code)
			return
		}
	})

	t.Run("successful create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		hotel := mock_models.NewMockHotelModelInterface(ctrl)
		room := mock_models.NewMockRoomModelInterface(ctrl)

		room.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(response.RoomID, nil)

		app := &Application{
			InfoLog:    log.New(io.Discard, "", 0),
			ErrorLog:   log.New(io.Discard, "", 0),
			HotelModel: hotel,
			RoomModel:  room,
		}

		req := httptest.NewRequest(http.MethodPost, "/hotels/room", strings.NewReader(toJsonString(t, request)))
		res := httptest.NewRecorder()

		app.roomCreate(res, req)

		if res.Code != http.StatusCreated {
			t.Errorf("expected code: %d but got %d", http.StatusCreated, res.Code)
			return
		}

		if res.Body.String() != toJsonString(t, response) {
			t.Errorf("expected body: %q but got %q", toJsonString(t, response), res.Body.String())
			return
		}
	})
}
