package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/goresttodo/data"
	"github.com/goresttodo/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todoData []models.Todo = data.TodoData

	responseDataInst := models.TodoSuccessResponse{Message: "Success", Data: todoData}

	json.NewEncoder(w).Encode(responseDataInst)
}
