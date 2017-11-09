package mocks

import (
	"github.com/stretchr/testify/mock"
)

//
// LoggerMock is a mock instance for Logger object.
//
type LoggerMock struct {
	mock.Mock
}

//
// Error method mocking.
//
func (l LoggerMock) Error(format string, args ...interface{}) {

	l.Called(format, args)
}

//
// Debug method mocking.
//
func (l LoggerMock) Debug(format string, args ...interface{}) {

	l.Called(format, args)
}

//
// Info method mocking.
//
func (l LoggerMock) Info(format string, args ...interface{}) {

	l.Called(format, args)
}
