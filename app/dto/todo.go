package dto

import "go-todo-app/app/models"

type TodosResponseDto struct {
	Data []models.Todo `json:"data"`
}
