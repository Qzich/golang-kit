package config

import (
	"net/url"

	"github.com/ameteiko/errors"
)

//
// URLInfo is a URL config parameter.
//
type URLInfo struct {
	StringParameter
}

//
// validate validates the URL config parameter.
//
func (c *URLInfo) validate() error {
	if err := c.StringParameter.validate(); nil != err {

		return err
	}
	var err error
	urlValue := c.GetValue()

	if _, err = url.Parse(urlValue); nil != err {

		return errors.WrapError(
			err,
			errors.Errorf(`incorrect URL value (%s)`, urlValue),
			ErrURLIncorrectValue,
		)
	}

	return nil
}
