// Copyright 2014 beego Author. All Rights Reserved.
//

package pagination

import (
	"fmt"
	"reflect"
)

// ToInt64 convert any numeric value to int64
func toInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}
