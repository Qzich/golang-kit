package config

import (
	"github.com/ameteiko/errors"
)

//
// ParameterInfoProvider is an interface for the String parameter entry.
//
type ParameterInfoProvider interface {
	//
	// GetValue returns a string parameter value.
	//
	GetValue() string

	//
	// GetValueLink returns the value link for passing to the flags parsing block.
	//
	GetValueLink() *string

	//
	// GetName returns a configuration parameter name.
	//
	GetName() string

	//
	// validate validates the string configuration parameter value not to be an empty string.
	//
	validate() error
}

//
// StringParameter is a string configuration parameter.
// String parameter is a base parameter type for each configuration parameter like database connection string, web
// service URL, etc.
//
type StringParameter struct {
	name  string
	value string
}

//
// GetValue returns a string parameter value.
//
func (p *StringParameter) GetValue() string {
	return p.value
}

//
// GetValueLink returns value link.
//
func (p *StringParameter) GetValueLink() *string {
	return &p.value
}

//
// GetName returns a configuration parameter name.
//
func (p *StringParameter) GetName() string {
	return p.name
}

//
// validate validates the string configuration parameter value not to be an empty string.
//
func (p *StringParameter) validate() error {
	if "" == p.GetValue() {

		return errors.Wrapf(ErrConfigParameterIsEmpty, `configuration parameter (%s) is empty`, p.GetName())
	}

	return nil
}

//
// RegisterHTTPPortParser registers an HTTP port config parser.
//
func (c *Config) RegisterHTTPPortParser() {

	c.RegisterConfigParameter(ConfigHTTPPort)
}
