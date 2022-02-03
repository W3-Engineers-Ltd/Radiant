// Copyright 2021 radiant
//

package mock

import (
	"context"
	"net/http"

	"github.com/W3-Engineers-Ltd/Radiant/client/httplib"
)

// MockResponse will return mock response if find any suitable mock data
// if you want to test your code using httplib, you need this.
type MockResponseFilter struct {
	ms []*Mock
}

func NewMockResponseFilter() *MockResponseFilter {
	return &MockResponseFilter{
		ms: make([]*Mock, 0, 1),
	}
}

func (m *MockResponseFilter) FilterChain(next httplib.Filter) httplib.Filter {
	return func(ctx context.Context, req *httplib.RadiantHTTPRequest) (*http.Response, error) {
		ms := mockFromCtx(ctx)
		ms = append(ms, m.ms...)
		for _, mock := range ms {
			if mock.cond.Match(ctx, req) {
				return mock.resp, mock.err
			}
		}
		return next(ctx, req)
	}
}

func (m *MockResponseFilter) MockByPath(path string, resp *http.Response, err error) {
	m.Mock(NewSimpleCondition(path), resp, err)
}

func (m *MockResponseFilter) Clear() {
	m.ms = make([]*Mock, 0, 1)
}

// Mock add mock data
// If the cond.Match(...) = true, the resp and err will be returned
func (m *MockResponseFilter) Mock(cond RequestCondition, resp *http.Response, err error) {
	m.ms = append(m.ms, NewMock(cond, resp, err))
}
