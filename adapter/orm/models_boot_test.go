// Copyright 2021 radiant
//

package orm

import (
	"testing"
)

type User struct {
	Id int
}

type Seller struct {
	Id int
}

func TestRegisterModelWithPrefix(t *testing.T) {
	RegisterModelWithPrefix("test", &User{}, &Seller{})
}
