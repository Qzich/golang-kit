package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subosito/gotenv"
)

//
// Common variables.
//
var (
	HTTPPortParameter  = EnvParameters[ConfigHTTPPort]
	CassandraParameter = EnvParameters[ConfigCassandra]
)

func TestNewConfig_WithoutRegisteredParameters_Passes(t *testing.T) {
	config := NewConfig()

	// No parameter registration goes here.

	assert.Equal(t, 0, len(config.parameters))
}

func TestRegisterConfigParameter_WithAValidParameter_RegistersAParameter(t *testing.T) {
	config := NewConfig()

	config.RegisterConfigParameter(ConfigHTTPPort)

	assert.Equal(t, 1, len(config.parameters))
}

func TestRegisterConfigParameter_WithTwoValidParameters_RegistersBothParameters(t *testing.T) {
	config := NewConfig()

	config.RegisterConfigParameter(ConfigHTTPPort)
	config.RegisterConfigParameter(ConfigCassandra)

	assert.Equal(t, 2, len(config.parameters))
}

func TestRegisterConfigParameter_WithSequentialParameterRegistration_RegistersJustOneInstance(t *testing.T) {
	config := NewConfig()

	config.RegisterConfigParameter(ConfigHTTPPort)
	config.RegisterConfigParameter(ConfigHTTPPort)

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

func TestParse_WithAParameterRegistration_Passes(t *testing.T) {
	config := NewConfig()

	config.RegisterConfigParameter(ConfigHTTPPort)
	err := config.Parse()

	assert.Empty(t, err)
}

func TestGetParameter_WithoutAParameterRegistration_PassesAndReturnsAnEmptyValue(t *testing.T) {
	config := NewConfig()

	parsingErr := config.Parse()
	httpConfigParameter := config.GetParameterValue(ConfigHTTPPort)

	assert.Empty(t, parsingErr)
	assert.Empty(t, httpConfigParameter)
}

func TestGetParameter_WithAParameterRegistrationButWithoutParameterValue_PassesAndReturnsAnEmptyValue(t *testing.T) {
	config := NewConfig()
	setEnvVariable(HTTPPortParameter, "")

	config.RegisterConfigParameter(ConfigHTTPPort)
	parsingErr := config.Parse()
	httpConfigParameter := config.GetParameterValue(ConfigHTTPPort)

	assert.Empty(t, parsingErr)
	assert.Empty(t, httpConfigParameter)
}

func TestGetParameter_WithAParameterRegistrationAndValueSet_PassesAndReturnsValue(t *testing.T) {
	httpPortValue := ":8080"
	setEnvVariable(HTTPPortParameter, httpPortValue)
	config := NewConfig()

	config.RegisterConfigParameter(ConfigHTTPPort)
	parsingErr := config.Parse()
	httpConfigParameter := config.GetParameterValue(ConfigHTTPPort)

	assert.Empty(t, parsingErr)
	assert.Equal(t, httpPortValue, httpConfigParameter)
}

func TestGetParameter_WithSeveralParametersRegistrationAndValueSet_PassesAndReturnsValue(t *testing.T) {
	// HTTP port
	httpPortValue := ":8080"
	setEnvVariable(HTTPPortParameter, httpPortValue)
	// Cassandra connection string
	cassandraConnectionString := "cassandra://127.0.0.1/virgil_card"
	setEnvVariable(CassandraParameter, cassandraConnectionString)
	// Config object
	config := NewConfig()

	config.RegisterConfigParameter(ConfigHTTPPort)
	config.RegisterConfigParameter(ConfigCassandra)
	parsingErr := config.Parse()
	httpConfigParameter := config.GetParameterValue(ConfigHTTPPort)
	cassandraConfigParameter := config.GetParameterValue(ConfigCassandra)

	assert.Empty(t, parsingErr)
	assert.Equal(t, httpPortValue, httpConfigParameter)
	assert.Equal(t, cassandraConnectionString, cassandraConfigParameter)
}

func TestGetParameter_WithForANotRegisteredParameter_ReturnsAnError(t *testing.T) {
	// HTTP port
	httpPortValue := ":8080"
	setEnvVariable(HTTPPortParameter, httpPortValue)
	// Config object
	config := NewConfig()

	config.RegisterConfigParameter(ConfigHTTPPort)
	parsingErr := config.Parse()
	cassandraConfigParameter := config.GetParameterValue(ConfigCassandra)

	assert.Empty(t, parsingErr)
	assert.Empty(t, cassandraConfigParameter)
}

//
// setEnvVariable sets the testing environment variable.
//
func setEnvVariable(name, value string) {
	gotenv.OverApply(strings.NewReader(fmt.Sprintf(`%s="%s"`, name, value)))
}
