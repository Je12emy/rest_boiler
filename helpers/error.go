package helpers

import (
	"fmt"
)

type ApiError struct {
	Code    int
	Message string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func NotFoundError(name string, id int) string {
	return fmt.Sprintf("%s with ID %d was not found", name, id)
}
