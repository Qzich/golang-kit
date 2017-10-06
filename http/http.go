package http

import (
	"net/http"

	"github.com/ameteiko/golang-kit/errors"
)

//
// WriteResponseError writes the information about the error to the response.
//
func WriteResponseError(responseWriter http.ResponseWriter, e errors.HTTPErrorInfoProvider) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(e.GetHTTPStatus())
	responseWriter.Write([]byte(e.Error()))
}
