package errors

import (
	"fmt"
	"net/http"

)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

// TODO clean up mess with ApplicationError and ApiError!!
type ApplicationError struct {
	AMessage    string `json:"message"`
	AStatusCode int    `json:"status"`
	ACode       string `json:"code"`
}

func (e *ApplicationError) Status() int {
	return e.AStatusCode
}


func (e *ApplicationError) Message() string {
	return e.AMessage
}


func (e *ApplicationError) Error() string {
	return e.ACode
}

func (e *ApplicationError) String() string {
	return fmt.Sprintf( "ApplicationError: %s Status: %s Code: %s", e.AMessage, e.AStatusCode, e.ACode)
}

func NewBadRequestError(message string) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage: message,
		AStatusCode: http.StatusBadRequest,
		ACode: "bad request",
	}
	return apiErr
}

func NewInternalServerError(message string) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage: message,
		AStatusCode: http.StatusInternalServerError, //500
		ACode: "internal server error",
	}
	return apiErr
}

func NewNotFoundError(message string) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage: message,
		AStatusCode: http.StatusNotFound, //404
		ACode: "not found",
	}
	return apiErr
}

func NewNotYetImplementedError(message string) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage: message,
		AStatusCode: http.StatusNotImplemented, //501
		ACode: "not yet implemented",
	}
	return apiErr
}

func NewTokenValidationError(message string) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage: message,
		AStatusCode: http.StatusUnauthorized, //401
		ACode: "token validation error",
	}
	return apiErr
}

func NewTokenExpiredError(message string) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage: message,
		AStatusCode: http.StatusUnauthorized, //401
		ACode: "token already expired",
	}
	return apiErr
}