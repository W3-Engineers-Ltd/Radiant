// Copyright 2021 radiant
//

package httplib

import (
	"context"
	"net/http"
)

type FilterChain func(next Filter) Filter

type Filter func(ctx context.Context, req *RadiantHTTPRequest) (*http.Response, error)
