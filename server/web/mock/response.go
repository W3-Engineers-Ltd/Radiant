// Copyright 2021 radiant
//

package mock

import (
	"encoding/json"
	"net/http"
)

// HttpResponse mock response, which should be used in tests
type HttpResponse struct {
	body       []byte
	header     http.Header
	StatusCode int
}

// NewMockHttpResponse you should only use this in your test code
func NewMockHttpResponse() *HttpResponse {
	return &HttpResponse{
		body:   make([]byte, 0),
		header: make(http.Header),
	}
}

// Header return headers
func (m *HttpResponse) Header() http.Header {
	return m.header
}

// Write append the body
func (m *HttpResponse) Write(bytes []byte) (int, error) {
	m.body = append(m.body, bytes...)
	return len(bytes), nil
}

// WriteHeader set the status code
func (m *HttpResponse) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}

// JsonUnmarshal convert the body to object
func (m *HttpResponse) JsonUnmarshal(value interface{}) error {
	return json.Unmarshal(m.body, value)
}

// BodyToString return the body as the string
func (m *HttpResponse) BodyToString() string {
	return string(m.body)
}

// Reset will reset the status to init status
// Usually, you want to reuse this instance you may need to call Reset
func (m *HttpResponse) Reset() {
	m.body = make([]byte, 0)
	m.header = make(http.Header)
	m.StatusCode = 0
}
