package cfg

import (
	"net/url"
	"regexp"

	"github.com/ameteiko/golang-kit/errors"
)

//
// URLInfoProvider declares all the URL info getters.
//
type URLInfoProvider interface {
	ParameterInfoProvider

	GetHost() string
	GetNotVersionedURL() string
}

//
// URLInfo is a URL config parameter.
//
type URLInfo struct {
	host string

	*StringParameter
}

//
// newURLInfo returns a new URL info object instance.
//
func newURLInfo() *URLInfo {
	return &URLInfo{
		StringParameter: &StringParameter{},
	}
}

//
// validate validates the URL config parameter.
//
func (c *URLInfo) validate() error {
	var err error
	if err = c.StringParameter.validate(); nil != err {
		return err
	}

	urlValue := c.GetValue()
	urlInfo, err := url.Parse(urlValue)
	if nil != err {
		return errors.WrapError(
			ErrURLIncorrectValue,
			errors.WithMessage(err, `kit-cfg@URLInfo.validate [value (%s)]`, urlValue),
		)
	}
	c.host = urlInfo.Host

	return nil
}

//
// GetNotVersionedURL returns a URL value without version ending part
//
func (c *URLInfo) GetNotVersionedURL() (string, error) {
	re := regexp.MustCompile("(.*)/v[0-9]*")
	match := re.FindStringSubmatch(c.GetValue())
	if nil == match {
		return "", ErrURLDoesNotContainVersionPart
	}

	return match[1], nil
}

//
// GetHost returns host value.
//
func (c *URLInfo) GetHost() string {

	return c.host
}
