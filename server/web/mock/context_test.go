// Copyright 2021 radiant
//

package mock

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/server/web"
)

type TestController struct {
	web.Controller
}

func TestMockContext(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/hello?name=tom", bytes.NewReader([]byte{}))
	assert.Nil(t, err)
	ctx, resp := NewMockContext(req)
	ctrl := &TestController{
		Controller: web.Controller{
			Ctx: ctx,
		},
	}
	ctrl.HelloWorld()
	result := resp.BodyToString()
	assert.Equal(t, "name=tom", result)
}

// GET hello?name=XXX
func (c *TestController) HelloWorld() {
	name := c.GetString("name")
	c.Ctx.WriteString(fmt.Sprintf("name=%s", name))
}

func (c *TestController) HelloSession() {
	err := c.SessionRegenerateID()
	if err != nil {
		c.Ctx.WriteString("error")
		return
	}
	_ = c.SetSession("name", "Tom")
	c.Ctx.WriteString("set")
}

func (c *TestController) HelloSessionName() {
	name := c.CruSession.Get(context.Background(), "name")
	c.Ctx.WriteString(name.(string))
}
