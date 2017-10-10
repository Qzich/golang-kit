package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ameteiko/golang-kit/test/helper"
)

func TestURLValidate_WithAnEmptyString_ReturnsAnError(t *testing.T) {
	ui := newURLInfo()

	err := ui.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrConfigParameterIsEmpty, err)
}

func TestURLValidate_WithAnIncorrectString_ReturnsAnError(t *testing.T) {
	incorrectConnectionString := "*:?//"
	ui := newURLInfo()
	ui.value = incorrectConnectionString

	err := ui.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrURLIncorrectValue, err)
}

func TestURLValidate_WithACorrectString_Passes(t *testing.T) {
	url := "https://google.com"
	ui := newURLInfo()
	ui.value = url

	err := ui.validate()

	assert.Empty(t, err)
	assert.Equal(t, url, ui.GetValue())
}
