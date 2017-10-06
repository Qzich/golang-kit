package errors

import (
	"fmt"
	"net/http"
)

//
// HTTPErrorInfoProvider is an interface that provides information about the HTTP error.
//
type HTTPErrorInfoProvider interface {
	error

	GetHTTPStatus() int
	GetErrorCode() int
	GetErrorMessage() string
}

//
// HTTPError error object.
//
type HTTPError struct {
	status  int    `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//
// NewHTTP400Error returns an instance of the HTTP 400 (Bad request) error.
//
func NewHTTP400Error(code int, message string) HTTPError {
	return HTTPError{http.StatusBadRequest, code, message}
}

//
// NewHTTP403Error returns an instance of the HTTP 403 (Forbidden) error.
//
func NewHTTP403Error(code int, message string) HTTPError {
	return HTTPError{http.StatusForbidden, code, message}
}

//
// NewHTTP500Error returns an instance of the HTTP 500 (Internal server error) error.
//
func NewHTTP500Error(code int, message string) HTTPError {
	return HTTPError{http.StatusInternalServerError, code, message}
}

//
// GetErrorCode returns the error Code.
//
func (e HTTPError) GetErrorCode() int {
	return e.Code
}

//
// GetErrorMessage returns the error Message.
//
func (e HTTPError) GetErrorMessage() string {
	return e.Message
}

//
// GetHTTPStatus returns the HTTP status.
//
func (e HTTPError) GetHTTPStatus() int {
	return e.status
}

//
// Error returns a string representation for the error.
//
func (e HTTPError) Error() string {
	if "" != e.Message && 0 != e.Code {
		return fmt.Sprintf(`{"code": %d, "message": "%s"}`, e.GetErrorCode(), e.GetErrorMessage())
	}

	return ""
}
