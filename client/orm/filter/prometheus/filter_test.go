// Copyright 2020 radiant
//

package prometheus

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

func TestFilterChainBuilderFilterChain1(t *testing.T) {
	next := func(ctx context.Context, inv *orm.Invocation) []interface{} {
		inv.Method = "coming"
		return []interface{}{}
	}
	builder := &FilterChainBuilder{}
	filter := builder.FilterChain(next)

	assert.NotNil(t, summaryVec)
	assert.NotNil(t, filter)

	inv := &orm.Invocation{}
	filter(context.Background(), inv)
	assert.Equal(t, "coming", inv.Method)

	inv = &orm.Invocation{
		Method:      "Hello",
		TxStartTime: time.Now(),
	}
	builder.reportTxn(context.Background(), inv)

	inv = &orm.Invocation{
		Method: "Begin",
	}

	ctx := context.Background()
	// it will be ignored
	builder.report(ctx, inv, time.Second)

	inv.Method = "Commit"
	builder.report(ctx, inv, time.Second)

	inv.Method = "Update"
	builder.report(ctx, inv, time.Second)
}
