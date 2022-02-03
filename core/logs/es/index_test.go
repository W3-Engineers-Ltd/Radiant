// Copyright 2020
//

package es

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

func TestDefaultIndexNaming_IndexName(t *testing.T) {
	tm := time.Date(2020, 9, 12, 1, 34, 45, 234, time.UTC)
	lm := &logs.LogMsg{
		When: tm,
	}

	res := (&defaultIndexNaming{}).IndexName(lm)
	assert.Equal(t, "2020.09.12", res)
}
