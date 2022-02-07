// Copyright 2014 beego Author. All Rights Reserved.
//

package session

import (
	"github.com/W3-Engineers-Ltd/Radiant/server/web/session"
)

// EncodeGob encode the obj to gob
func EncodeGob(obj map[interface{}]interface{}) ([]byte, error) {
	return session.EncodeGob(obj)
}

// DecodeGob decode data to map
func DecodeGob(encoded []byte) (map[interface{}]interface{}, error) {
	return session.DecodeGob(encoded)
}
