package helpers

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error interface{}
}

func ThrowError(code int, error interface{}, w http.ResponseWriter) {
	message := errorResponse{
		Error: error,
	}
	SendResponse(code, message, w)
}

func SendResponse(code int, data interface{}, w http.ResponseWriter) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
