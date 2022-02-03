// Copyright 2021 radiant
//

package mock

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/server/web"
)

func TestSessionProvider(t *testing.T) {
	sp := NewSessionProvider("file")
	assert.NotNil(t, sp)

	req, err := http.NewRequest("GET", "http://localhost:8080/hello?name=tom", bytes.NewReader([]byte{}))
	assert.Nil(t, err)
	ctx, resp := NewMockContext(req)
	ctrl := &TestController{
		Controller: web.Controller{
			Ctx: ctx,
		},
	}
	ctrl.HelloSession()
	result := resp.BodyToString()
	assert.Equal(t, "set", result)

	resp.Reset()
	ctrl.HelloSessionName()
	result = resp.BodyToString()

	assert.Equal(t, "Tom", result)
}
