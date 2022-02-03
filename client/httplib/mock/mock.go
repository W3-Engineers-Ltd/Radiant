// Copyright 2020 radiant
//

package mock

import (
	"context"
	"net/http"

	"github.com/W3-Engineers-Ltd/Radiant/client/httplib"
	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

const mockCtxKey = "radiant-httplib-mock"

func init() {
	InitMockSetting()
}

type Stub interface {
	Mock(cond RequestCondition, resp *http.Response, err error)
	Clear()
	MockByPath(path string, resp *http.Response, err error)
}

var mockFilter = &MockResponseFilter{}

func InitMockSetting() {
	httplib.AddDefaultFilter(mockFilter.FilterChain)
}

func StartMock() Stub {
	return mockFilter
}

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

type Mock struct {
	cond RequestCondition
	resp *http.Response
	err  error
}

func NewMockByPath(path string, resp *http.Response, err error) *Mock {
	return NewMock(NewSimpleCondition(path), resp, err)
}

func NewMock(con RequestCondition, resp *http.Response, err error) *Mock {
	return &Mock{
		cond: con,
		resp: resp,
		err:  err,
	}
}
