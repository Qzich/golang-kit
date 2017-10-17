package jwt

import (
	"github.com/ameteiko/golang-kit/errors"
)

var (
	ErrSignatureDecode = errors.NewError("signature decoding error")
)
