package cfg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ameteiko/golang-kit/test/helper"
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
	helper.AssertError(t, ErrRedisConnectionStringIsEmpty, err)
}

func TestRedisParseConnectionString_WithAnIncorrectConnectionString_ReturnsAnError(t *testing.T) {
	incorrectConnectionString := "*:?//"
	ri := new(RedisConnectionInfo)
	ri.value = incorrectConnectionString

	err := ri.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrRedisConnectionStringIsIncorrect, err)
}

func TestRedisParseConnectionString_WithAnEmptyURLScheme_ReturnsAnError(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = "host.com"

	err := ri.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrRedisProtocolIsIncorrect, err)
}

func TestRedisParseConnectionString_WithAnIncorrectURLScheme_ReturnsAnError(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = "mongodb://host.com"

	err := ri.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrRedisProtocolIsIncorrect, err)
}

func TestRedisParseConnectionString_WithAnEmptyHost_ReturnsAnError(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = "redis://:3443"

	err := ri.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrRedisHostIsEmpty, err)
}

func TestRedisParseConnectionString_WithAnEmptyPort_ReturnsAnError(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = "redis://host.com"

	err := ri.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrRedisPortIsEmpty, err)
}

func TestRedisParseConnectionString_WithACorrectConnectionString_Passes(t *testing.T) {
	ri := new(RedisConnectionInfo)
	ri.value = fmt.Sprintf("redis://%s:%s", redisHost, redisPort)

	err := ri.validate()

	assert.Empty(t, err)
	assert.Equal(t, redisHost, ri.GetHost())
	assert.Equal(t, redisPort, ri.GetPort())
}
