package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHttpParser_WithParameters_PassesAndRegistersAParser(t *testing.T) {
	// HTTP port
	httpPort := ":8080"
	httpPortEnvVariableName := EnvParameters[ConfigHTTPPort]
	setEnvVariable(httpPortEnvVariableName, httpPort)
	// Config object
	config := NewConfig()

	config.RegisterConfigParameter(ConfigHTTPPort)
	errParsing := config.Parse()
	port := config.GetParameterValue(ConfigHTTPPort)

	assert.Empty(t, errParsing)
	assert.Equal(t, httpPort, port)
}
