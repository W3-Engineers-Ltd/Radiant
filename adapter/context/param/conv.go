package param

import (
	"reflect"

	radicalcontext "github.com/W3-Engineers-Ltd/Radiant/adapter/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/context/param"
)

// ConvertParams converts http method params to values that will be passed to the method controller as arguments
func ConvertParams(methodParams []*MethodParam, methodType reflect.Type, ctx *radicalcontext.Context) (result []reflect.Value) {
	nps := make([]*param.MethodParam, 0, len(methodParams))
	for _, mp := range methodParams {
		nps = append(nps, (*param.MethodParam)(mp))
	}
	return param.ConvertParams(nps, methodType, (*context.Context)(ctx))
}
