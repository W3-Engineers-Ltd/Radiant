// Copyright 2014 beego Author. All Rights Reserved.
//

package utils

import (
	"strings"
	"testing"
)

func TestGetFuncName(t *testing.T) {
	name := GetFuncName(TestGetFuncName)
	t.Log(name)
	if !strings.HasSuffix(name, ".TestGetFuncName") {
		t.Error("get func name error")
	}
}
