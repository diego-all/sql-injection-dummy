package main

import (
	"fmt"
	"net/http"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type envelope map[string]interface{}

func (app *application) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func (app *application) DeleteUserSQLi(w http.ResponseWriter, r *http.Request) {

	var payload jsonResponse
	id := r.URL.Query().Get("id")

	fmt.Println("El id es:", id)

	app.infoLog.Println(r.URL, id)

	err := app.models.User.DeleteUserSQLi(id)
	if err != nil {
		app.errorLog.Println("Couldn't delete user")
		// send back a response
		payload.Error = true
		payload.Message = "Couldn't delete user"
		return
	}

	// send back a response
	payload.Error = false
	payload.Message = "User deleted correctly"

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {

	var payload jsonResponse

	id := r.URL.Query().Get("id")
	fmt.Println("El id es:", id)

	app.infoLog.Println(r.URL, id)

	err := app.models.User.DeleteUser(id)

	if err != nil {
		// send back a response
		payload.Error = true
		payload.Message = "Couldn't delete user" // sqli case

		err = app.writeJSON(w, http.StatusOK, payload)
		if err != nil {
			app.errorLog.Println(err)
		}

		app.errorLog.Println("Couldn't delete user")

		return
	}

	// send back a response
	payload.Error = false
	payload.Message = "User deleted correctly"

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}

}
