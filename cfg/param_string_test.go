package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHttpParser_WithParameters_PassesAndRegistersAParser(t *testing.T) {
	// HTTP port
	httpPort := ":8080"
	httpPortEnvVariableName := string(HTTPAddress)
	setEnvVariable(httpPortEnvVariableName, httpPort)
	// Config object
	config := NewConfig()

	config.Register(HTTPAddress)
	errParsing := config.Parse()
	port := config.GetValue(HTTPAddress)

	assert.Empty(t, errParsing)
	assert.Equal(t, httpPort, port)
}
