//
// Package cfg provides a syntax sugar for the the application configuration parameters access.
//
package cfg

import (
	"os"
	"time"

	"github.com/namsral/flag"

	"github.com/ameteiko/golang-kit/errors"
)

//
// Parameter is a custom configuration parameter type to be used as configuration parameter options enum.
//
type Parameter string

//
// Configuration constants.
//
const (
	DefaultHTTPReadTimeout  = time.Second * 1
	DefaultHTTPWriteTimeout = time.Second * 1
)

//
// Predefined configuration parameter types.
//
const (
	Cassandra          Parameter = "CASSANDRA"
	HTTPAddress                  = "HTTP_ADDRESS"
	DevPortalURL                 = "URL_DEV_PORTAL"
	Redis                        = "REDIS"
	PrivateKey                   = "PRIVATE_KEY"
	PrivateKeyPassword           = "PRIVATE_KEY_PASSWORD"
	PublicKey                    = "PUBLIC_KEY"
	LogSeverity                  = "LOG_SEVERITY"
)

//
// Configer interface provides a base interface for application configuration parameters.
// It allows parameter registration and its value retrieval.
//
type Configer interface {
	//
	// Register registers a predefined configuration parameter for the application.
	//
	Register(Parameter)

	//
	// GetValue returns a string parameter value.
	//
	GetValue(Parameter) string
}

//
// CustomParamsConfiger provides an interface for custom parameters registration.
//
type CustomParamsConfiger interface {
	//
	// RegisterStringParameter registers a custom string configuration parameter.
	//
	RegisterStringParameter(Parameter)

	//
	// GetStringParameter returns a parameter info.
	//
	GetStringParameter(Parameter) (ParameterInfoProvider, error)

	//
	// RegisterBase64Parameter registers a custom base64-encoded configuration parameter.
	//
	RegisterBase64Parameter(Parameter)

	//
	// GetBase64Parameter returns a Base64-encoded parameter info.
	//
	GetBase64Parameter(Parameter) (Base64StringInfoProvider, error)

	//
	// RegisterURLParameter registers a custom URL configuration parameter.
	//
	RegisterURLParameter(Parameter)

	//
	// GetURLParameter returns a URL parameter info.
	//
	GetURLParameter(Parameter) (URLInfoProvider, error)
}

//
// DefaultsConfiger returns all the default application parameters.
//
type DefaultsConfiger interface {
	//
	// GetHTTPReadTimeout a default HTTP read timeout.
	//
	GetHTTPReadTimeout() time.Duration

	//
	// GetHTTPWriteTimeout a default HTTP write timeout.
	//
	GetHTTPWriteTimeout() time.Duration
}

//
// ConfigParametersRetriever interface provides all syntax sugar for a configuration parameters registering and
// retrieving. It's deliberately segregated into separate interface.
//
type ConfigParametersRetriever interface {
	//
	// GetCassandraConnectionInfo returns a Cassandra connection info.
	//
	GetCassandraConnectionInfo() CassandraConnectionInfoProvider

	//
	// GetRedisConnectionInfo returns a Redis connection info.
	//
	GetRedisConnectionInfo() RedisConnectionInfoProvider

	//
	// GetPrivateKeyInfo returns a private key info.
	//
	GetPrivateKeyInfo() Base64StringInfoProvider

	//
	// GetPrivateKeyPasswordInfo returns a private key password info.
	//
	GetPrivateKeyPasswordInfo() Base64StringInfoProvider

	//
	// GetLogInfo returns a log severity info.
	//
	GetLogInfo() LogInfoProvider
}

//
// Config object provides an access for all the application configuration parameters.
//
type Config struct {
	parameters       map[Parameter]ParameterInfoProvider
	httpReadTimeout  time.Duration
	httpWriteTimeout time.Duration
}

//
// NewConfig returns an instance of Config object
//
func NewConfig() *Config {
	config := Config{
		httpReadTimeout:  DefaultHTTPReadTimeout,
		httpWriteTimeout: DefaultHTTPWriteTimeout,
	}
	config.parameters = make(map[Parameter]ParameterInfoProvider)

	return &config
}

//
// Register registers a configuration parameter for the application.
//
func (c *Config) Register(param Parameter) {
	c.parameters[param] = getConfigParameterEntry(param)
}

//
// RegisterStringParameter registers a custom string configuration parameter.
//
func (c *Config) RegisterStringParameter(param Parameter) {
	c.parameters[param] = newStringParameter(param)
}

//
// RegisterBase64Parameter registers a custom base64-encoded configuration parameter.
//
func (c *Config) RegisterBase64Parameter(param Parameter) {
	c.parameters[param] = newBase64Parameter(param)
}

//
// RegisterURLParameter registers a custom URL configuration parameter.
//
func (c *Config) RegisterURLParameter(param Parameter) {
	c.parameters[param] = newURLParameter(param)
}

//
// Parse parses all application configuration parameters.
//
func (c *Config) Parse() error {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	for param, paramEntry := range c.parameters {
		func(parameterValue *string) {
			paramName := string(param)
			flags.StringVar(parameterValue, paramName, "", "")
		}(paramEntry.GetValueLink())
	}
	flags.Parse(os.Args[1:])
	if err := c.validateParameters(); nil != err {
		return errors.WithMessage(err, "kit-cfg@Config.Parse")
	}

	return nil
}

//
// GetValue returns a parameter value by its name.
//
func (c *Config) GetValue(parameter Parameter) string {
	parameterEntry, ok := c.parameters[parameter]
	if !ok {

		return ""
	}

	return parameterEntry.GetValue()
}

//
// GetHTTPReadTimeout a default HTTP read timeout.
//
func (c *Config) GetHTTPReadTimeout() time.Duration {

	return c.httpReadTimeout
}

//
// GetHTTPWriteTimeout a default HTTP write timeout.
//
func (c *Config) GetHTTPWriteTimeout() time.Duration {

	return c.httpWriteTimeout
}

//
// GetCassandraConnectionInfo returns cassandra connection info object.
//
func (c *Config) GetCassandraConnectionInfo() CassandraConnectionInfoProvider {

	return c.parameters[Cassandra].(CassandraConnectionInfoProvider)
}

//
// GetRedisConnectionInfo returns redis connection info object.
//
func (c *Config) GetRedisConnectionInfo() RedisConnectionInfoProvider {

	return c.parameters[Redis].(RedisConnectionInfoProvider)
}

//
// GetPrivateKeyInfo returns a private key info.
//
func (c *Config) GetPrivateKeyInfo() Base64StringInfoProvider {
	privateKeyParameter, _ := c.GetBase64Parameter(PrivateKey)

	return privateKeyParameter
}

//
// GetPrivateKeyPasswordInfo returns a private key password info.
//
func (c *Config) GetPrivateKeyPasswordInfo() Base64StringInfoProvider {
	privateKeyPasswordParameter, _ := c.GetBase64Parameter(PrivateKeyPassword)

	return privateKeyPasswordParameter
}

//
// GetLogInfo returns a log info.
//
func (c *Config) GetLogInfo() LogInfoProvider {

	return c.parameters[LogSeverity].(LogInfoProvider)
}

//
// GetBase64Parameter returns a Base64-encoded parameter info.
//
func (c *Config) GetStringParameter(param Parameter) (ParameterInfoProvider, error) {
	parameter, ok := c.parameters[param].(ParameterInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetStringParameter",
		)
	}

	return parameter, nil
}

//
// GetBase64Parameter returns a Base64-encoded parameter info.
//
func (c *Config) GetBase64Parameter(param Parameter) (Base64StringInfoProvider, error) {
	parameter, ok := c.parameters[param].(Base64StringInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetBase64Parameter",
		)
	}

	return parameter, nil
}

//
// GetURLParameter returns a URL parameter info.
//
func (c *Config) GetURLParameter(param Parameter) (URLInfoProvider, error) {
	parameter, ok := c.parameters[param].(URLInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetURLParameter",
		)
	}

	return parameter, nil
}

//
// validateParameters validates all registered parameters.
//
func (c *Config) validateParameters() error {
	for _, parameterEntry := range c.parameters {
		if err := parameterEntry.validate(); nil != err {
			//println(parameterEntry.GetName())
			//println(parameterEntry.GetValue())
			//println(err.Error())
			//os.Exit(2)
			return errors.WithMessage(
				err,
				`kit-cfg@Config.validateParameters [parameter (%s) value (%s)]`,
				parameterEntry.GetName(), parameterEntry.GetValue(),
			)
		}
	}

	return nil
}

//
// getConfigParameterEntry is a config parameter entry factory.
//
func getConfigParameterEntry(parameter Parameter) ParameterInfoProvider {
	switch parameter {
	case Cassandra:
		return &CassandraConnectionInfo{StringParameter: newStringParameter(parameter)}
	case Redis:
		return &RedisConnectionInfo{StringParameter: newStringParameter(parameter)}
	case DevPortalURL:
		return newURLParameter(parameter)
	case PrivateKey:
		fallthrough
	case PrivateKeyPassword:
		return newBase64Parameter(parameter)
	case LogSeverity:
		return &LogInfo{StringParameter: newStringParameter(parameter)}
	default:
		return newStringParameter(parameter)
	}

	return nil
}

//
// newStringParameter returns a new instance of the string parameter.
//
func newStringParameter(param Parameter) *StringParameter {

	return &StringParameter{name: string(param)}
}

//
// newBase64Parameter returns a new instance of the base64-encoded parameter.
//
func newBase64Parameter(param Parameter) *Base64StringInfo {

	return &Base64StringInfo{StringParameter: newStringParameter(param)}
}

//
// newURLParameter returns a new instance of the URL parameter.
//
func newURLParameter(param Parameter) *URLInfo {

	return &URLInfo{StringParameter: newStringParameter(param)}
}
