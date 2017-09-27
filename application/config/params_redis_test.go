package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// Testing constants.
//
const (
	redisHost = "host1.com"
	redisPort = "2345"
)

func TestRedisParseConnectionString_WithAnEmptyString_ReturnsAnError(t *testing.T) {
	ri := new(RedisConnectionInfo)

	err := ri.validate()

	assert.Error(t, err)
	assert.Equal(t, ErrRedisConnectionStringIsEmpty, err)
}

func TestRedisParseConnectionString_WithAnIncorrectConnectionString_ReturnsAnError(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = "Incorrect connection string"

	err := ri.validate()

	assert.Error(t, err)
}

func TestRedisParseConnectionString_WithAnEmptyURLScheme_ReturnsAnError(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = "host.com"

	err := ri.validate()

	assert.Error(t, err)
	//helper.AssertError(t, ErrConnectionStringSchemeIsIncorrect, err)
}

func TestRedisParseConnectionString_WithAnIncorrectURLScheme_ReturnsAnError(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = "mongodb://host.com"

	err := ri.validate()

	assert.Error(t, err)
	//helper.AssertError(t, ErrConnectionStringSchemeIsIncorrect, err)
}

func TestRedisParseConnectionString_WithAnEmptyPort_ReturnsAnError(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = "redis://host.com"

	err := ri.validate()

	assert.Error(t, err)
	//helper.AssertError(t, ErrConnectionStringKeyspaceIsEmpty, err)
}

func TestRedisParseConnectionString_WithACorrectConnectionString_Passes(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = fmt.Sprintf("redis://%s:%s", redisHost, redisPort)

	err := ri.validate()

	assert.Empty(t, err)
	assert.Equal(t, redisHost, ri.GetHost())
	assert.Equal(t, redisPort, ri.GetPort())
}
