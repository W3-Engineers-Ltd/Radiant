package ledis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvider_SessionInit(t *testing.T) {
	// using old style
	savePath := `http://host:port/,100`
	cp := &Provider{}
	cp.SessionInit(context.Background(), 12, savePath)
	assert.Equal(t, "http://host:port/", cp.SavePath)
	assert.Equal(t, 100, cp.Db)
	assert.Equal(t, int64(12), cp.maxlifetime)

	savePath = `
{ "save_path": "my save path", "db": 100}
`
	cp = &Provider{}
	cp.SessionInit(context.Background(), 12, savePath)
	assert.Equal(t, "my save path", cp.SavePath)
	assert.Equal(t, 100, cp.Db)
	assert.Equal(t, int64(12), cp.maxlifetime)
}
