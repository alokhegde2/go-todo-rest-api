package services

import (
	"github.com/goresttodo/data"
	"github.com/goresttodo/models"
)

func AddTodoService(value models.Todo) {
	data.TodoData = append(data.TodoData, value)
}
