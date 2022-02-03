// Copyright 2020
//

package bean

import (
	"context"
)

// AutoWireBeanFactory wire the bean based on ApplicationContext and context.Context
type AutoWireBeanFactory interface {
	// AutoWire will wire the bean.
	AutoWire(ctx context.Context, appCtx ApplicationContext, bean interface{}) error
}
