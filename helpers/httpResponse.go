package helpers

import (
	"encoding/json"
	"net/http"
)

/*
Error response object for simple error message response, for example:
"message": "ID was not found"
*/
type HttpErrorMessage struct {
	Error string
}

/*
Error response object for validation errors, for example:
"Error": {
	"Todo": [
		"The Todo field is required"
	]
}
*/
type HttpValidationError struct {
	Error map[string][]string
}

// Error response object for returning errors in controllers
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
