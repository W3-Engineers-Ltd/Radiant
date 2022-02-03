// Copyright 2020 radiant
//

package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoNothingQuerySetter(t *testing.T) {
	setter := &DoNothingQuerySetter{}
	setter.GroupBy().Filter("").Limit(10).
		Distinct().Exclude("a").FilterRaw("", "").
		ForceIndex().ForUpdate().IgnoreIndex().
		Offset(11).OrderBy().RelatedSel().SetCond(nil).UseIndex()

	assert.True(t, setter.Exist())
	err := setter.One(nil)
	assert.Nil(t, err)
	i, err := setter.Count()
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = setter.Delete()
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = setter.All(nil)
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = setter.Update(nil)
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = setter.RowsToMap(nil, "", "")
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = setter.RowsToStruct(nil, "", "")
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = setter.Values(nil)
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = setter.ValuesFlat(nil, "")
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	i, err = setter.ValuesList(nil)
	assert.Equal(t, int64(0), i)
	assert.Nil(t, err)

	ins, err := setter.PrepareInsert()
	assert.Nil(t, err)
	assert.Nil(t, ins)

	assert.NotNil(t, setter.GetCond())
}
