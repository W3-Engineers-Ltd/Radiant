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
func SetDefaultSetting(setting radiantHTTPSettings) {
	httplib.SetDefaultSetting(httplib.radiantHTTPSettings(setting))
}

// NewradiantRequest return *radiantHttpRequest with specific method
func NewradiantRequest(rawurl, method string) *radiantHTTPRequest {
	return &radiantHTTPRequest{
		delegate: httplib.NewradiantRequest(rawurl, method),
	}
}

// Get returns *radiantHttpRequest with GET method.
func Get(url string) *radiantHTTPRequest {
	return NewradiantRequest(url, "GET")
}

// Post returns *radiantHttpRequest with POST method.
func Post(url string) *radiantHTTPRequest {
	return NewradiantRequest(url, "POST")
}

// Put returns *radiantHttpRequest with PUT method.
func Put(url string) *radiantHTTPRequest {
	return NewradiantRequest(url, "PUT")
}

// Delete returns *radiantHttpRequest DELETE method.
func Delete(url string) *radiantHTTPRequest {
	return NewradiantRequest(url, "DELETE")
}

// Head returns *radiantHttpRequest with HEAD method.
func Head(url string) *radiantHTTPRequest {
	return NewradiantRequest(url, "HEAD")
}

// radiantHTTPSettings is the http.Client setting
type radiantHTTPSettings httplib.radiantHTTPSettings

// radiantHTTPRequest provides more useful methods for requesting one url than http.Request.
type radiantHTTPRequest struct {
	delegate *httplib.radiantHTTPRequest
}

// GetRequest return the request object
func (b *radiantHTTPRequest) GetRequest() *http.Request {
	return b.delegate.GetRequest()
}

// Setting Change request settings
func (b *radiantHTTPRequest) Setting(setting radiantHTTPSettings) *radiantHTTPRequest {
	b.delegate.Setting(httplib.radiantHTTPSettings(setting))
	return b
}

// SetBasicAuth sets the request's Authorization header to use HTTP Basic Authentication with the provided username and password.
func (b *radiantHTTPRequest) SetBasicAuth(username, password string) *radiantHTTPRequest {
	b.delegate.SetBasicAuth(username, password)
	return b
}

// SetEnableCookie sets enable/disable cookiejar
func (b *radiantHTTPRequest) SetEnableCookie(enable bool) *radiantHTTPRequest {
	b.delegate.SetEnableCookie(enable)
	return b
}

// SetUserAgent sets User-Agent header field
func (b *radiantHTTPRequest) SetUserAgent(useragent string) *radiantHTTPRequest {
	b.delegate.SetUserAgent(useragent)
	return b
}

// Retries sets Retries times.
// default is 0 means no retried.
// -1 means retried forever.
// others means retried times.
func (b *radiantHTTPRequest) Retries(times int) *radiantHTTPRequest {
	b.delegate.Retries(times)
	return b
}

func (b *radiantHTTPRequest) RetryDelay(delay time.Duration) *radiantHTTPRequest {
	b.delegate.RetryDelay(delay)
	return b
}

// SetTimeout sets connect time out and read-write time out for radiantRequest.
func (b *radiantHTTPRequest) SetTimeout(connectTimeout, readWriteTimeout time.Duration) *radiantHTTPRequest {
	b.delegate.SetTimeout(connectTimeout, readWriteTimeout)
	return b
}

// SetTLSClientConfig sets tls connection configurations if visiting https url.
func (b *radiantHTTPRequest) SetTLSClientConfig(config *tls.Config) *radiantHTTPRequest {
	b.delegate.SetTLSClientConfig(config)
	return b
}

// Header add header item string in request.
func (b *radiantHTTPRequest) Header(key, value string) *radiantHTTPRequest {
	b.delegate.Header(key, value)
	return b
}

// SetHost set the request host
func (b *radiantHTTPRequest) SetHost(host string) *radiantHTTPRequest {
	b.delegate.SetHost(host)
	return b
}

// SetProtocolVersion Set the protocol version for incoming requests.
// Client requests always use HTTP/1.1.
func (b *radiantHTTPRequest) SetProtocolVersion(vers string) *radiantHTTPRequest {
	b.delegate.SetProtocolVersion(vers)
	return b
}

// SetCookie add cookie into request.
func (b *radiantHTTPRequest) SetCookie(cookie *http.Cookie) *radiantHTTPRequest {
	b.delegate.SetCookie(cookie)
	return b
}

// SetTransport set the setting transport
func (b *radiantHTTPRequest) SetTransport(transport http.RoundTripper) *radiantHTTPRequest {
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
func (b *radiantHTTPRequest) SetProxy(proxy func(*http.Request) (*url.URL, error)) *radiantHTTPRequest {
	b.delegate.SetProxy(proxy)
	return b
}

// SetCheckRedirect specifies the policy for handling redirects.
//
// If CheckRedirect is nil, the Client uses its default policy,
// which is to stop after 10 consecutive requests.
func (b *radiantHTTPRequest) SetCheckRedirect(redirect func(req *http.Request, via []*http.Request) error) *radiantHTTPRequest {
	b.delegate.SetCheckRedirect(redirect)
	return b
}

// Param adds query param in to request.
// params build query string as ?key1=value1&key2=value2...
func (b *radiantHTTPRequest) Param(key, value string) *radiantHTTPRequest {
	b.delegate.Param(key, value)
	return b
}

// PostFile add a post file to the request
func (b *radiantHTTPRequest) PostFile(formname, filename string) *radiantHTTPRequest {
	b.delegate.PostFile(formname, filename)
	return b
}

// Body adds request raw body.
// it supports string and []byte.
func (b *radiantHTTPRequest) Body(data interface{}) *radiantHTTPRequest {
	b.delegate.Body(data)
	return b
}

// XMLBody adds request raw body encoding by XML.
func (b *radiantHTTPRequest) XMLBody(obj interface{}) (*radiantHTTPRequest, error) {
	_, err := b.delegate.XMLBody(obj)
	return b, err
}

// YAMLBody adds request raw body encoding by YAML.
func (b *radiantHTTPRequest) YAMLBody(obj interface{}) (*radiantHTTPRequest, error) {
	_, err := b.delegate.YAMLBody(obj)
	return b, err
}

// JSONBody adds request raw body encoding by JSON.
func (b *radiantHTTPRequest) JSONBody(obj interface{}) (*radiantHTTPRequest, error) {
	_, err := b.delegate.JSONBody(obj)
	return b, err
}

// DoRequest will do the client.Do
func (b *radiantHTTPRequest) DoRequest() (resp *http.Response, err error) {
	return b.delegate.DoRequest()
}

// String returns the body string in response.
// it calls Response inner.
func (b *radiantHTTPRequest) String() (string, error) {
	return b.delegate.String()
}

// Bytes returns the body []byte in response.
// it calls Response inner.
func (b *radiantHTTPRequest) Bytes() ([]byte, error) {
	return b.delegate.Bytes()
}

// ToFile saves the body data in response to one file.
// it calls Response inner.
func (b *radiantHTTPRequest) ToFile(filename string) error {
	return b.delegate.ToFile(filename)
}

// ToJSON returns the map that marshals from the body bytes as json in response .
// it calls Response inner.
func (b *radiantHTTPRequest) ToJSON(v interface{}) error {
	return b.delegate.ToJSON(v)
}

// ToXML returns the map that marshals from the body bytes as xml in response .
// it calls Response inner.
func (b *radiantHTTPRequest) ToXML(v interface{}) error {
	return b.delegate.ToXML(v)
}

// ToYAML returns the map that marshals from the body bytes as yaml in response .
// it calls Response inner.
func (b *radiantHTTPRequest) ToYAML(v interface{}) error {
	return b.delegate.ToYAML(v)
}

// Response executes request client gets response mannually.
func (b *radiantHTTPRequest) Response() (*http.Response, error) {
	return b.delegate.Response()
}

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return httplib.TimeoutDialer(cTimeout, rwTimeout)
}
