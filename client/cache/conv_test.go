package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	t1 := "test1"

	assert.Equal(t, "test1", GetString(t1))
	t2 := []byte("test2")
	assert.Equal(t, "test2", GetString(t2))
	t3 := 1
	assert.Equal(t, "1", GetString(t3))
	var t4 int64 = 1
	assert.Equal(t, "1", GetString(t4))
	t5 := 1.1
	assert.Equal(t, "1.1", GetString(t5))
	assert.Equal(t, "", GetString(nil))
}

func TestGetInt(t *testing.T) {
	t1 := 1
	assert.Equal(t, 1, GetInt(t1))
	var t2 int32 = 32
	assert.Equal(t, 32, GetInt(t2))

	var t3 int64 = 64
	assert.Equal(t, 64, GetInt(t3))
	t4 := "128"

	assert.Equal(t, 128, GetInt(t4))
	assert.Equal(t, 0, GetInt(nil))
}

func TestGetInt64(t *testing.T) {
	var i int64 = 1
	t1 := 1
	assert.Equal(t, i, GetInt64(t1))
	var t2 int32 = 1

	assert.Equal(t, i, GetInt64(t2))
	var t3 int64 = 1
	assert.Equal(t, i, GetInt64(t3))
	t4 := "1"
	assert.Equal(t, i, GetInt64(t4))
	assert.Equal(t, int64(0), GetInt64(nil))
}

func TestGetFloat64(t *testing.T) {
	f := 1.11
	var t1 float32 = 1.11
	assert.Equal(t, f, GetFloat64(t1))
	t2 := 1.11
	assert.Equal(t, f, GetFloat64(t2))
	t3 := "1.11"
	assert.Equal(t, f, GetFloat64(t3))

	var f2 float64 = 1
	t4 := 1
	assert.Equal(t, f2, GetFloat64(t4))

	assert.Equal(t, float64(0), GetFloat64(nil))
}

func TestGetBool(t *testing.T) {
	t1 := true
	assert.True(t, GetBool(t1))
	t2 := "true"
	assert.True(t, GetBool(t2))

	assert.False(t, GetBool(nil))
}
