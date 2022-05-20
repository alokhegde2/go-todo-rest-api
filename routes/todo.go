package routes

import (
	"github.com/goresttodo/controllers"
	"github.com/gorilla/mux"
)

func Todo(r *mux.Router) {
	r.HandleFunc("/addTodo", controllers.AddTodo).Methods("POST")
	r.HandleFunc("/getTodos", controllers.GetTodos).Methods("GET")
}
