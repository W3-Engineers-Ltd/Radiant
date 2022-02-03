// Copyright 2020 radiant
//

package opentracing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

func TestFilterChainBuilder_FilterChain(t *testing.T) {
	builder := &FilterChainBuilder{
		CustomSpanFunc: func(span opentracing.Span, ctx *context.Context) {
			span.SetTag("aa", "bbb")
		},
	}

	ctx := context.NewContext()
	r, _ := http.NewRequest("GET", "/prometheus/user", nil)
	w := httptest.NewRecorder()
	ctx.Reset(w, r)
	ctx.Input.SetData("RouterPattern", "my-route")

	filterFunc := builder.FilterChain(func(ctx *context.Context) {
		ctx.Input.SetData("opentracing", true)
	})

	filterFunc(ctx)
	assert.True(t, ctx.Input.GetData("opentracing").(bool))
}
