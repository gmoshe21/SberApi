package models

type NewTask struct {
	UUID			string `json:"uuid" validate:"required"`
	Title     		string `json:"title" validate:"required"`
	Description     string `json:"description" validate:"required"`
	Data 			string `json:"data" validate:"required"`
}

type Task struct {
	UUID			string `json:"uuid" validate:"required"`
}