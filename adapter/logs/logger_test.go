// Copyright 2021 radiant Author. All Rights Reserved.
//

package logs

import (
	"testing"
)

func TestRadicalLoggerInfo(t *testing.T) {
	log := NewLogger(1000)
	log.SetLogger("file", `{"net":"tcp","addr":":7020"}`)
}
