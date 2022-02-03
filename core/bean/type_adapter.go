// Copyright 2020
//

package bean

import (
	"context"
)

// TypeAdapter is an abstraction that define some behavior of target type
// usually, we don't use this to support basic type since golang has many restriction for basic types
// This is an important extension point
type TypeAdapter interface {
	DefaultValue(ctx context.Context, dftValue string) (interface{}, error)
}
