package config

import (
	"github.com/ameteiko/golang-kit/errors"
)

//
// Configuration parameter validation errors.
//
var (
	ErrConfigParameterIsEmpty = errors.NewError("requested configuration parameter value is empty")
)

//
// Cassandra errors.
//
var (
	ErrConnectionStringIsEmpty           = errors.NewError("database connection string is empty")
	ErrConnectionStringSchemeIsIncorrect = errors.NewError("database connection string protocol is incorrect")
	ErrConnectionStringHostsAreEmpty     = errors.NewError("database hosts are not specified")
	ErrConnectionStringKeyspaceIsEmpty   = errors.NewError("keyspace connection clause cannot be empty")
	ErrConnectionStringPasswordIsEmpty   = errors.NewError("password clause cannot be empty if username was set")
)

//
// URL errors
//
var (
//ErrConnectionStringIsEmpty           = errors.NewError("database connection string is empty")
)
