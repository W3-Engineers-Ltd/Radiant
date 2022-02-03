// Copyright 2020
//

package bean

import (
	"context"
	"time"
)

// TimeTypeAdapter process the time.Time
type TimeTypeAdapter struct {
	Layout string
}

// DefaultValue parse the DftValue to time.Time
// and if the DftValue == now
// time.Now() is returned
func (t *TimeTypeAdapter) DefaultValue(ctx context.Context, dftValue string) (interface{}, error) {
	if dftValue == "now" {
		return time.Now(), nil
	}
	return time.Parse(t.Layout, dftValue)
}
