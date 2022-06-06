package helpers

import (
	"encoding/json"
	"net/http"
)

// Error response object for HTTP requests, meant to be used on tests
type HttpErrorMessage struct {
	Error string
}

// Error response object for validation errors
type HttpErrorObject struct {
	Error map[string][]string
}

// Error response object
type errorResponse struct {
	Error interface{}
}

// Send any kind of error to the ResponseWriter
func ThrowError(code int, error interface{}, w http.ResponseWriter) {
	message := errorResponse{
		Error: error,
	}
	SendResponse(code, message, w)
}

// Send a response to the ResponseWriter
func SendResponse(code int, data interface{}, w http.ResponseWriter) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// Check for an API Error and throw it
func SendApiError(error error, w http.ResponseWriter) {
	// If we can assert the error into an ApiError send it
	if err, ok := error.(*ApiError); ok {
		ThrowError(err.Code, err.Message, w)
		return
	}
	// If it's unkown send it with a 500 status code
	ThrowError(http.StatusInternalServerError, error.Error(), w)
	return
}
