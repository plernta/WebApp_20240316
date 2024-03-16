package main

import (
	"net/http"
	"welldone/hong/cool_web_app/handlers"

	"fmt"
	"log"
	"os"

	"welldone/hong/cool_web_app/database"

	"github.com/gorilla/mux"
)

func main() {
	// connect to the database
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := database.Connect(connString)
	if err != nil {
		// print error and exit
		log.Fatal(err)
	}

	//create a new app object
	app := handlers.NewApp(db)

	//db.Query("SELECT * FROM users")

	// router is making choices base on the url which handler to call
	r := mux.NewRouter()

	// connect URL to the handler
	r.HandleFunc("/", app.IndexPage)
	r.HandleFunc("/users/{username}", app.UserPage)

	log.Printf("Server started on port 3000")

	// start the server
	http.ListenAndServe(":3000", r)
	//:3000 is the port number

	//http.ListenAndServe(":3000/about", r)

}
