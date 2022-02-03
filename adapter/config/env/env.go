// Copyright 2021 radiant Author. All Rights Reserved.

// Package env is used to parse environment.
package env

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/config/env"
)

// Get returns a value by key.
// If the key does not exist, the default value will be returned.
func Get(key string, defVal string) string {
	return env.Get(key, defVal)
}

// MustGet returns a value by key.
// If the key does not exist, it will return an error.
func MustGet(key string) (string, error) {
	return env.MustGet(key)
}

// Set sets a value in the ENV copy.
// This does not affect the child process environment.
func Set(key string, value string) {
	env.Set(key, value)
}

// MustSet sets a value in the ENV copy and the child process environment.
// It returns an error in case the set operation failed.
func MustSet(key string, value string) error {
	return env.MustSet(key, value)
}

// GetAll returns all keys/values in the current child process environment.
func GetAll() map[string]string {
	return env.GetAll()
}
