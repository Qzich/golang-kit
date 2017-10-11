package cfg

import (
	"encoding/base64"

	"github.com/ameteiko/golang-kit/errors"
)

//
// Base64StringInfoProvider declares base64-encoded string getters.
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

	*StringParameter
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

	if err = p.StringParameter.validate(); nil != err {
		return err
	}

	if p.decodedValue, err = base64.StdEncoding.DecodeString(p.GetValue()); nil != err {
		return errors.WrapError(
			errors.WithMessage(
				err,
				`kit-cfg@Base64StringInfo.validate [parameter (%s), value ($s)]`,
				p.GetValue(),
				p.GetName(),
			),
			ErrBase64IncorrectValue,
		)
	}

	return nil
}
