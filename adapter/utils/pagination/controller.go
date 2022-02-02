package pagination

import (
	"github.com/W3-Engineers-Ltd/Radiant/adapter/context"
	beecontext "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/pagination"
)

// SetPaginator Instantiates a Paginator and assigns it to context.Input.Data("paginator").
func SetPaginator(ctx *context.Context, per int, nums int64) (paginator *Paginator) {
	return (*Paginator)(pagination.SetPaginator((*beecontext.Context)(ctx), per, nums))
}
