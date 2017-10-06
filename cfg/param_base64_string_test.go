package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ameteiko/golang-kit/test/helper"
	"encoding/base64"
)

func TestBase64StringValidate_WithAnEmptyString_ReturnsAnError(t *testing.T) {
	si := Base64StringInfo{}

	err := si.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrConfigParameterIsEmpty, err)
}

func TestBase64StringValidate_WithNotABase64EncodedString_ReturnsAnError(t *testing.T) {
	notABase64EncodedValue := "!@#$%^&*()_"
	si := Base64StringInfo{}
	si.value = notABase64EncodedValue

	err := si.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrBase64IncorrectValue, err)
}

func TestBase64StringValidate_WithABase64EncodedString_Passes(t *testing.T) {
	rawValue := []byte("Some string")
	si := Base64StringInfo{}
	si.value = base64.StdEncoding.EncodeToString(rawValue)

	err := si.validate()

	assert.Empty(t, err)
	assert.Equal(t, rawValue, si.GetDecodedValue())
}
