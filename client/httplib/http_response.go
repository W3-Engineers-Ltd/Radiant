// Copyright 2021 radiant
//

package httplib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// NewHttpResponseWithJsonBody will try to convert the data to json format
// usually you only use this when you want to mock http Response
func NewHttpResponseWithJsonBody(data interface{}) *http.Response {
	var body []byte
	if str, ok := data.(string); ok {
		body = []byte(str)
	} else if bts, ok := data.([]byte); ok {
		body = bts
	} else {
		body, _ = json.Marshal(data)
	}
	return &http.Response{
		ContentLength: int64(len(body)),
		Body:          ioutil.NopCloser(bytes.NewReader(body)),
	}
}
