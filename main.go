package main

import (
	"fmt"
	"net/http"

	"je12emy/todo_app/controllers"
	"je12emy/todo_app/services"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Healthy")
	})

	todoController := controllers.NewTodoController(services.NewTodoService())
	todoController.RegisterRoutes(router)

	http.ListenAndServe(":8080", router)
}
