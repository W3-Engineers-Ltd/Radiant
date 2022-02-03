// Copyright 2020
//

package ssdb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvider_SessionInit(t *testing.T) {
	// using old style
	savePath := `localhost:8080`
	cp := &Provider{}
	cp.SessionInit(context.Background(), 12, savePath)
	assert.Equal(t, "localhost", cp.Host)
	assert.Equal(t, 8080, cp.Port)
	assert.Equal(t, int64(12), cp.maxLifetime)

	savePath = `
{ "host": "localhost", "port": 8080}
`
	cp = &Provider{}
	cp.SessionInit(context.Background(), 12, savePath)
	assert.Equal(t, "localhost", cp.Host)
	assert.Equal(t, 8080, cp.Port)
	assert.Equal(t, int64(12), cp.maxLifetime)
}
