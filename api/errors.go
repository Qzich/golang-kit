package api

import (
	"github.com/ameteiko/golang-kit/errors"
)

//
// API package errors.
//
var (
	ErrRequestCreationError         = errors.NewError("net/http request creation error")
	ErrRequestToExternalAPI         = errors.NewError("error on external API call")
	ErrExternalAPIResponseReadError = errors.NewError("error on external API response read")
)
