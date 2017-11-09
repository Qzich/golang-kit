package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"

	"github.com/ameteiko/golang-kit/errors"
	"github.com/ameteiko/golang-kit/log"
)

//
// Redis client instance.
//
type Redis struct {
	client *redis.Client // TODO:: I want to get rid of this hardcoded dependency here!!!
	logger log.Logger
}

//
// NewRedisClient returns a new Redis client instance.
//
func NewRedisClient(c *redis.Client, l log.Logger) *Redis {

	return &Redis{c, l}
}

//
// SetKey sets a value for key for duration.
//
func (r *Redis) SetKey(key string, val interface{}, expiration time.Duration) error {
	var value string
	switch v := val.(type) {
	case string:
		value = v
	case []byte:
		value = string(v)
	default:
		vv, err := json.Marshal(v)
		if nil != err {
			r.logger.Debug("redis value marshalling error (%s)", val)
			return errors.WithMessage(err, "kit.cache@Redis.SetKey [value (%s) marshalling error for key (%s)]", val, key)
		}
		value = string(vv)
	}

	if err := r.client.Set(key, value, expiration).Err(); nil != err {
		r.logger.Debug("redis client set value error (%s) for key (%s)", val, key)
		return errors.WithMessage(err, "kit.cache@Redis.SetKey [key (%s) setting error for val (%s)]", key, val)
	}

	return nil
}

//
// GetKey returns key value.
//
func (r *Redis) GetKey(key string) ([]byte, error) {
	result, err := r.client.Get(key).Result()
	if nil != err {
		return nil, errors.WithMessage(err, "kit.cache@Redis.SetKey [error reading key (%s) value]", key)
	}

	return []byte(result), nil
}
