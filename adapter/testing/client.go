package testing

import "github.com/W3-Engineers-Ltd/Radiant/client/httplib/testing"

// TestHTTPRequest radiant test request client
type TestHTTPRequest testing.TestHTTPRequest

// Get returns test client in GET method
func Get(path string) *TestHTTPRequest {
	return (*TestHTTPRequest)(testing.Get(path))
}

// Post returns test client in POST method
func Post(path string) *TestHTTPRequest {
	return (*TestHTTPRequest)(testing.Post(path))
}

// Put returns test client in PUT method
func Put(path string) *TestHTTPRequest {
	return (*TestHTTPRequest)(testing.Put(path))
}

// Delete returns test client in DELETE method
func Delete(path string) *TestHTTPRequest {
	return (*TestHTTPRequest)(testing.Delete(path))
}

// Head returns test client in HEAD method
func Head(path string) *TestHTTPRequest {
	return (*TestHTTPRequest)(testing.Head(path))
}
