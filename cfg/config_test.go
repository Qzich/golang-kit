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
	HTTPAddressParameter = string(HTTPAddress)
	CassandraParameter   = string(Cassandra)
)

func TestNewConfig_WithoutRegisteredParameters_Passes(t *testing.T) {
	config := NewConfig()

	// No parameter registration goes here.

	assert.Equal(t, 0, len(config.parameters))
}

func TestRegisterConfigParameter_WithAValidParameter_RegistersAParameter(t *testing.T) {
	config := NewConfig()

	config.Register(HTTPAddress)

	assert.Equal(t, 1, len(config.parameters))
}

func TestRegisterConfigParameter_WithTwoValidParameters_RegistersBothParameters(t *testing.T) {
	config := NewConfig()

	config.Register(HTTPAddress)
	config.Register(Cassandra)

	assert.Equal(t, 2, len(config.parameters))
}

func TestRegisterConfigParameter_WithSequentialParameterRegistration_RegistersJustOneInstance(t *testing.T) {
	config := NewConfig()

	config.Register(HTTPAddress)
	config.Register(HTTPAddress)

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

	config.Register(HTTPAddress)
	err := config.Parse()

	assert.Error(t, err)
	helper.AssertError(t, ErrConfigParameterIsEmpty, err)
}

func TestGetParameter_WithoutAParameterRegistration_PassesAndReturnsAnEmptyValue(t *testing.T) {
	config := NewConfig()

	parsingErr := config.Parse()
	httpConfigParameter := config.GetValue(HTTPAddress)

	assert.Empty(t, parsingErr)
	assert.Empty(t, httpConfigParameter)
}

func TestGetParameter_WithAParameterRegistrationButWithoutParameterValue_ReturnsAnError(t *testing.T) {
	config := NewConfig()
	setEnvVariable(HTTPAddressParameter, "")

	config.Register(HTTPAddress)
	parsingErr := config.Parse()
	httpConfigParameter := config.GetValue(HTTPAddress)

	assert.Error(t, parsingErr)
	helper.AssertError(t, ErrConfigParameterIsEmpty, parsingErr)
	assert.Empty(t, httpConfigParameter)
}

func TestGetParameter_WithAParameterRegistrationAndValueSet_PassesAndReturnsValue(t *testing.T) {
	httpPortValue := ":8080"
	setEnvVariable(HTTPAddressParameter, httpPortValue)
	config := NewConfig()

	config.Register(HTTPAddress)
	parsingErr := config.Parse()
	httpConfigParameter := config.GetValue(HTTPAddress)

	assert.Empty(t, parsingErr)
	assert.Equal(t, httpPortValue, httpConfigParameter)
}

func TestGetParameter_WithSeveralParametersRegistrationAndValueSet_PassesAndReturnsValue(t *testing.T) {
	// HTTP port
	httpPortValue := ":8080"
	setEnvVariable(HTTPAddressParameter, httpPortValue)
	// Cassandra connection string
	cassandraConnectionString := "cassandra://127.0.0.1/virgil_card"
	setEnvVariable(CassandraParameter, cassandraConnectionString)
	// Config object
	config := NewConfig()

	config.Register(HTTPAddress)
	config.Register(Cassandra)
	parsingErr := config.Parse()
	httpConfigParameter := config.GetValue(HTTPAddress)
	cassandraConfigParameter := config.GetValue(Cassandra)

	assert.Empty(t, parsingErr)
	assert.Equal(t, httpPortValue, httpConfigParameter)
	assert.Equal(t, cassandraConnectionString, cassandraConfigParameter)
}

func TestGetParameter_WithForANotRegisteredParameter_ReturnsAnError(t *testing.T) {
	// HTTP port
	httpPortValue := ":8080"
	setEnvVariable(HTTPAddressParameter, httpPortValue)
	// Config object
	config := NewConfig()

	config.Register(HTTPAddress)
	parsingErr := config.Parse()
	cassandraConfigParameter := config.GetValue(Cassandra)

	assert.Empty(t, parsingErr)
	assert.Empty(t, cassandraConfigParameter)
}

//
// setEnvVariable sets the testing environment variable.
//
func setEnvVariable(name, value string) {
	gotenv.OverApply(strings.NewReader(fmt.Sprintf(`%s="%s"`, name, value)))
}
