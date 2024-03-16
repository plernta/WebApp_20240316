package handlers

import (
	"net/http"

	//print on terminal
	"database/sql"
	"errors"

	"github.com/gorilla/mux"
)

func (app *App) IndexPage(w http.ResponseWriter, r *http.Request) {
	//open index.html template and send its content to the client
	renderTemplate(w, "index.html", nil)
}

type User struct {
	Id       int
	Username string
	Email    string
}

func (app *App) UserPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	// query the database for the user using SQL language
	row := app.DB.QueryRow("SELECT id, username, email FROM users WHERE username = $1", username)
	// create an empty user object
	var user User // try to fill the user object with the data from the row
	err := row.Scan(&user.Id, &user.Username, &user.Email)
	//check if the user was not found
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		http.NotFound(w, r)
		return
	}
	//check if there was a real error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "user.html", map[string]any{
		"User": user,
	})
}
