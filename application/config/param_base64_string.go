package config

import (
	"encoding/base64"
	"github.com/ameteiko/errors"
)

//
// RedisConnectionInfoProvider declares all the connection info getters.
//
type Base64StringInfoProvider interface {
	GetDecodedValue() []byte

	ParameterInfoProvider
}

//
// Base64StringInfo is a base64-encoded string configuration parameter.
//
type Base64StringInfo struct {
	decodedValue []byte
	StringParameter
}

//
// GetDecodedValue returns a decoded parameter value.
//
func (p *Base64StringInfo) GetDecodedValue() []byte {

	return p.decodedValue
}

//
// validate validates the string configuration parameter value not to be an empty string.
//
func (p *Base64StringInfo) validate() error {
	var err error

	if err := p.StringParameter.validate(); nil != err {

		return err
	}

	if p.decodedValue, err = base64.StdEncoding.DecodeString(p.GetValue()); nil != err {

		return errors.WrapError(
			err,
			errors.Errorf(`unable to base64-decode parameter (%s)`, p.GetValue()),
			ErrBase64IncorrectValue,
		)
	}

	return nil
}
