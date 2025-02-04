/*
Custom Error Implementation:
- Consistent error format across API
- Separates technical details from user messages
- Reusable error definitions

Client receives:
- Message: Client friendly error message
- StatusCode: HTTP status code
- Description: Optional detailed error info
*/

package utility

import "net/http"

type ErrorCode struct {
	Message     string `json:"message"`
	StatusCode  int    `json:"status_code"`
	Description string `json:"description"`
}

// NewFullErrorCode creates a new ErrorCode with a message, status code, and description
func NewFullErrorCode(message string, statusCode int, description string) *ErrorCode {
	return &ErrorCode{
		Message:     message,
		StatusCode:  statusCode,
		Description: description,
	}
}

// NewErrorCode creates a new ErrorCode with a message and status code
func NewErrorCode(message string, statusCode int) *ErrorCode {
	return &ErrorCode{
		Message:    message,
		StatusCode: statusCode,
	}
}

// WithDescription adds a description to an ErrorCode
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
		"Invalid Token.",
		http.StatusUnauthorized,
	)

	Unauthorized = NewErrorCode(
		"Unauthorized",
		http.StatusUnauthorized,
	)

	DuplicateEmail = NewErrorCode(
		"User already exists.",
		http.StatusConflict,
	)

	EmailFormat = NewErrorCode(
		"Invalid Email Format.",
		http.StatusBadRequest,
	)

	PasswordFormat = NewErrorCode(
		"Invalid Password.",
		http.StatusBadRequest,
	)

	InvalidPollID = NewErrorCode(
		"Poll does not exist.",
		http.StatusBadRequest,
	)

	InvalidCategoryID = NewErrorCode(
		"The poll does not contain the given category.",
		http.StatusBadRequest,
	)

	PollExpired = NewErrorCode(
		"Unable to add new vote as poll has finished.",
		http.StatusBadRequest,
	)

	AlreadyVoted = NewErrorCode(
		"You have already voted in this poll.",
		http.StatusBadRequest,
	)

	RouteParameterCast = NewErrorCode(
		"Invalid route parameter format",
		http.StatusBadRequest,
	)

	QueryParameterCast = NewErrorCode(
		"Invalid query parameter format",
		http.StatusBadRequest,
	)

	AlreadyEnded = NewErrorCode(
		"Cannot perform operations on an ended poll.",
		http.StatusBadRequest,
	)

	NotPollOwner = NewErrorCode(
		"Only the poll owner can perform this operation.",
		http.StatusBadRequest,
	)

	ValidationError = NewErrorCode(
		"Validation Error.",
		http.StatusBadRequest,
	)
	DateShouldBeFuture = NewErrorCode(
		"Date should be in the future.",
		http.StatusBadRequest,
	)
)
