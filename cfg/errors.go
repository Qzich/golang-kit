package cfg

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
	ErrCassandraConnectionStringIsEmpty     = errors.NewError("cassandra connection string is empty")
	ErrCassandraConnectionStringIsIncorrect = errors.NewError("cassandra connection string is incorrect")
	ErrCassandraProtocolIsIncorrect         = errors.NewError("cassandra connection protocol is incorrect")
	ErrCassandraHostsAreEmpty               = errors.NewError("cassandra hosts are empty")
	ErrCassandraKeyspaceIsEmpty             = errors.NewError("cassandra cassandraKeyspace is empty")
	ErrCassandraPasswordIsEmpty             = errors.NewError("cassandra password cannot be empty if username was set")
)

//
// Redis errors.
//
var (
	ErrRedisConnectionStringIsEmpty     = errors.NewError("redis connection string is empty")
	ErrRedisConnectionStringIsIncorrect = errors.NewError("redis connection string is incorrect")
	ErrRedisProtocolIsIncorrect         = errors.NewError("redis connection protocol is incorrect")
	ErrRedisHostIsEmpty                 = errors.NewError("redis host is empty")
	ErrRedisPortIsEmpty                 = errors.NewError("redis port is empty")
)

//
// URL errors.
//
var (
	ErrURLIncorrectValue    = errors.NewError("URL configuration parameter is incorrect")
	ErrBase64IncorrectValue = errors.NewError("parameter is not a base64-encoded string")
)

//
// Log errors.
//
var (
	ErrLogSeverityIncorrectValue = errors.NewError("log severity is incorrect")
)
