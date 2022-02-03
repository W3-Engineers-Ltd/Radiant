// Copyright 2020
//

package es

import (
	"fmt"

	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

// IndexNaming generate the index name
type IndexNaming interface {
	IndexName(lm *logs.LogMsg) string
}

var indexNaming IndexNaming = &defaultIndexNaming{}

// SetIndexNaming will register global IndexNaming
func SetIndexNaming(i IndexNaming) {
	indexNaming = i
}

type defaultIndexNaming struct{}

func (d *defaultIndexNaming) IndexName(lm *logs.LogMsg) string {
	return fmt.Sprintf("%04d.%02d.%02d", lm.When.Year(), lm.When.Month(), lm.When.Day())
}
