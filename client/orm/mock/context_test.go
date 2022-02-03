// Copyright 2021 radiant
//

package mock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCtx(t *testing.T) {
	ms := make([]*Mock, 0, 4)
	ctx := CtxWithMock(context.Background(), ms...)
	res := mockFromCtx(ctx)
	assert.Equal(t, ms, res)
}
