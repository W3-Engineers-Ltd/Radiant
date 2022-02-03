// Copyright 2020
//

package couchbase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvider_SessionInit(t *testing.T) {
	// using old style
	savePath := `http://host:port/,Pool,Bucket`
	cp := &Provider{}
	cp.SessionInit(context.Background(), 12, savePath)
	assert.Equal(t, "http://host:port/", cp.SavePath)
	assert.Equal(t, "Pool", cp.Pool)
	assert.Equal(t, "Bucket", cp.Bucket)
	assert.Equal(t, int64(12), cp.maxlifetime)

	savePath = `
{ "save_path": "my save path", "pool": "mypool", "bucket": "mybucket"}
`
	cp = &Provider{}
	cp.SessionInit(context.Background(), 12, savePath)
	assert.Equal(t, "my save path", cp.SavePath)
	assert.Equal(t, "mypool", cp.Pool)
	assert.Equal(t, "mybucket", cp.Bucket)
	assert.Equal(t, int64(12), cp.maxlifetime)
}
