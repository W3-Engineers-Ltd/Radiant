// Copyright 2021 radiant
//

package httplib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHttpResponseWithJsonBody(t *testing.T) {
	// string
	resp := NewHttpResponseWithJsonBody("{}")
	assert.Equal(t, int64(2), resp.ContentLength)

	resp = NewHttpResponseWithJsonBody([]byte("{}"))
	assert.Equal(t, int64(2), resp.ContentLength)

	resp = NewHttpResponseWithJsonBody(&user{
		Name: "Tom",
	})
	assert.True(t, resp.ContentLength > 0)
}
