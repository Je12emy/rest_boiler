package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"je12emy/todo_app/dto"
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
	r.HandleFunc("/todo", t.CreateTodo).Methods(http.MethodPost)
	return r
}

func Test_TodoController_Returns_All_Todos(t *testing.T) {
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

func Test_TodoController_Returns_A_Todo_By_Its_ID(t *testing.T) {
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

func Test_TodoController_Returns_NotFound_When_The_ID_Is_Not_Found(t *testing.T) {
	// Arrange
	setup()
	var response helpers.HttpErrorMessage

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

func Test_TodoController_Returns_200Ok_When_Creating_A_Todo(t *testing.T) {
	// Arrange
	setup()
	var response models.Todo
	var createTodoDto = dto.CreateTodo{
		Todo: "Buy milk",
	}

	var requestBody bytes.Buffer
	err := json.NewEncoder(&requestBody).Encode(createTodoDto)

	if err != nil {
		t.Errorf("Expected error to be nil but got %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/todo", &requestBody)
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
	assert.Equal(t, response.Todo, createTodoDto.Todo, "Response Todo did not match with the request")
	assert.Equal(t, response.Completed, false, "Response Todo did not match with the request")
}

func Test_TodoController_Returns_BadRequest_When_The_Body_Is_Empty(t *testing.T) {
	// Arrange
	setup()
	var response helpers.HttpErrorObject
	var createTodoDto string = "{}"
	var requestBody bytes.Buffer

	err := json.NewEncoder(&requestBody).Encode(createTodoDto)
	if err != nil {
		t.Errorf("Expected error to be nil but got %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/todo", nil)
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

	assert.Equal(t, result.StatusCode, http.StatusBadRequest, "Response should have been 400 BadRequest")
	assert.Equal(t, "The Todo field is required", response.Error["Todo"][0], "Response should have been 400 BadRequest")
}
