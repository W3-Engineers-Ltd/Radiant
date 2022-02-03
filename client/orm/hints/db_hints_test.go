// Copyright 2020 radiant-dev
//

package hints

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewHintTime(t *testing.T) {
	key := "qweqwe"
	value := time.Second
	hint := NewHint(key, value)

	assert.Equal(t, hint.GetKey(), key)
	assert.Equal(t, hint.GetValue(), value)
}

func TestNewHintInt(t *testing.T) {
	key := "qweqwe"
	value := 281230
	hint := NewHint(key, value)

	assert.Equal(t, hint.GetKey(), key)
	assert.Equal(t, hint.GetValue(), value)
}

func TestNewHintFloat(t *testing.T) {
	key := "qweqwe"
	value := 21.2459753
	hint := NewHint(key, value)

	assert.Equal(t, hint.GetKey(), key)
	assert.Equal(t, hint.GetValue(), value)
}

func TestForceIndex(t *testing.T) {
	s := []string{`f_index1`, `f_index2`, `f_index3`}
	hint := ForceIndex(s...)
	assert.Equal(t, hint.GetValue(), s)
	assert.Equal(t, hint.GetKey(), KeyForceIndex)
}

func TestForceIndex0(t *testing.T) {
	var s []string
	hint := ForceIndex(s...)
	assert.Equal(t, hint.GetValue(), s)
	assert.Equal(t, hint.GetKey(), KeyForceIndex)
}

func TestIgnoreIndex(t *testing.T) {
	s := []string{`i_index1`, `i_index2`, `i_index3`}
	hint := IgnoreIndex(s...)
	assert.Equal(t, hint.GetValue(), s)
	assert.Equal(t, hint.GetKey(), KeyIgnoreIndex)
}

func TestIgnoreIndex0(t *testing.T) {
	var s []string
	hint := IgnoreIndex(s...)
	assert.Equal(t, hint.GetValue(), s)
	assert.Equal(t, hint.GetKey(), KeyIgnoreIndex)
}

func TestUseIndex(t *testing.T) {
	s := []string{`u_index1`, `u_index2`, `u_index3`}
	hint := UseIndex(s...)
	assert.Equal(t, hint.GetValue(), s)
	assert.Equal(t, hint.GetKey(), KeyUseIndex)
}

func TestUseIndex0(t *testing.T) {
	var s []string
	hint := UseIndex(s...)
	assert.Equal(t, hint.GetValue(), s)
	assert.Equal(t, hint.GetKey(), KeyUseIndex)
}

func TestForUpdate(t *testing.T) {
	hint := ForUpdate()
	assert.Equal(t, hint.GetValue(), true)
	assert.Equal(t, hint.GetKey(), KeyForUpdate)
}

func TestDefaultRelDepth(t *testing.T) {
	hint := DefaultRelDepth()
	assert.Equal(t, hint.GetValue(), true)
	assert.Equal(t, hint.GetKey(), KeyRelDepth)
}

func TestRelDepth(t *testing.T) {
	hint := RelDepth(157965)
	assert.Equal(t, hint.GetValue(), 157965)
	assert.Equal(t, hint.GetKey(), KeyRelDepth)
}

func TestLimit(t *testing.T) {
	hint := Limit(1579625)
	assert.Equal(t, hint.GetValue(), int64(1579625))
	assert.Equal(t, hint.GetKey(), KeyLimit)
}

func TestOffset(t *testing.T) {
	hint := Offset(int64(1572123965))
	assert.Equal(t, hint.GetValue(), int64(1572123965))
	assert.Equal(t, hint.GetKey(), KeyOffset)
}

func TestOrderBy(t *testing.T) {
	hint := OrderBy(`-ID`)
	assert.Equal(t, hint.GetValue(), `-ID`)
	assert.Equal(t, hint.GetKey(), KeyOrderBy)
}
