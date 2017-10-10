package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// Testing parameter names.
//
const (
	TestParameter Parameter = "TEST_PARAMETER"
)

func TestRegisterHttpParser_WithParameters_PassesAndRegistersAParser(t *testing.T) {
	// HTTP port
	httpPort := "8080"
	httpPortEnvVariableName := string(TestParameter)
	setEnvVariable(httpPortEnvVariableName, httpPort)
	// Config object

	config := NewConfig()
	config.Register(TestParameter)
	errParsing := config.Parse()
	port := config.GetValue(TestParameter)

	assert.Empty(t, errParsing)
	assert.Equal(t, httpPort, port)
}
