package errors

import (
	"fmt"

	"github.com/ameteiko/errors"
)

//
// New returns a new error instance.
//
func New(message string) error {

	return errors.New(message)
}

//
// Errorf returns a new formatted error instance.
//
func Errorf(format string, args ...interface{}) error {

	return errors.Errorf(format, args)
}

//
// WithStack records an error stack trace at the point it was invoked.
//
func WithStack(err error) error {

	return errors.WithStack(err)
}

//
// Wrap wraps an error with a formatted message.
//
func Wrap(err error, format string, args ...interface{}) error {

	return errors.Wrapf(err, format, args...)
}

//
// WrapError wraps an error with provided errors.
//
func WrapError(err error, outerErrors ...error) error {

	return errors.WrapError(err, outerErrors...)
}

//
// WithMessage annotates an error with a new message.
//
func WithMessage(err error, format string, args ...interface{}) error {

	return errors.WithMessage(err, fmt.Sprintf(format, args...))
}

//
// Cause returns the underlying cause of the error, if possible.
//
func Cause(err error, causers ...interface{}) error {

	return errors.Cause(err, causers...)
}
