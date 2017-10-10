package cfg

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ameteiko/golang-kit/test/helper"
	"github.com/stretchr/testify/assert"
	"github.com/subosito/gotenv"
)

//
// Common variables.
//
var (
	EnvCards4CardID  = string(Cards4CardID)
	EnvCards4ReadURL = string(Cards4ReadURL)
)

func TestNewConfig_WithoutRegisteredParameters_Passes(t *testing.T) {
	config := NewConfig()

	// No parameter registration goes here.

	assert.Equal(t, 0, len(config.parameters))
}

func TestRegisterConfigParameter_WithAValidParameter_RegistersAParameter(t *testing.T) {
	config := NewConfig()

	config.Register(Cards4CardID)

	assert.Equal(t, 1, len(config.parameters))
}

func TestRegisterConfigParameter_WithTwoValidParameters_RegistersBothParameters(t *testing.T) {
	config := NewConfig()

	config.Register(Cards4CardID)
	config.Register(DevPortalURL)

	assert.Equal(t, 2, len(config.parameters))
}

func TestRegisterConfigParameter_WithSequentialParameterRegistration_RegistersJustOneInstance(t *testing.T) {
	config := NewConfig()

	config.Register(DevPortalURL)
	config.Register(DevPortalURL)

	assert.Equal(t, 1, len(config.parameters))
}

func TestParse_SequentialCalls_PassesAndReturnsNoError(t *testing.T) {
	config := NewConfig()

	// No parameter registration goes here.
	err1 := config.Parse()
	err2 := config.Parse()

	assert.Empty(t, err1)
	assert.Empty(t, err2)
}

func TestParse_WithoutParameterRegistration_Passes(t *testing.T) {
	config := NewConfig()

	// No parameter registration goes here.
	err := config.Parse()

	assert.Empty(t, err)
}

func TestParse_WithAParameterRegistration_ReturnsAnError(t *testing.T) {
	config := NewConfig()

	config.Register(DevPortalURL)
	err := config.Parse()

	assert.Error(t, err)
	helper.AssertError(t, ErrConfigParameterIsEmpty, err)
}

func TestGetParameter_WithoutAParameterRegistration_PassesAndReturnsAnEmptyValue(t *testing.T) {
	config := NewConfig()

	parsingErr := config.Parse()
	httpConfigParameter := config.GetValue(DevPortalURL)

	assert.Empty(t, parsingErr)
	assert.Empty(t, httpConfigParameter)
}

func TestGetParameter_WithAParameterRegistrationButWithoutParameterValue_ReturnsAnError(t *testing.T) {
	config := NewConfig()
	setEnvVariable(EnvCards4ReadURL, "")

	config.Register(DevPortalURL)
	parsingErr := config.Parse()
	httpConfigParameter := config.GetValue(Cards4ReadURL)

	assert.Error(t, parsingErr)
	helper.AssertError(t, ErrConfigParameterIsEmpty, parsingErr)
	assert.Empty(t, httpConfigParameter)
}

func TestGetParameter_WithAParameterRegistrationAndValueSet_PassesAndReturnsValue(t *testing.T) {
	httpPortValue := "CARD ID"
	setEnvVariable(EnvCards4CardID, httpPortValue)
	config := NewConfig()

	config.Register(Cards4CardID)
	parsingErr := config.Parse()
	httpConfigParameter := config.GetValue(Cards4CardID)

	assert.Empty(t, parsingErr)
	assert.Equal(t, httpPortValue, httpConfigParameter)
}

func TestGetParameter_WithSeveralParametersRegistrationAndValueSet_PassesAndReturnsValue(t *testing.T) {
	urlValue := "https://virgil.com"
	setEnvVariable(EnvCards4ReadURL, urlValue)
	idValue := "Card ID"
	setEnvVariable(EnvCards4CardID, idValue)
	// Config object
	config := NewConfig()

	config.Register(Cards4ReadURL)
	config.Register(Cards4CardID)
	parsingErr := config.Parse()
	url := config.GetValue(Cards4ReadURL)
	id := config.GetValue(Cards4CardID)

	assert.Empty(t, parsingErr)
	assert.Equal(t, urlValue, url)
	assert.Equal(t, idValue, id)
}

func TestGetParameter_WithForANotRegisteredParameter_ReturnsAnError(t *testing.T) {
	cardID := "Card ID"
	setEnvVariable(EnvCards4CardID, cardID)
	config := NewConfig()

	config.Register(Cards4CardID)
	parsingErr := config.Parse()
	url := config.GetValue(Cards4ReadURL)

	assert.Empty(t, parsingErr)
	assert.Empty(t, url)
}

//
// setEnvVariable sets the testing environment variable.
//
func setEnvVariable(name, value string) {
	gotenv.OverApply(strings.NewReader(fmt.Sprintf(`%s="%s"`, name, value)))
}
