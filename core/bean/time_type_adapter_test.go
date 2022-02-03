// Copyright 2020
//

package bean

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeTypeAdapter_DefaultValue(t *testing.T) {
	typeAdapter := &TimeTypeAdapter{Layout: "2006-01-02 15:04:05"}
	tm, err := typeAdapter.DefaultValue(context.Background(), "2018-02-03 12:34:11")
	assert.Nil(t, err)
	assert.NotNil(t, tm)
}
