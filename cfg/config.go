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
	DevPortalURL Parameter = "DEVPORTAL_URL"

	Cards4PublicKey = "CARDS4_PUBLIC_KEY"
	Cards4ReadURL   = "CARDS4_READ_URL"
	Cards4CardID    = "CARDS4_CARD_ID"

	Cards5URL = "CARDS5_URL"
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

	//
	// RegisterCassandraParameter registers Cassandra configuration parameter.
	//
	RegisterCassandraParameter(Parameter)

	//
	// GetCassandraParameter returns Cassandra configuration parameter value.
	//
	GetCassandraParameter(Parameter) (CassandraConnectionInfoProvider, error)

	//
	// RegisterRedisParameter registers Redis configuration parameter.
	//
	RegisterRedisParameter(Parameter)

	//
	// GetRedisParameter returns Redis configuration parameter value.
	//
	GetRedisParameter(Parameter) (RedisConnectionInfoProvider, error)

	//
	// RegisterLogParameter registers log configuration parameter.
	//
	RegisterLogParameter(Parameter)

	//
	// GetLogParameter returns log configuration parameter value.
	//
	GetLogParameter(Parameter) (LogInfoProvider, error)
}

//
// CommonsConfiger returns all the common application parameters.
//
type CommonsConfiger interface {
	//
	// GetCards5URL returns Cards v5 url.
	//
	GetCards5URL() (URLInfoProvider, error)

	//
	// GetCards4ReadURL returns Cards v4 read url.
	//
	GetCards4ReadURL() (URLInfoProvider, error)

	//
	// GetCards4CardID returns Cards v4 card id.
	//
	GetCards4CardID() (ParameterInfoProvider, error)

	//
	// GetCards4PublicKey returns Cards v4 public key.
	//
	GetCards4PublicKey() (Base64StringInfoProvider, error)
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
// RegisterCassandraParameter registers a cassandra connection configuration parameter.
//
func (c *Config) RegisterCassandraParameter(param Parameter) {
	c.parameters[param] = newCassandraParameter(param)
}

//
// RegisterRedisParameter registers redis connection configuration parameter.
//
func (c *Config) RegisterRedisParameter(param Parameter) {
	c.parameters[param] = newRedisParameter(param)
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
// RegisterLogParameter registers log configuration parameter.
//
func (c *Config) RegisterLogParameter(param Parameter) {
	c.parameters[param] = newLogParameter(param)
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
// GetCards5URL returns Virgil Cards service URL.
//
func (c *Config) GetCards5URL() (URLInfoProvider, error) {
	url, ok := c.parameters[Cards5URL].(URLInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetCards5URL",
		)
	}

	return url, nil
}

//
// GetCards4ReadURL returns Virgil Cards service read URL.
//
func (c *Config) GetCards4ReadURL() (URLInfoProvider, error) {
	url, ok := c.parameters[Cards4ReadURL].(URLInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetCards4ReadURL",
		)
	}

	return url, nil
}

//
// GetCards4CardID returns Virgil Cards service Virgil Card ID.
//
func (c *Config) GetCards4CardID() (ParameterInfoProvider, error) {
	id, ok := c.parameters[Cards4CardID].(ParameterInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetCards4CardID",
		)
	}

	return id, nil
}

//
// GetCards4PublicKey returns Virgil Cards service public key.
//
func (c *Config) GetCards4PublicKey() (Base64StringInfoProvider, error) {
	key, ok := c.parameters[Cards4PublicKey].(Base64StringInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetCards4PublicKey",
		)
	}

	return key, nil
}

//
// GetLogParameter returns log configuration parameter value.
//
func (c *Config) GetLogParameter(param Parameter) (LogInfoProvider, error) {
	parameter, ok := c.parameters[param].(LogInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetLogParameter",
		)
	}

	return parameter, nil
}

//
// GetCassandraParameter returns cassandra config parameter.
//
func (c *Config) GetCassandraParameter(param Parameter) (CassandraConnectionInfoProvider, error) {
	parameter, ok := c.parameters[param].(CassandraConnectionInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetCassandraParameter",
		)
	}

	return parameter, nil
}

//
// GetRedisParameter returns cassandra config parameter.
//
func (c *Config) GetRedisParameter(param Parameter) (RedisConnectionInfoProvider, error) {
	parameter, ok := c.parameters[param].(RedisConnectionInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetRedisParameter",
		)
	}

	return parameter, nil
}

//
// GetStringParameter returns a string parameter info.
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
	url, ok := c.parameters[param].(URLInfoProvider)
	if !ok {
		return nil, errors.WithMessage(
			errors.ErrGetMisregisteredConfigParameter,
			"kit-cfg@Config.GetURLParameter",
		)
	}

	return url, nil
}

//
// validateParameters validates all registered parameters.
//
func (c *Config) validateParameters() error {
	for _, parameterEntry := range c.parameters {
		if err := parameterEntry.validate(); nil != err {
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
	case DevPortalURL:
		return newURLParameter(parameter)
	case Cards5URL:
		return newURLParameter(parameter)
	case Cards4ReadURL:
		return newURLParameter(parameter)
	case Cards4PublicKey:
		return newBase64Parameter(parameter)
	case Cards4CardID:
		return newStringParameter(parameter)
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

//
// newCassandraParameter registers a new Cassandra parameter.
//
func newCassandraParameter(param Parameter) *CassandraConnectionInfo {

	return &CassandraConnectionInfo{StringParameter: newStringParameter(param)}
}

//
// newRedisParameter registers a new Redis parameter.
//
func newRedisParameter(param Parameter) *RedisConnectionInfo {

	return &RedisConnectionInfo{StringParameter: newStringParameter(param)}
}

//
// newCassandraParameter registers a new Cassandra parameter.
//
func newLogParameter(param Parameter) *LogInfo {

	return &LogInfo{StringParameter: newStringParameter(param)}
}
