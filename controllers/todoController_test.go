package controllers

import (
	"encoding/json"
	"io/ioutil"
	"je12emy/todo_app/models"
	"je12emy/todo_app/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	todoController TodoController
	recorder       *httptest.ResponseRecorder
	todos          []models.Todo
)

func setup() {
	todoController = *NewTodoController(services.NewTodoService())
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

func TestGetAll(t *testing.T) {
	// Arrange
	setup()
	request := httptest.NewRequest(http.MethodGet, "/todo", nil)
	var response []models.Todo

	// Act
	todoController.GetAll(recorder, request)
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

func TestGetById(t *testing.T)  {
	// Arrange
	setup()
	request := httptest.NewRequest(http.MethodGet, "/todo/1", nil)
	var response models.Todo

	// Act
	todoController.GetById(recorder, request)
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
