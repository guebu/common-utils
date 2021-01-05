package errors

import (
	"fmt"
	"net/http"

)

type ApiError interface {
	Status() int
	Message() string
	Error() string
	Causes() []interface{}
}

// TODO clean up mess with ApplicationError and ApiError!!
type ApplicationError struct {
	AMessage    string `json:"message"`
	AStatusCode int    `json:"status"`
	AnError     string `json:"code"`
	SomeCauses 	[]interface{} `json:"causes"`
}

func (e *ApplicationError) Status() int {
	return e.AStatusCode
}

func (e *ApplicationError) Message() string {
	return e.AMessage
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("Message: %s - Status-Code: %d - Error: %s - Causes:  [ %v ]", e.AMessage, e.AStatusCode, e.AnError, e.SomeCauses)
}

func (e *ApplicationError) String() string {
	return fmt.Sprintf(   "ApplicationError: %s Status: %s Code: %s Causes: [ %v ]", e.AMessage, e.AStatusCode, e.AnError, e.SomeCauses)
}

func (e *ApplicationError) Causes() []interface{} {
	return e.SomeCauses
}

func NewBadRequestError(message string, err error) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage:    message,
		AStatusCode: http.StatusBadRequest,
		AnError:     "bad request",
	}

	if err != nil {
		apiErr.SomeCauses = append(apiErr.SomeCauses, err.Error() )
	}

	return apiErr
}

func NewInternalServerError(message string, err error) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage:    message,
		AStatusCode: http.StatusInternalServerError, //500
		AnError:     "internal server error",
	}

	if err != nil {
		apiErr.SomeCauses = append(apiErr.SomeCauses, err.Error() )
	}

	return apiErr
}

func NewNotFoundError(message string, err error) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage:    message,
		AStatusCode: http.StatusNotFound, //404
		AnError:     "not found",
	}

	if err != nil {
		apiErr.SomeCauses = append(apiErr.SomeCauses, err.Error() )
	}

	return apiErr
}

func NewNotYetImplementedError(message string) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage:    message,
		AStatusCode: http.StatusNotImplemented, //501
		AnError:     "not yet implemented",
	}
	return apiErr
}

func NewTokenValidationError(message string, err error) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage:    message,
		AStatusCode: http.StatusUnauthorized, //401
		AnError:     "token validation error",
	}

	if err != nil {
		apiErr.SomeCauses = append(apiErr.SomeCauses, err.Error() )
	}

	return apiErr
}

func NewTokenExpiredError(message string, err error) *ApplicationError {
	apiErr := &ApplicationError{
		AMessage:    message,
		AStatusCode: http.StatusUnauthorized, //401
		AnError:     "token already expired",
	}

	if err != nil {
		apiErr.SomeCauses = append(apiErr.SomeCauses, err.Error() )
	}

	return apiErr
}