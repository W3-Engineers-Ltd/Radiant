// Copyright 2020 radiant
//

package mock

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/client/httplib"
)

func TestStartMock(t *testing.T) {
	// httplib.defaultSetting.FilterChains = []httplib.FilterChain{mockFilter.FilterChain}

	stub := StartMock()
	// defer stub.Clear()

	expectedResp := httplib.NewHttpResponseWithJsonBody([]byte(`{}`))
	expectedErr := errors.New("expected err")

	stub.Mock(NewSimpleCondition("/abc"), expectedResp, expectedErr)

	resp, err := OriginalCodeUsingHttplib()

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, expectedResp, resp)
}

// TestStartMock_Isolation Test StartMock that
// mock only work for this request
func TestStartMockIsolation(t *testing.T) {
	// httplib.defaultSetting.FilterChains = []httplib.FilterChain{mockFilter.FilterChain}
	// setup global stub
	stub := StartMock()
	globalMockResp := httplib.NewHttpResponseWithJsonBody([]byte(`{}`))
	globalMockErr := errors.New("expected err")
	stub.Mock(NewSimpleCondition("/abc"), globalMockResp, globalMockErr)

	expectedResp := httplib.NewHttpResponseWithJsonBody(struct {
		A string `json:"a"`
	}{
		A: "aaa",
	})
	expectedErr := errors.New("expected err aa")
	m := NewMockByPath("/abc", expectedResp, expectedErr)
	ctx := CtxWithMock(context.Background(), m)

	resp, err := OriginnalCodeUsingHttplibPassCtx(ctx)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, expectedResp, resp)
}

func OriginnalCodeUsingHttplibPassCtx(ctx context.Context) (*http.Response, error) {
	return httplib.Get("http://localhost:7777/abc").DoRequestWithCtx(ctx)
}

func OriginalCodeUsingHttplib() (*http.Response, error) {
	return httplib.Get("http://localhost:7777/abc").DoRequest()
}
