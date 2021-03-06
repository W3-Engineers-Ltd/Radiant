// Copyright 2020 radiant
//

package log

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/client/httplib"
)

func TestFilterChain(t *testing.T) {
	next := func(ctx context.Context, req *httplib.RadiantHTTPRequest) (*http.Response, error) {
		time.Sleep(100 * time.Millisecond)
		return &http.Response{
			StatusCode: 404,
		}, nil
	}
	builder := NewFilterChainBuilder()
	filter := builder.FilterChain(next)
	req := httplib.Get("https://github.com/notifications?query=repo%3Aastaxie%2Fradiant")
	resp, err := filter(context.Background(), req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestContains(t *testing.T) {
	jsonType := "application/json"
	cases := []struct {
		Name        string
		Types       []string
		ContentType string
		Expected    bool
	}{
		{"case1", []string{jsonType}, jsonType, true},
		{"case2", []string{"text/plain"}, jsonType, false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := contains(c.Types, c.ContentType); ans != c.Expected {
				t.Fatalf("Types: %v, ContentType: %v, expected %v, but %v got",
					c.Types, c.ContentType, c.Expected, ans)
			}
		})
	}
}
