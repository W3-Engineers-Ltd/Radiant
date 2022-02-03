// Copyright 2021 radiant
//

package mock

import (
	"context"

	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

type mockCtxKeyType string

const mockCtxKey = mockCtxKeyType("radiant-orm-mock")

func CtxWithMock(ctx context.Context, mock ...*Mock) context.Context {
	return context.WithValue(ctx, mockCtxKey, mock)
}

func mockFromCtx(ctx context.Context) []*Mock {
	ms := ctx.Value(mockCtxKey)
	if ms != nil {
		if res, ok := ms.([]*Mock); ok {
			return res
		}
		logs.Error("mockCtxKey found in context, but value is not type []*Mock")
	}
	return nil
}
