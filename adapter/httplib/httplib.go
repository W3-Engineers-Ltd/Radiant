// Package httplib is used as http.Client
// Usage:
//
// import "github.com/W3-Engineers-Ltd/Radiant/client/httplib"
//
//	b := httplib.Post("http://radiant.me/")
//	b.Param("username","astaxie")
//	b.Param("password","123456")
//	b.PostFile("uploadfile1", "httplib.pdf")
//	b.PostFile("uploadfile2", "httplib.txt")
//	str, err := b.String()
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(str)
//
//  more docs http://radiant.me/docs/module/httplib.md
package httplib

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/W3-Engineers-Ltd/Radiant/client/httplib"
)

// SetDefaultSetting Overwrite default settings
func SetDefaultSetting(setting RadiantHTTPSettings) {
	httplib.SetDefaultSetting(httplib.RadiantHTTPSettings(setting))
}

// NewRadiantRequest return *RadiantHttpRequest with specific method
func NewRadiantRequest(rawurl, method string) *RadiantHTTPRequest {
	return &RadiantHTTPRequest{
		delegate: httplib.NewRadiantRequest(rawurl, method),
	}
}

// Get returns *RadiantHttpRequest with GET method.
func Get(url string) *RadiantHTTPRequest {
	return NewRadiantRequest(url, "GET")
}

// Post returns *RadiantHttpRequest with POST method.
func Post(url string) *RadiantHTTPRequest {
	return NewRadiantRequest(url, "POST")
}

// Put returns *RadiantHttpRequest with PUT method.
func Put(url string) *RadiantHTTPRequest {
	return NewRadiantRequest(url, "PUT")
}

// Delete returns *RadiantHttpRequest DELETE method.
func Delete(url string) *RadiantHTTPRequest {
	return NewRadiantRequest(url, "DELETE")
}

// Head returns *RadiantHttpRequest with HEAD method.
func Head(url string) *RadiantHTTPRequest {
	return NewRadiantRequest(url, "HEAD")
}

// RadiantHTTPSettings is the http.Client setting
type RadiantHTTPSettings httplib.RadiantHTTPSettings

// RadiantHTTPRequest provides more useful methods for requesting one url than http.Request.
type RadiantHTTPRequest struct {
	delegate *httplib.RadiantHTTPRequest
}

// GetRequest return the request object
func (b *RadiantHTTPRequest) GetRequest() *http.Request {
	return b.delegate.GetRequest()
}

// Setting Change request settings
func (b *RadiantHTTPRequest) Setting(setting RadiantHTTPSettings) *RadiantHTTPRequest {
	b.delegate.Setting(httplib.RadiantHTTPSettings(setting))
	return b
}

// SetBasicAuth sets the request's Authorization header to use HTTP Basic Authentication with the provided username and password.
func (b *RadiantHTTPRequest) SetBasicAuth(username, password string) *RadiantHTTPRequest {
	b.delegate.SetBasicAuth(username, password)
	return b
}

// SetEnableCookie sets enable/disable cookiejar
func (b *RadiantHTTPRequest) SetEnableCookie(enable bool) *RadiantHTTPRequest {
	b.delegate.SetEnableCookie(enable)
	return b
}

// SetUserAgent sets User-Agent header field
func (b *RadiantHTTPRequest) SetUserAgent(useragent string) *RadiantHTTPRequest {
	b.delegate.SetUserAgent(useragent)
	return b
}

// Retries sets Retries times.
// default is 0 means no retried.
// -1 means retried forever.
// others means retried times.
func (b *RadiantHTTPRequest) Retries(times int) *RadiantHTTPRequest {
	b.delegate.Retries(times)
	return b
}

func (b *RadiantHTTPRequest) RetryDelay(delay time.Duration) *RadiantHTTPRequest {
	b.delegate.RetryDelay(delay)
	return b
}

// SetTimeout sets connect time out and read-write time out for RadiantRequest.
func (b *RadiantHTTPRequest) SetTimeout(connectTimeout, readWriteTimeout time.Duration) *RadiantHTTPRequest {
	b.delegate.SetTimeout(connectTimeout, readWriteTimeout)
	return b
}

// SetTLSClientConfig sets tls connection configurations if visiting https url.
func (b *RadiantHTTPRequest) SetTLSClientConfig(config *tls.Config) *RadiantHTTPRequest {
	b.delegate.SetTLSClientConfig(config)
	return b
}

// Header add header item string in request.
func (b *RadiantHTTPRequest) Header(key, value string) *RadiantHTTPRequest {
	b.delegate.Header(key, value)
	return b
}

// SetHost set the request host
func (b *RadiantHTTPRequest) SetHost(host string) *RadiantHTTPRequest {
	b.delegate.SetHost(host)
	return b
}

// SetProtocolVersion Set the protocol version for incoming requests.
// Client requests always use HTTP/1.1.
func (b *RadiantHTTPRequest) SetProtocolVersion(vers string) *RadiantHTTPRequest {
	b.delegate.SetProtocolVersion(vers)
	return b
}

// SetCookie add cookie into request.
func (b *RadiantHTTPRequest) SetCookie(cookie *http.Cookie) *RadiantHTTPRequest {
	b.delegate.SetCookie(cookie)
	return b
}

// SetTransport set the setting transport
func (b *RadiantHTTPRequest) SetTransport(transport http.RoundTripper) *RadiantHTTPRequest {
	b.delegate.SetTransport(transport)
	return b
}

// SetProxy set the http proxy
// example:
//
//	func(req *http.Request) (*url.URL, error) {
// 		u, _ := url.ParseRequestURI("http://127.0.0.1:8118")
// 		return u, nil
// 	}
func (b *RadiantHTTPRequest) SetProxy(proxy func(*http.Request) (*url.URL, error)) *RadiantHTTPRequest {
	b.delegate.SetProxy(proxy)
	return b
}

// SetCheckRedirect specifies the policy for handling redirects.
//
// If CheckRedirect is nil, the Client uses its default policy,
// which is to stop after 10 consecutive requests.
func (b *RadiantHTTPRequest) SetCheckRedirect(redirect func(req *http.Request, via []*http.Request) error) *RadiantHTTPRequest {
	b.delegate.SetCheckRedirect(redirect)
	return b
}

// Param adds query param in to request.
// params build query string as ?key1=value1&key2=value2...
func (b *RadiantHTTPRequest) Param(key, value string) *RadiantHTTPRequest {
	b.delegate.Param(key, value)
	return b
}

// PostFile add a post file to the request
func (b *RadiantHTTPRequest) PostFile(formname, filename string) *RadiantHTTPRequest {
	b.delegate.PostFile(formname, filename)
	return b
}

// Body adds request raw body.
// it supports string and []byte.
func (b *RadiantHTTPRequest) Body(data interface{}) *RadiantHTTPRequest {
	b.delegate.Body(data)
	return b
}

// XMLBody adds request raw body encoding by XML.
func (b *RadiantHTTPRequest) XMLBody(obj interface{}) (*RadiantHTTPRequest, error) {
	_, err := b.delegate.XMLBody(obj)
	return b, err
}

// YAMLBody adds request raw body encoding by YAML.
func (b *RadiantHTTPRequest) YAMLBody(obj interface{}) (*RadiantHTTPRequest, error) {
	_, err := b.delegate.YAMLBody(obj)
	return b, err
}

// JSONBody adds request raw body encoding by JSON.
func (b *RadiantHTTPRequest) JSONBody(obj interface{}) (*RadiantHTTPRequest, error) {
	_, err := b.delegate.JSONBody(obj)
	return b, err
}

// DoRequest will do the client.Do
func (b *RadiantHTTPRequest) DoRequest() (resp *http.Response, err error) {
	return b.delegate.DoRequest()
}

// String returns the body string in response.
// it calls Response inner.
func (b *RadiantHTTPRequest) String() (string, error) {
	return b.delegate.String()
}

// Bytes returns the body []byte in response.
// it calls Response inner.
func (b *RadiantHTTPRequest) Bytes() ([]byte, error) {
	return b.delegate.Bytes()
}

// ToFile saves the body data in response to one file.
// it calls Response inner.
func (b *RadiantHTTPRequest) ToFile(filename string) error {
	return b.delegate.ToFile(filename)
}

// ToJSON returns the map that marshals from the body bytes as json in response .
// it calls Response inner.
func (b *RadiantHTTPRequest) ToJSON(v interface{}) error {
	return b.delegate.ToJSON(v)
}

// ToXML returns the map that marshals from the body bytes as xml in response .
// it calls Response inner.
func (b *RadiantHTTPRequest) ToXML(v interface{}) error {
	return b.delegate.ToXML(v)
}

// ToYAML returns the map that marshals from the body bytes as yaml in response .
// it calls Response inner.
func (b *RadiantHTTPRequest) ToYAML(v interface{}) error {
	return b.delegate.ToYAML(v)
}

// Response executes request client gets response mannually.
func (b *RadiantHTTPRequest) Response() (*http.Response, error) {
	return b.delegate.Response()
}

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return httplib.TimeoutDialer(cTimeout, rwTimeout)
}
