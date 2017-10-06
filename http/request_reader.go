package http

import (
	"io/ioutil"
	"net/http"

	"github.com/ameteiko/golang-kit/errors"
)

// requestReader interface lists all the actions on the request object.
type requestReader interface {
	readBody(request *http.Request) ([]byte, error)
	readURLParameter(request *http.Request, parameterName string) string
}

//
// requestRead is a request reader object.
//
type requestRead struct{}

//
// readBody returns the request body contents.
//
func (r *requestRead) readBody(request *http.Request) ([]byte, error) {
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, errors.WithMessage(err, `kit-http@requestRead.readBody`)
	}

	return requestBody, nil
}

//
// readURLParameter returns the value for the URL parameter.
//
func (r *requestRead) readURLParameter(request *http.Request, parameterName string) string {

	return request.URL.Query().Get(parameterName)
}
