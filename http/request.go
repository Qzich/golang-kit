package http

import (
	"context"
)

//
// RequestHandler is an interface for all HTTP handlers.
//
type RequestHandler interface {
	//
	// Handle handler an HTTP request.
	//
	Handle([]byte, context.Context, *Responder) error
}
