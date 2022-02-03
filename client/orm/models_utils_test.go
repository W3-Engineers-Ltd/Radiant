// Copyright 2020
//

package orm

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type NotApplicableModel struct {
	Id int
}

func (n *NotApplicableModel) IsApplicableTableForDB(db string) bool {
	return db == "default"
}

func TestIsApplicableTableForDB(t *testing.T) {
	assert.False(t, isApplicableTableForDB(reflect.ValueOf(&NotApplicableModel{}), "defa"))
	assert.True(t, isApplicableTableForDB(reflect.ValueOf(&NotApplicableModel{}), "default"))
}
