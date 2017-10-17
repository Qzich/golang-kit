package helper

import (
	"fmt"
	"testing"

	baseErrors "github.com/ameteiko/errors"
	"github.com/ameteiko/golang-kit/errors"
)

//
// AssertError is a utility function to check is any of errors inside the error stack passed as an err parameter
// contains the error of type expectedError.
//
func AssertError(t *testing.T, expectedError error, err error) bool {
	typedError := baseErrors.Cause(err, (*errors.ErrorInfoProvider)(nil))
	if nil == typedError || typedError.Error() != expectedError.Error() {
		msg := fmt.Sprintf("Error type is not as expected.\n"+
			"expected: %s\n"+
			"actual: %s",
			expectedError, err)
		return fail(t, msg)
	}

	return true
}

//
// AssertHTTPError is a utility function to check is any of errors inside the error stack passed as an err parameter
// contains the error of type expectedError.
//
func AssertHTTPError(t *testing.T, expectedError error, err error) bool {
	typedError := baseErrors.Cause(err, (*errors.HTTPErrorInfoProvider)(nil))
	if nil == typedError || typedError.Error() != expectedError.Error() {
		msg := fmt.Sprintf("Error type is not as expected.\n"+
			"expected: %s\n"+
			"actual: %s",
			expectedError, err)
		return fail(t, msg)
	}

	return true
}

//
// fail fails the test execution.
//
func fail(t *testing.T, message string) bool {
	t.Errorf(message)

	return false
}
