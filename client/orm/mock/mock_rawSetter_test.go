// Copyright 2020 radiant
//

package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoNothingRawSetter(t *testing.T) {
	rs := &DoNothingRawSetter{}
	i, err := rs.ValuesList(nil)
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = rs.Values(nil)
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = rs.ValuesFlat(nil)
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = rs.RowsToStruct(nil, "", "")
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = rs.RowsToMap(nil, "", "")
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = rs.QueryRows()
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	err = rs.QueryRow()
	// assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	s, err := rs.Exec()
	assert.Nil(t, err)
	assert.Nil(t, s)

	p, err := rs.Prepare()
	assert.Nil(t, err)
	assert.Nil(t, p)

	rrs := rs.SetArgs()
	assert.Equal(t, rrs, rs)
}
