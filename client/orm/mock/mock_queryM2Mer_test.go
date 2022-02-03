// Copyright 2020 radiant
//

package mock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

func TestDoNothingQueryM2Mer(t *testing.T) {
	m2m := &DoNothingQueryM2Mer{}

	i, err := m2m.Clear()
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = m2m.Count()
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = m2m.Add()
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = m2m.Remove()
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	assert.True(t, m2m.Exist(nil))
}

func TestNewQueryM2MerCondition(t *testing.T) {
	cond := NewQueryM2MerCondition("", "")
	res := cond.Match(context.Background(), &orm.Invocation{})
	assert.True(t, res)
	cond = NewQueryM2MerCondition("hello", "")
	assert.False(t, cond.Match(context.Background(), &orm.Invocation{}))

	cond = NewQueryM2MerCondition("", "A")
	assert.False(t, cond.Match(context.Background(), &orm.Invocation{
		Args: []interface{}{0, "B"},
	}))

	assert.True(t, cond.Match(context.Background(), &orm.Invocation{
		Args: []interface{}{0, "A"},
	}))
}
