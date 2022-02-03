// Copyright 2020 radiant
//

package prometheus

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

func TestFilterChain(t *testing.T) {
	filter := (&FilterChainBuilder{}).FilterChain(func(ctx *context.Context) {
		// do nothing
		ctx.Input.SetData("invocation", true)
	})

	ctx := context.NewContext()
	r, _ := http.NewRequest("GET", "/prometheus/user", nil)
	w := httptest.NewRecorder()
	ctx.Reset(w, r)
	ctx.Input.SetData("RouterPattern", "my-route")
	filter(ctx)
	assert.True(t, ctx.Input.GetData("invocation").(bool))
	time.Sleep(1 * time.Second)
}

func TestFilterChainBuilder_report(t *testing.T) {
	ctx := context.NewContext()
	r, _ := http.NewRequest("GET", "/prometheus/user", nil)
	w := httptest.NewRecorder()
	ctx.Reset(w, r)
	fb := &FilterChainBuilder{}
	// without router info
	report(time.Second, ctx, fb.buildVec())

	ctx.Input.SetData("RouterPattern", "my-route")
	report(time.Second, ctx, fb.buildVec())
}
