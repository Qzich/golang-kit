package config

import (
	"net/url"

	"github.com/ameteiko/golang-kit/errors"
)

//
// RegisterDevPortalURLParser registers a dev portal URL config parser.
//
func (c *Config) RegisterDevPortalURLParser() {
	c.RegisterConfigParameter(ConfigDevPortalURL)
}

//
// URLInfo is a URL config parameter.
//
type URLInfo struct {
	StringParameter

	hosts      []string
	keyspace   string
	dataCenter string
	user       string
	password   string
}

//
// URLInfoProvider declares all the URL getters.
//
type URLInfoProvider interface {
	//
	// GetURL returns the URL value.
	//
	GetURL() string
}

//
// GetURL returns the URL value.
//
func (c *URLInfo) GetURL() string {
	return c.GetValue()
}

//
// validate validates the cassandra connection string parameter.
//
func (c *URLInfo) validate() error {
	var err error
	urlParameter := c.GetValue()

	if "" == urlParameter {
		return ErrConnectionStringIsEmpty
	}

	if _, err = url.Parse(c.GetValue()); nil != err {
		return errors.Wrapf(
			err,
			`incorrect database connection string (%s)`,
			c.GetValue(),
		)
	}

	return nil
}
