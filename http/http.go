package http

import (
	"net/http"

	"github.com/ameteiko/golang-kit/errors"
)

//
// WriteResponseError writes the information about the error to the response.
//
func WriteResponseError(responseWriter http.ResponseWriter, err error) {
	httpError, ok := err.(errors.HTTPErrorInfoProvider)
	if ok {
		responseWriter.WriteHeader(httpError.GetHTTPStatus())
	} else {
		responseWriter.WriteHeader(http.StatusBadRequest)
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write([]byte(err.Error()))
}
