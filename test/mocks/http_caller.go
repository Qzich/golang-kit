package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ameteiko/golang-kit/api"
)

//
// HTTPCallerMock is a mock instance for HTTP object.
//
type HTTPCallerMock struct {
	mock.Mock
}

//
// Call method mocking.
//
func (c HTTPCallerMock) Call(httpResource api.HTTPResourceInfoProvider, response interface{}) error {

	args := c.Called(httpResource, response)

	return args.Error(0)
}
