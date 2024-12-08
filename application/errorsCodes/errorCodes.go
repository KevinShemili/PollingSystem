package errorsCodes

import "net/http"

type ErrorCode struct {
	Name        string
	Code        int
	Description string
}

func (e *ErrorCode) GetName() string {
	return e.Name
}

func (e *ErrorCode) GetCode() int {
	return e.Code
}

func (e *ErrorCode) GetDescription() string {
	return e.Description
}

var DATABASE_CONNECTION_ERROR = ErrorCode{
	Name: "Database connection error.",
	Code: http.StatusInternalServerError}

var DATABASE_NOT_CREATED = ErrorCode{
	Name:        "Database could not be created.",
	Code:        http.StatusInternalServerError,
	Description: "Database did not exist, attemp to create it failed."}
