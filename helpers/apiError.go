package helpers

import (
	"fmt"
)

// Custom error type for throwing API errors
type ApiError struct {
	Code    int
	Message string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// A simple preformated 404 error
func NotFoundError(name string, id int) string {
	return fmt.Sprintf("%s with ID %d was not found", name, id)
}
