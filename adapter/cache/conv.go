package cache

import (
	"github.com/W3-Engineers-Ltd/Radiant/client/cache"
)

// GetString convert interface to string.
func GetString(v interface{}) string {
	return cache.GetString(v)
}

// GetInt convert interface to int.
func GetInt(v interface{}) int {
	return cache.GetInt(v)
}

// GetInt64 convert interface to int64.
func GetInt64(v interface{}) int64 {
	return cache.GetInt64(v)
}

// GetFloat64 convert interface to float64.
func GetFloat64(v interface{}) float64 {
	return cache.GetFloat64(v)
}

// GetBool convert interface to bool.
func GetBool(v interface{}) bool {
	return cache.GetBool(v)
}
