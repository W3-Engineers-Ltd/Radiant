// Copyright 2021 radiant
//

package mock

import (
	"net/http"

	radiantCtx "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

func NewMockContext(req *http.Request) (*radiantCtx.Context, *HttpResponse) {
	ctx := radiantCtx.NewContext()
	resp := NewMockHttpResponse()
	ctx.Reset(resp, req)
	return ctx, resp
}
