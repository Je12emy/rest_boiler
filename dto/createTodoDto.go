package dto

type CreateTodo struct {
	Todo string `json:"Todo" validate:"required,alphanumeric"`
}
