package utility

import "net/http"

type ErrorCode struct {
	Message     string `json:"message"`
	StatusCode  int    `json:"status_code"`
	Description string `json:"description"`
}

func NewFullErrorCode(message string, statusCode int, description string) *ErrorCode {
	return &ErrorCode{
		Message:     message,
		StatusCode:  statusCode,
		Description: description,
	}
}

func NewErrorCode(message string, statusCode int) *ErrorCode {
	return &ErrorCode{
		Message:    message,
		StatusCode: statusCode,
	}
}

func (e *ErrorCode) WithDescription(description string) *ErrorCode {
	return &ErrorCode{
		Message:     e.Message,
		StatusCode:  e.StatusCode,
		Description: description,
	}
}

var (
	DatabaseConnectionError = NewErrorCode(
		"Unable to connect to the database.",
		http.StatusInternalServerError,
	)

	DatabaseCreationError = NewErrorCode(
		"Database could not be created.",
		http.StatusInternalServerError,
	)

	IncorrectEmail = NewErrorCode(
		"Incorrect email or password.",
		http.StatusNotFound,
	)

	InternalServerError = NewErrorCode(
		"Internal Server Error.",
		http.StatusInternalServerError,
	)

	BindFailure = NewErrorCode(
		"Binding Failure.",
		http.StatusInternalServerError,
	)

	InvalidToken = NewErrorCode(
		"Invalid Token Credentials.",
		http.StatusUnauthorized,
	)
)
