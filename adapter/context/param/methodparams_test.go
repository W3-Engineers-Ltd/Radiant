// Copyright 2020 radiant
//

package param

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethodParamString(t *testing.T) {
	method := New("myName", IsRequired, InHeader, Default("abc"))
	s := method.String()
	assert.Equal(t, `param.New("myName", param.IsRequired, param.InHeader, param.Default("abc"))`, s)
}

func TestMake(t *testing.T) {
	res := Make()
	assert.Equal(t, 0, len(res))
	res = Make(New("myName", InBody))
	assert.Equal(t, 1, len(res))
}
