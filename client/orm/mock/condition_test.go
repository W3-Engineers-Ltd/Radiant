// Copyright 2021 radiant
//

package mock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

func TestSimpleCondition_Match(t *testing.T) {
	cond := NewSimpleCondition("", "")
	res := cond.Match(context.Background(), &orm.Invocation{})
	assert.True(t, res)
	cond = NewSimpleCondition("hello", "")
	assert.False(t, cond.Match(context.Background(), &orm.Invocation{}))

	cond = NewSimpleCondition("", "A")
	assert.False(t, cond.Match(context.Background(), &orm.Invocation{
		Method: "B",
	}))

	assert.True(t, cond.Match(context.Background(), &orm.Invocation{
		Method: "A",
	}))
}
