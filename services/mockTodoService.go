package services

import (
	"errors"
	"fmt"
	"je12emy/todo_app/dto"
	"je12emy/todo_app/models"
)

type TodoService struct{}

var todos []models.Todo

func NewTodoService() *TodoService {
	todos = []models.Todo{
		{
			ID:        0,
			Todo:      "Buy milk",
			Completed: false,
		},
		{
			ID:        1,
			Todo:      "Buy cheese",
			Completed: false,
		},
		{
			ID:        2,
			Todo:      "Buy eggs",
			Completed: false,
		},
	}
	return &TodoService{}
}

func (t TodoService) GetAll() []models.Todo {
	return todos
}

func (t TodoService) GetById(id int) (models.Todo, error) {
	if id > len(todos) {
		return models.Todo{}, errors.New(fmt.Sprintf("Todo with id of %d not found", id))
	}
	todo := todos[id]
	return todo, nil
}

func (t TodoService) ToggleTodo(id int) (models.Todo, error) {
	if id > len(todos) {
		return models.Todo{}, errors.New(fmt.Sprintf("Todo with id of %d not found", id))
	}

	todo := &todos[id]
	todo.Completed = !todo.Completed
	return *todo, nil
}

func (t TodoService) DeleteById(id int) (models.Todo, error) {
	if id > len(todos) {
		return models.Todo{}, errors.New(fmt.Sprintf("Todo with id of %d not found", id))
	}

	todo := todos[id]
	// All items after id + 1
	todos = append(todos[:id], todos[id+1:]...)

	return todo, nil
}

func (t TodoService) CreateTodo(createTodo dto.CreateTodo) (models.Todo, error) {
	todo := models.Todo{
		ID:        uint(len(todos)),
		Todo:      createTodo.Todo,
		Completed: false,
	}

	todos = append(todos, todo)

	return todo, nil
}
