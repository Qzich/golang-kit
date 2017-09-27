package config

import (
	"net/url"

	"github.com/ameteiko/golang-kit/errors"
	"strings"
)

//
// GetRedisConnectionInfo returns redis connection info object.
//
func (c *Config) GetRedisConnectionInfo() RedisConnectionInfoProvider {
	config := c.parameters[ConfigRedis].(RedisConnectionInfoProvider)

	return config
}

//
// Redis connection URL constants.
//
const (
	RedisConnectionProtocol = "redis"
)

//
// RedisConnectionInfo is a redis connection info parameter.
//
type RedisConnectionInfo struct {
	StringParameter

	host string
	port string
}

//
// RedisConnectionInfoProvider declares all the connection info getters.
//
type RedisConnectionInfoProvider interface {
	ParseConnectionString(string) error

	GetHost() string
	GetPort() string
}

//
// GetHost returns the host.
//
func (c *RedisConnectionInfo) GetHost() string {
	return c.host
}

//
// GetPort returns port value.
//
func (c *RedisConnectionInfo) GetPort() string {
	return c.port
}

//
// validate validates the redis connection string parameter.
//
func (c *RedisConnectionInfo) validate() error {
	var err error
	var connectionInfo *url.URL
	connectionString := c.GetValue()

	if "" == connectionString {
		return ErrRedisConnectionStringIsEmpty
	}

	if connectionInfo, err = url.Parse(connectionString); nil != err {
		return errors.Wrapf(
			err,
			`incorrect redis connection string (%s)`,
			connectionString,
		)
	}

	if _, err := validateConnectionProtocolClause(connectionInfo, RedisConnectionProtocol); nil != err {
		return errors.Wrapf(
			err,
			`incorrect protocol value for connection string (%s)`,
			connectionString,
		)
	}

	if c.host, err = validateRedisHostClause(connectionInfo); nil != err {
		return errors.Wrapf(
			err,
			`incorrect host value for connection string (%s)`,
			connectionString,
		)
	}

	if c.port, err = validateRedisPortClause(connectionInfo); nil != err {
		return errors.Wrapf(
			err,
			`incorrect port value for connection string (%s)`,
			connectionString,
		)
	}

	return nil
}

//
// validateHostClause validates the hosts clause.
//
func validateRedisHostClause(url *url.URL) (string, error) {
	host := strings.Split(url.Host, ":")[0]

	if "" == host {
		return "", ErrRedisHostIsEmpty
	}

	return host, nil
}

//
// validateRedisPortClause validates the hosts clause.
//
func validateRedisPortClause(url *url.URL) (string, error) {
	port := url.Port()
	if "" == port {
		return "", ErrRedisPortIsEmpty
	}

	return port, nil
}
