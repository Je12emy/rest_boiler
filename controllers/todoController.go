package controllers

import (
	"net/http"
	"strconv"

	"je12emy/todo_app/controllers/middlewares"
	"je12emy/todo_app/dto"
	"je12emy/todo_app/helpers"
	"je12emy/todo_app/services"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/govalidator"
)

type TodoController struct {
	todoService services.ITodoService
}

func NewTodoController(service services.ITodoService) *TodoController {
	return &TodoController{
		todoService: service,
	}
}

func (t TodoController) RegisterRoutes(r *mux.Router) {
	todoRouter := r.PathPrefix("/todo").Subrouter()

	todoRouter.HandleFunc("/", middlewares.MultipleMiddleware(t.GetAll, middlewares.AddJsonContentType)).Methods(http.MethodGet)
	todoRouter.HandleFunc("/", middlewares.MultipleMiddleware(t.CreateTodo, middlewares.AddJsonContentType)).Methods(http.MethodPost)
	todoRouter.HandleFunc("/{id}", middlewares.MultipleMiddleware(t.GetById, middlewares.AddJsonContentType)).Methods(http.MethodGet)
	todoRouter.HandleFunc("/{id}", middlewares.MultipleMiddleware(t.DeleteById, middlewares.AddJsonContentType)).Methods(http.MethodDelete)
	todoRouter.HandleFunc("/{id}/toggle", middlewares.MultipleMiddleware(t.ToggleTodo, middlewares.AddJsonContentType)).Methods(http.MethodPost)
}

func (t TodoController) GetAll(w http.ResponseWriter, r *http.Request) {
	todos := t.todoService.GetAll()
	helpers.SendResponse(http.StatusOK, todos, w)
}

func (t TodoController) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todo, err := t.todoService.GetById(id)

	if err != nil {
		helpers.SendApiError(err, w)
		return
	}

	helpers.SendResponse(http.StatusOK, todo, w)
}

func (t TodoController) ToggleTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		helpers.ThrowError(http.StatusBadRequest, "The provided Id should have been a number", w)
		return
	}

	todo, err := t.todoService.ToggleTodo(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	helpers.SendResponse(http.StatusOK, todo, w)
}

func (t TodoController) DeleteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		helpers.ThrowError(http.StatusNotFound, "The provided Id should have been a number", w)
		return
	}

	todo, err := t.todoService.DeleteById(id)

	if err != nil {
		helpers.ThrowError(http.StatusNotFound, err.Error, w)
		return
	}

	helpers.SendResponse(http.StatusOK, todo, w)
}

func (t TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"Todo": []string{"required"},
	}

	var createTodo dto.CreateTodo

	// TODO try moving me into a customer middleware maybe or into a helper?
	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		RequiredDefault: true,
		Data:            &createTodo,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()

	if len(e) != 0 {
		// err := map[string]interface{}{"error": e}
		helpers.ThrowError(http.StatusBadRequest, e, w)
		return
	}

	todo, err := t.todoService.CreateTodo(createTodo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	helpers.SendResponse(http.StatusOK, todo, w)
}
