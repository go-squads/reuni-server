package helper

import (
	"fmt"
)

type HttpError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewHttpError(status int, message string) *HttpError {
	return &HttpError{
		Message: message,
		Status:  status,
	}
}

func (e HttpError) Error() string {
	return fmt.Sprintf(`{"status:": "%v",  "message": "%v"}`, e.Status, e.Message)
}
