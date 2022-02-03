// Copyright 2020 radiant
//

package mock

import (
	"context"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

var stub = newOrmStub()

func init() {
	orm.AddGlobalFilterChain(stub.FilterChain)
}

type Stub interface {
	Mock(m *Mock)
	Clear()
}

type OrmStub struct {
	ms []*Mock
}

func StartMock() Stub {
	return stub
}

func newOrmStub() *OrmStub {
	return &OrmStub{
		ms: make([]*Mock, 0, 4),
	}
}

func (o *OrmStub) Mock(m *Mock) {
	o.ms = append(o.ms, m)
}

func (o *OrmStub) Clear() {
	o.ms = make([]*Mock, 0, 4)
}

func (o *OrmStub) FilterChain(next orm.Filter) orm.Filter {
	return func(ctx context.Context, inv *orm.Invocation) []interface{} {
		ms := mockFromCtx(ctx)
		ms = append(ms, o.ms...)

		for _, mock := range ms {
			if mock.cond.Match(ctx, inv) {
				if mock.cb != nil {
					mock.cb(inv)
				}
				return mock.resp
			}
		}
		return next(ctx, inv)
	}
}
