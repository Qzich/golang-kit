//
// Package config provides a syntax sugar for the the application configuration parameters access.
//
package config

import (
	"github.com/namsral/flag"
	"os"
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
)

//
// EnvParameters are the predefined configuration parameter types names.
//
var EnvParameters = map[Parameter]string{
	ConfigCassandra:    "CASSANDRA",
	ConfigHTTPPort:     "HTTP_PORT",
	ConfigDevPortalURL: "URL_DEV_PORTAL",
	ConfigRedis:        "REDIS",
}

//
// Configer interface provides a base interface for application configuration parameters.
// It allows parameter registration and its value retrieval.
// ITODO: Add named configuration parameters with string values.
//
type Configer interface {
	//
	// RegisterConfigParameter registers a predefined configuration parameter for the application.
	//
	RegisterConfigParameter(configurationParameter Parameter) *Config

	//
	// GetParameterValue returns a string parameter value.
	//
	GetParameter(parameter Parameter) (string, error)
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
	GetRedisConnectionInfo() CassandraConnectionInfoProvider

	//
	// Returns a Development Portal URL
	//
	//GetDevPortalURL() string
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
	default:
		return &StringParameter{name: parameterName}
	}
}
