package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/goresttodo/models"
	"github.com/goresttodo/services"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo

	json.NewDecoder(r.Body).Decode(&todo)

	services.AddTodoService(todo)

	responseData := models.ErrorResponse{Message: "New Task Added"}

	json.NewEncoder(w).Encode(responseData)
}
