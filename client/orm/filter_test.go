// Copyright 2020 radiant
//

package orm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGlobalFilterChain(t *testing.T) {
	AddGlobalFilterChain(func(next Filter) Filter {
		return func(ctx context.Context, inv *Invocation) []interface{} {
			return next(ctx, inv)
		}
	})
	assert.Equal(t, 1, len(globalFilterChains))
	globalFilterChains = nil
}
