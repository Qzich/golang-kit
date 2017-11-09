package cache

import (
	"time"
)

//
// Client is a cache client interface.
//
type Client interface {
	//
	// SetKey sets a value for key for duration.
	//
	SetKey(key string, value interface{}, expiration time.Duration) error

	//
	// GetKey returns key value.
	//
	GetKey(key string) ([]byte, error)
}
