// Copyright 2020
//

package bean

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

func TestDefaultValueFilterChainBuilderFilterChain(t *testing.T) {
	builder := NewDefaultValueFilterChainBuilder(nil, true, true)
	o := orm.NewFilterOrmDecorator(&defaultValueTestOrm{}, builder.FilterChain)

	// test insert
	entity := &DefaultValueTestEntity{}
	_, _ = o.Insert(entity)
	assert.Equal(t, 12, entity.Age)
	assert.Equal(t, 13, entity.AgeInOldStyle)
	assert.Equal(t, 0, entity.AgeIgnore)

	// test InsertOrUpdate
	entity = &DefaultValueTestEntity{}
	orm.RegisterModel(entity)

	_, _ = o.InsertOrUpdate(entity)
	assert.Equal(t, 12, entity.Age)
	assert.Equal(t, 13, entity.AgeInOldStyle)

	// we won't set the default value because we find the pk is not Zero value
	entity.Id = 3
	entity.AgeInOldStyle = 0
	_, _ = o.InsertOrUpdate(entity)
	assert.Equal(t, 0, entity.AgeInOldStyle)

	entity = &DefaultValueTestEntity{}

	// the entity is not array, it will be ignored
	_, _ = o.InsertMulti(3, entity)
	assert.Equal(t, 0, entity.Age)
	assert.Equal(t, 0, entity.AgeInOldStyle)

	_, _ = o.InsertMulti(3, []*DefaultValueTestEntity{entity})
	assert.Equal(t, 12, entity.Age)
	assert.Equal(t, 13, entity.AgeInOldStyle)
}

type defaultValueTestOrm struct {
	orm.DoNothingOrm
}

type DefaultValueTestEntity struct {
	Id            int
	Age           int `default:"12"`
	AgeInOldStyle int `orm:"default(13);radical()"`
	AgeIgnore     int
}
