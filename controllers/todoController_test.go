package controllers

import (
	"encoding/json"
	"io/ioutil"
	"je12emy/todo_app/helpers"
	"je12emy/todo_app/models"
	"je12emy/todo_app/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	todoController TodoController
	recorder       *httptest.ResponseRecorder
	todos          []models.Todo
	router         http.Handler
)

func setup() {
	todoController = *NewTodoController(services.NewTodoService())
	router = initRouter(todoController)
	recorder = httptest.NewRecorder()
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
}

func initRouter(t TodoController) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/todo/{id}", t.GetById).Methods(http.MethodGet)
	r.HandleFunc("/todo", t.GetAll).Methods(http.MethodGet)
	return r
}

func TestGetAll(t *testing.T) {
	// Arrange
	setup()
	var response []models.Todo
	request := httptest.NewRequest(http.MethodGet, "/todo", nil)
	sut := router

	// Act
	sut.ServeHTTP(recorder, request)
	result := recorder.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	err = json.Unmarshal(data, &response)

	// Assert
	if err != nil {
		t.Errorf("Expected error to be nil but got %v", err)
	}

	assert.Equal(t, result.StatusCode, http.StatusOK, "Response should have been 200 Ok")
	assert.Equal(t, response, todos, "Response did not match the expected result")
}

func TestGetById(t *testing.T) {
	// Arrange
	setup()
	var response models.Todo
	request := httptest.NewRequest(http.MethodGet, "/todo/1", nil)
	sut := router

	// Act
	sut.ServeHTTP(recorder, request)
	result := recorder.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	err = json.Unmarshal(data, &response)

	// Assert
	if err != nil {
		t.Errorf("Expected error to be nil but got %v", err)
	}

	assert.Equal(t, result.StatusCode, http.StatusOK, "Response should have been 200 Ok")
	assert.Equal(t, response, todos[1], "Response did not match the expected result")
}

func TestGetByIdNotFound(t *testing.T) {
	// Arrange
	setup()
	type ApiError struct {
		Error string `json:"Error"`
	}

	var response ApiError

	request := httptest.NewRequest(http.MethodGet, "/todo/99", nil)
	sut := router

	// Act
	sut.ServeHTTP(recorder, request)
	result := recorder.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	err = json.Unmarshal(data, &response)

	// Assert
	if err != nil {
		t.Errorf("Expected error to be nil but got %v", err)
	}

	assert.Equal(t, result.StatusCode, http.StatusNotFound, "Response should have been 404 Not Found")
	assert.Equal(t, response.Error, helpers.NotFoundError(models.TodoModelName, 99))
}
