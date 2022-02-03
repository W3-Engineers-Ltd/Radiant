// Copyright 2020 radiant
//

package orm

import (
	"context"
)

// FilterChain is used to build a Filter
// don't forget to call next(...) inside your Filter
type FilterChain func(next Filter) Filter

// Filter's behavior is a little big strange.
// it's only be called when users call methods of Ormer
// return value is an array. it's a little bit hard to understand,
// for example, the Ormer's Read method only return error
// so the filter processing this method should return an array whose first element is error
// and, Ormer's ReadOrCreateWithCtx return three values, so the Filter's result should contains three values
type Filter func(ctx context.Context, inv *Invocation) []interface{}

var globalFilterChains = make([]FilterChain, 0, 4)

// AddGlobalFilterChain adds a new FilterChain
// All orm instances built after this invocation will use this filterChain,
// but instances built before this invocation will not be affected
func AddGlobalFilterChain(filterChain ...FilterChain) {
	globalFilterChains = append(globalFilterChains, filterChain...)
}
