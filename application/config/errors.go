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
	ErrCassandraConnectionStringIsEmpty = errors.NewError("cassandra connection string is empty")
	ErrCassandraProtocolIsIncorrect     = errors.NewError("cassandra connection protocol is incorrect")
	ErrCassandraHostsAreEmpty           = errors.NewError("cassandra hosts are empty")
	ErrCassandraKeyspaceIsEmpty         = errors.NewError("cassandra cassandraKeyspace is empty")
	ErrCassandraPasswordIsEmpty         = errors.NewError("cassandra cassandraPassword cannot be empty if cassandraUsername was set")
)

//
// Redis errors.
//
var (
	ErrRedisConnectionStringIsEmpty = errors.NewError("redis connection string is empty")
	ErrRedisHostIsEmpty             = errors.NewError("redis host is empty")
	ErrRedisPortIsEmpty             = errors.NewError("redis port is empty")
)

//
// URL errors
//
var (
//ErrConnectionStringIsEmpty           = errors.NewError("database connection string is empty")
)
