// Copyright 2021 radiant
//

package mock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

func TestOrmStub_FilterChain(t *testing.T) {
	os := newOrmStub()
	inv := &orm.Invocation{
		Args: []interface{}{10},
	}
	i := 1
	os.FilterChain(func(ctx context.Context, inv *orm.Invocation) []interface{} {
		i++
		return nil
	})(context.Background(), inv)

	assert.Equal(t, 2, i)

	m := NewMock(NewSimpleCondition("", ""), nil, func(inv *orm.Invocation) {
		arg := inv.Args[0]
		j := arg.(int)
		inv.Args[0] = j + 1
	})
	os.Mock(m)

	os.FilterChain(nil)(context.Background(), inv)
	assert.Equal(t, 11, inv.Args[0])

	inv.Args[0] = 10
	ctxMock := NewMock(NewSimpleCondition("", ""), nil, func(inv *orm.Invocation) {
		arg := inv.Args[0]
		j := arg.(int)
		inv.Args[0] = j + 3
	})

	os.FilterChain(nil)(CtxWithMock(context.Background(), ctxMock), inv)
	assert.Equal(t, 13, inv.Args[0])
}
