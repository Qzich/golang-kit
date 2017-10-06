package http

import (
	"encoding/json"
	"net/http"

	"github.com/ameteiko/golang-kit/errors"
)

//
// Responder is a HTTP handler response interface.
//
type Responder interface {
	//
	// GetBody returns response body.
	//
	GetBody() []byte

	//
	// SetBody sets response body.
	//
	SetBody(responseObject interface{}) error

	//
	// GetStatus returns response HTTP status.
	//
	GetStatus() int

	//
	// SetStatus sets response HTTP status.
	//
	SetStatus(int)

	//
	// SetCreatedStatus sets the response status to HTTP 201.
	//
	SetCreatedStatus()

	//
	// SetBadRequestStatus sets the response status to HTTP 404.
	//
	SetBadRequestStatus()
}

//
// Response is an HTTP response object.
//
type Response struct {
	body   []byte
	status int
}

//
// NewResponse creates a new Response instance.
//
func NewResponse() *Response {

	return &Response{status: http.StatusOK}
}

//
// GetBody returns a response body.
//
func (r *Response) GetBody() []byte {

	return r.body
}

//
// SetBody sets a response body.
// The method expects a JSON-annotated response.
//
func (r *Response) SetBody(responseObject interface{}) error {
	var err error

	if r.body, err = json.Marshal(responseObject); nil != err {
		return errors.WrapError(
			errors.WithMessage(err, `unable to encode an HTTP response body`),
			errors.ErrResponseEncodingFailed,
		)
	}

	return nil
}

//
// GetStatus returns a response status.
//
func (r *Response) GetStatus() int {

	return r.status
}

//
// SetStatus sets a response status.
//
func (r *Response) SetStatus(status int) {
	r.status = status
}

//
// SetBadRequestStatus sets a response status.
//
func (r *Response) SetBadRequestStatus() {
	r.status = http.StatusBadRequest
}

//
// SetCreatedStatus sets a response status.
//
func (r *Response) SetCreatedStatus() {
	r.status = http.StatusCreated
}
