// Copyright 2020 radiant
//

package prometheus

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/client/httplib"
)

func TestFilterChainBuilderFilterChain(t *testing.T) {
	next := func(ctx context.Context, req *httplib.RadiantHTTPRequest) (*http.Response, error) {
		time.Sleep(100 * time.Millisecond)
		return &http.Response{
			StatusCode: 404,
		}, nil
	}
	builder := &FilterChainBuilder{}
	filter := builder.FilterChain(next)
	req := httplib.Get("https://github.com/notifications?query=repo%3Aastaxie%2Fradiant")
	resp, err := filter(context.Background(), req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}
