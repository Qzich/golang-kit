package cache

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"

	"github.com/ameteiko/golang-kit/test/mocks"
)

//
// Common test variables.
//
var (
	Key            = "key"
	StringValue    = "string value"
	ByteArrayValue = []byte("string value")
)

//
// JSONSerialized type.
//
type JSONSerialized struct {
	ID        string `json:"id,omitempty"`
	Value     string `json:"value,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	Hash      []byte
}

//
// SetKey :: with a string value :: sets the value
//
func TestSetKeyWithStringValue(t *testing.T) {
	rc := getRedisClient()

	rc.SetKey(Key, StringValue, 0)
	value, err := rc.GetKey(Key)

	assert.Empty(t, err)
	assert.Equal(t, StringValue, string(value))
}

//
// SetKey :: with a byte array value :: sets the value
//
func TestSetKeyWithByteArrayValue(t *testing.T) {
	rc := getRedisClient()

	rc.SetKey(Key, ByteArrayValue, 0)
	value, err := rc.GetKey(Key)

	assert.Empty(t, err)
	assert.Equal(t, ByteArrayValue, value)
}

//
// SetKey :: with a byte array value :: sets the value
//
func TestSetKeyWithAnObjectValue(t *testing.T) {
	rc := getRedisClient()
	obj := JSONSerialized{
		"1234",
		"Value",
		123456,
		[]byte("Hash"),
	}
	serializedObj, _ := json.Marshal(obj)

	rc.SetKey(Key, obj, 0)
	value, err := rc.GetKey(Key)

	assert.Empty(t, err)
	assert.Equal(t, serializedObj, value)
}

//
// SetKey :: with an empty object :: passes
//
func TestSetKeyWithAnEmptyObject(t *testing.T) {
	rc := getRedisClient()

	errSet := rc.SetKey(Key, struct{}{}, 0)
	value, err := rc.GetKey(Key)

	assert.Empty(t, errSet)
	assert.Empty(t, err)
	assert.Equal(t, value, []byte("{}"))
}

//
// SetKey :: with an object that is serialized to empty struct :: passes
//
func TestSetKeyWithAnObjectWhichIsNotSerializable(t *testing.T) {
	rc := getRedisClient()

	errSet := rc.SetKey(Key, struct{ name string }{"name"}, 0)
	value, err := rc.GetKey(Key)

	assert.Empty(t, errSet)
	assert.Empty(t, err)
	assert.Equal(t, value, []byte("{}"))
}

//
// SetKey :: with an expiration timeout :: expires the value after timeout
//
func TestSetKeyWithExpiration(t *testing.T) {
	rc := getRedisClient()
	expiration := time.Millisecond * 10

	rc.SetKey(Key, StringValue, expiration)
	time.Sleep(expiration + 1)
	value, err := rc.GetKey(Key)

	assert.NotEmpty(t, err)
	assert.Empty(t, value)
}

//
// getRedisClient returns a redis client instance.
//
func getRedisClient() *Redis {
	c := redis.NewClient(&redis.Options{
		Addr: ":32768",
	})

	return NewRedisClient(c, mocks.LoggerMock{})
}
