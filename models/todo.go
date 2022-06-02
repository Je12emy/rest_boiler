package models

const TodoModelName string = "Todo"

type Todo struct {
	ID        uint
	Todo      string
	Completed bool
}
