package models

type TodoSuccessResponse struct {
	Message string `json:"message"`
	Data    []Todo `json:"data"`
}
