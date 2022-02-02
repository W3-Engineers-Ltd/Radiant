package session

import (
	"testing"
)

func TestGob(t *testing.T) {
	a := make(map[interface{}]interface{})
	a["username"] = "astaxie"
	a[12] = 234
	a["user"] = User{"asta", "xie"}
	b, err := EncodeGob(a)
	if err != nil {
		t.Error(err)
	}
	c, err := DecodeGob(b)
	if err != nil {
		t.Error(err)
	}
	if len(c) == 0 {
		t.Error("decodeGob empty")
	}
	if c["username"] != "astaxie" {
		t.Error("decode string error")
	}
	if c[12] != 234 {
		t.Error("decode int error")
	}
	if c["user"].(User).Username != "asta" {
		t.Error("decode struct error")
	}
}

type User struct {
	Username string
	NickName string
}
