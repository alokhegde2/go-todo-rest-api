package models

type Todo struct {
	Id          string `json:"id"`
	Task        string `json:"task"`
	CreatedDate string `json:"createdDate"`
	IsCompleted bool   `json:"isCompleted"`
}
