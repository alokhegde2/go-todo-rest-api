package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Make caps to export the function
func Server() {
	r := mux.NewRouter()

	Todo(r)
	//Running server
	log.Fatal(http.ListenAndServe(":3000", r))
}
