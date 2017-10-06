package errors

import (
	"net/http"
)

//
// Predefined application errors.
//
var (
	ErrGetMisregisteredConfigParameter = Error{
		"trying to get configuration parameter that was not registered properly",
	}

	ErrNotFound    = HTTPError{status: http.StatusNotFound}
	ErrRequestRead = NewHTTP400Error(
		30000,
		"Request content reading error. Check your request body.",
	)
	ErrRequestParsing = NewHTTP400Error(
		30001,
		"Request body parsing error. It must be a valid JSON object.",
	)
	ErrResponseEncodingFailed = NewHTTP400Error(
		30002,
		"Response encoding error. The request itself succeeded.",
	)

	ErrInternalServerError = NewHTTP500Error(
		10000,
		"Request serving internal error. Try again later.",
	)
)
