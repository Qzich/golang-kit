package cfg

import (
	"net/url"
	"strings"

	"fmt"
	"github.com/ameteiko/golang-kit/errors"
)

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
	host string
	port string

	*StringParameter
}

//
// RedisConnectionInfoProvider declares all the connection info getters.
//
type RedisConnectionInfoProvider interface {
	GetHost() string
	GetPort() string

	ParameterInfoProvider
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
	errMsg := fmt.Sprintf(`kit-cfg@CassandraConnectionInfo.validate [connection string (%s)]`, connectionString)

	if "" == connectionString {
		return ErrRedisConnectionStringIsEmpty
	}

	if connectionInfo, err = url.Parse(connectionString); nil != err {
		return errors.WrapError(ErrRedisConnectionStringIsIncorrect, errors.WithMessage(err, errMsg))
	}

	if _, err := validateConnectionProtocolClause(connectionInfo, RedisConnectionProtocol); nil != err {
		return errors.WrapError(ErrRedisProtocolIsIncorrect, errors.WithMessage(err, errMsg))
	}

	if c.host, err = validateRedisHostClause(connectionInfo); nil != err {
		return errors.WithMessage(err, errMsg)
	}

	if c.port, err = validateRedisPortClause(connectionInfo); nil != err {
		return errors.WithMessage(err, errMsg)
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
