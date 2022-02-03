// Copyright 2021 radiant Author. All Rights Reserved.
//

package utils

import "testing"

func TestRand_01(t *testing.T) {
	bs0 := RandomCreateBytes(16)
	bs1 := RandomCreateBytes(16)

	t.Log(string(bs0), string(bs1))
	if string(bs0) == string(bs1) {
		t.FailNow()
	}

	bs0 = RandomCreateBytes(4, []byte(`a`)...)

	if string(bs0) != "aaaa" {
		t.FailNow()
	}
}
