//
// Package config provides a syntax sugar for the the application configuration parameters access.
//
package config

import (
	"os"

	"github.com/namsral/flag"
)

//
// Parameter is a custom configuration parameter type to be used as configuration parameter options enum.
//
type Parameter int

//
// Predefined configuration parameter types.
//
const (
	ConfigCassandra Parameter = iota
	ConfigHTTPPort
	ConfigDevPortalURL
	ConfigRedis
	ConfigPrivateKey
	ConfigPrivateKeyPassword
)

//
// EnvParameters are the predefined configuration parameter types names.
//
var EnvParameters = map[Parameter]string{
	ConfigCassandra:          "CASSANDRA",
	ConfigHTTPPort:           "HTTP_PORT",
	ConfigDevPortalURL:       "URL_DEV_PORTAL",
	ConfigRedis:              "REDIS",
	ConfigPrivateKey:         "PRIVATE_KEY",
	ConfigPrivateKeyPassword: "PRIVATE_KEY_PASSWORD",
}

//
// Configer interface provides a base interface for application configuration parameters.
// It allows parameter registration and its value retrieval.
//
type Configer interface {
	//
	// RegisterConfigParameter registers a predefined configuration parameter for the application.
	//
	RegisterConfigParameter(configurationParameter Parameter)

	//
	// GetParameterValue returns a string parameter value.
	//
	GetParameterValue(parameter Parameter) string
}

//
// ConfigParametersDispatcher interface provides all syntax sugar for a configuration parameters registering and
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
	// Returns a private key info.
	//
	GetPrivateKeyInfo() Base64StringInfoProvider

	//
	// Returns a private key password info.
	//
	GetPrivateKeyPasswordInfo() Base64StringInfoProvider
}

//
// Config object provides an access for all the application configuration parameters.
//
type Config struct {
	parameters map[Parameter]ParameterInfoProvider
}

//
// NewConfig returns an instance of Config object
//
func NewConfig() *Config {
	config := Config{}
	config.parameters = make(map[Parameter]ParameterInfoProvider)

	return &config
}

//
// RegisterConfigParameter registers a configuration parameter for the application.
//
func (c *Config) RegisterConfigParameter(configurationParameter Parameter) {
	c.parameters[configurationParameter] = getConfigParameterEntry(configurationParameter)
}

//
// Parse parses all application configuration parameters.
//
func (c *Config) Parse() error {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	for paramType, paramEntry := range c.parameters {
		func(parameterValue *string) {
			envVariableName := EnvParameters[paramType]
			flags.StringVar(parameterValue, envVariableName, "", "")
		}(paramEntry.GetValueLink())
	}
	flags.Parse(os.Args[1:])
	c.validateParameters()

	return nil
}

//
// GetParameterValue returns a parameter value by its name.
//
func (c *Config) GetParameterValue(parameter Parameter) string {
	parameterEntry, ok := c.parameters[parameter]
	if !ok {

		return ""
	}

	return parameterEntry.GetValue()
}

//
// validateParameters validates all registered parameters.
//
func (c *Config) validateParameters() error {
	for _, parameterEntry := range c.parameters {
		if err := parameterEntry.validate(); nil != err {
			return err
		}
	}

	return nil
}

//
// getConfigParameterEntry is a config parameter entry factory.
//
func getConfigParameterEntry(parameter Parameter) ParameterInfoProvider {
	parameterName := EnvParameters[parameter]
	switch parameter {
	case ConfigCassandra:
		return &CassandraConnectionInfo{StringParameter: StringParameter{name: parameterName}}
	case ConfigRedis:
		return &RedisConnectionInfo{StringParameter: StringParameter{name: parameterName}}
	case ConfigDevPortalURL:
		return &URLInfo{StringParameter: StringParameter{name: parameterName}}
	case ConfigPrivateKey:
	case ConfigPrivateKeyPassword:
		return &Base64StringInfo{StringParameter: StringParameter{name: parameterName}}
	default:
		return &StringParameter{name: parameterName}
	}

	return nil
}
