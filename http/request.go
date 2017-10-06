package http

import (
	"net/http"
)

//
// RequestHandler is an interface for all HTTP handlers.
//
type RequestHandler interface {
	//
	// Handle handler an HTTP request.
	//
	Handle([]byte, Responder, *http.Request) error
}
