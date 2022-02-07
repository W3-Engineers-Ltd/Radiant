// Copyright 2014 beego Author. All Rights Reserved.
//

package config

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/config"
)

// NewFakeConfig return a fake Configer
func NewFakeConfig() Configer {
	config := config.NewFakeConfig()
	return &newToOldConfigerAdapter{delegate: config}
}
