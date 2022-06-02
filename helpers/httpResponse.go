package helpers

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	Error interface{}
}

func ThrowError(code int, error interface{}, w http.ResponseWriter) {
	message := httpResponse{
		Error: error,
	}
	SendResponse(code, message, w)
}

func SendResponse(code int, data interface{}, w http.ResponseWriter) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func SendError(error error, w http.ResponseWriter) {
	// If we can assert the error into an ApiError send it
	if err, ok := error.(*ApiError); ok {
		ThrowError(err.Code, err.Message, w)
		return
	}
	// If it's unkown send it with a 500 status code
	ThrowError(http.StatusInternalServerError, error.Error(), w)
	return
}
