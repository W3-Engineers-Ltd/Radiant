// Copyright 2021 radiant
//

package param

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/context"
)

// Demo is used to test, it's empty
func Demo(i int) {
}

func TestConvertParams(t *testing.T) {
	res := ConvertParams(nil, reflect.TypeOf(Demo), context.NewContext())
	assert.Equal(t, 0, len(res))
	ctx := context.NewContext()
	ctx.Input.RequestBody = []byte("11")
	res = ConvertParams([]*MethodParam{
		New("A", InBody),
	}, reflect.TypeOf(Demo), ctx)
	assert.Equal(t, int64(11), res[0].Int())
}
