package services

import (
	"je12emy/todo_app/dto"
	"je12emy/todo_app/models"
)

type ITodoService interface {
	GetAll() []models.Todo
	GetById(id int) (models.Todo, error)
	ToggleTodo(id int) (models.Todo, error)
	DeleteById(int int) (models.Todo, error)
	CreateTodo(createTodo dto.CreateTodo) (models.Todo, error)
}
