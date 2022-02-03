// Copyright 2021 radiant
//

package opentracing

import (
	"context"
	"strings"

	"github.com/opentracing/opentracing-go"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

// FilterChainBuilder provides an extension point
// this Filter's behavior looks a little bit strange
// for example:
// if we want to trace QuerySetter
// actually we trace invoking "QueryTable"
// the method Begin*, Commit and Rollback are ignored.
// When use using those methods, it means that they want to manager their transaction manually, so we won't handle them.
type FilterChainBuilder struct {
	// CustomSpanFunc users are able to custom their span
	CustomSpanFunc func(span opentracing.Span, ctx context.Context, inv *orm.Invocation)
}

func (builder *FilterChainBuilder) FilterChain(next orm.Filter) orm.Filter {
	return func(ctx context.Context, inv *orm.Invocation) []interface{} {
		operationName := builder.operationName(ctx, inv)
		if strings.HasPrefix(inv.Method, "Begin") || inv.Method == "Commit" || inv.Method == "Rollback" {
			return next(ctx, inv)
		}

		span, spanCtx := opentracing.StartSpanFromContext(ctx, operationName)
		defer span.Finish()
		res := next(spanCtx, inv)
		builder.buildSpan(span, spanCtx, inv)
		return res
	}
}

func (builder *FilterChainBuilder) buildSpan(span opentracing.Span, ctx context.Context, inv *orm.Invocation) {
	span.SetTag("orm.method", inv.Method)
	span.SetTag("orm.table", inv.GetTableName())
	span.SetTag("orm.insideTx", inv.InsideTx)
	span.SetTag("orm.txName", ctx.Value(orm.TxNameKey))
	span.SetTag("span.kind", "client")
	span.SetTag("component", "radiant")

	if builder.CustomSpanFunc != nil {
		builder.CustomSpanFunc(span, ctx, inv)
	}
}

func (builder *FilterChainBuilder) operationName(ctx context.Context, inv *orm.Invocation) string {
	if n, ok := ctx.Value(orm.TxNameKey).(string); ok {
		return inv.Method + "#tx(" + n + ")"
	}
	return inv.Method + "#" + inv.GetTableName()
}
