package main

import (
	"bomberman/api"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api.Router(r)
	http.Handle("/", r)

	// Start the server on localhost port 8080
	println("Server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
