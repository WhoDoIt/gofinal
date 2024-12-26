package app

import (
	"encoding/json"
	"net/http"
)

func (app *Application) methodNotAllowed(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (app *Application) internalError(w http.ResponseWriter, err error) {
	if err != nil {
		app.ErrorLog.Println(err.Error())
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) badRequest(w http.ResponseWriter, err error) {
	if err != nil {
		app.ErrorLog.Println(err.Error())
	}
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func (app *Application) returnJSON(w http.ResponseWriter, v any, status int) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		app.internalError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
