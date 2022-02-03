// Copyright 2021 radiant-dev
//

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKVs(t *testing.T) {
	key := "my-key"
	kvs := NewKVs(&SimpleKV{
		Key:   key,
		Value: 12,
	})

	assert.True(t, kvs.Contains(key))

	v := kvs.GetValueOr(key, 13)
	assert.Equal(t, 12, v)

	v = kvs.GetValueOr(`key-not-exists`, 8546)
	assert.Equal(t, 8546, v)
}
