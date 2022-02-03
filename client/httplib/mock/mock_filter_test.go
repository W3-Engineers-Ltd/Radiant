// Copyright 2020 radiant
//

package mock

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/client/httplib"
)

func TestMockResponseFilterFilterChain(t *testing.T) {
	req := httplib.Get("http://localhost:8080/abc/s")
	ft := NewMockResponseFilter()

	expectedResp := httplib.NewHttpResponseWithJsonBody(`{}`)
	expectedErr := errors.New("expected error")
	ft.Mock(NewSimpleCondition("/abc/s"), expectedResp, expectedErr)

	req.AddFilters(ft.FilterChain)

	resp, err := req.DoRequest()
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, expectedResp, resp)

	req = httplib.Get("http://localhost:8080/abcd/s")
	req.AddFilters(ft.FilterChain)

	resp, err = req.DoRequest()
	assert.NotEqual(t, expectedErr, err)
	assert.NotEqual(t, expectedResp, resp)

	req = httplib.Get("http://localhost:8080/abc/s")
	req.AddFilters(ft.FilterChain)
	expectedResp1 := httplib.NewHttpResponseWithJsonBody(map[string]string{})
	expectedErr1 := errors.New("expected error")
	ft.Mock(NewSimpleCondition("/abc/abs/bbc"), expectedResp1, expectedErr1)

	resp, err = req.DoRequest()
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, expectedResp, resp)

	req = httplib.Get("http://localhost:8080/abc/abs/bbc")
	req.AddFilters(ft.FilterChain)
	ft.Mock(NewSimpleCondition("/abc/abs/bbc"), expectedResp1, expectedErr1)
	resp, err = req.DoRequest()
	assert.Equal(t, expectedErr1, err)
	assert.Equal(t, expectedResp1, resp)
}
