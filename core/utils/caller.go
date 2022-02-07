// Copyright 2014 beego Author. All Rights Reserved.
//

package utils

import (
	"reflect"
	"runtime"
)

// GetFuncName get function name
func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
