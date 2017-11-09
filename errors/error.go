package errors

import (
	"github.com/ameteiko/errors"
)

//
// ErrorInfoProvider is an Error interface.
//
type ErrorInfoProvider interface {
	error

	GetMessage() string
}

//
// Error is a base application error object.
//
type Error struct {
	message string
}

//
// NewError returns a new Error instance.
// Need to return an object not a link for it because of typed checks.
//
func NewError(message string) error {

	return errors.WrapError(Error{message: message})
}

//
// GetMessage returns an error message.
// This method is needed (despite the fact it is similar to Error() one) to declare a separate interface for a typed
// checks.
//
func (e Error) GetMessage() string {
	return e.message
}

//
// Error returns an error message.
//
func (e Error) Error() string {
	return e.message
}
