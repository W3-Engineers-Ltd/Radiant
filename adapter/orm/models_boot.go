// Copyright 2014 beego Author. All Rights Reserved.
//

package orm

import (
	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

// RegisterModel register models
func RegisterModel(models ...interface{}) {
	orm.RegisterModel(models...)
}

// RegisterModelWithPrefix register models with a prefix
func RegisterModelWithPrefix(prefix string, models ...interface{}) {
	orm.RegisterModelWithPrefix(prefix, models...)
}

// RegisterModelWithSuffix register models with a suffix
func RegisterModelWithSuffix(suffix string, models ...interface{}) {
	orm.RegisterModelWithSuffix(suffix, models...)
}

// BootStrap bootstrap models.
// make all model parsed and can not add more models
func BootStrap() {
	orm.BootStrap()
}
