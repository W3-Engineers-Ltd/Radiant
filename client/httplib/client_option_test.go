// Copyright 2021 radiant
//

package httplib

import (
	"errors"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type respCarrier struct {
	bytes []byte
}

func (r *respCarrier) SetBytes(bytes []byte) {
	r.bytes = bytes
}

func (r *respCarrier) String() string {
	return string(r.bytes)
}

func TestOptionWithEnableCookie(t *testing.T) {
	client, err := NewClient("test", "http://httpbin.org/",
		WithEnableCookie(true))
	if err != nil {
		t.Fatal(err)
	}

	v := "smallfish"
	resp := &respCarrier{}
	err = client.Get(resp, "/cookies/set?k1="+v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

	err = client.Get(resp, "/cookies")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

	n := strings.Index(resp.String(), v)
	if n == -1 {
		t.Fatal(v + " not found in cookie")
	}
}

func TestOptionWithUserAgent(t *testing.T) {
	v := "radiant"
	client, err := NewClient("test", "http://httpbin.org/",
		WithUserAgent(v))
	if err != nil {
		t.Fatal(err)
	}

	resp := &respCarrier{}
	err = client.Get(resp, "/headers")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

	n := strings.Index(resp.String(), v)
	if n == -1 {
		t.Fatal(v + " not found in user-agent")
	}
}

func TestOptionWithCheckRedirect(t *testing.T) {
	client, err := NewClient("test", "https://goolnk.com/33BD2j",
		WithCheckRedirect(func(redirectReq *http.Request, redirectVia []*http.Request) error {
			return errors.New("Redirect triggered")
		}))
	if err != nil {
		t.Fatal(err)
	}
	err = client.Get(nil, "")
	assert.NotNil(t, err)
}

func TestOptionWithHTTPSetting(t *testing.T) {
	v := "radiant"
	var setting RadiantHTTPSettings
	setting.EnableCookie = true
	setting.UserAgent = v
	setting.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          50,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	setting.ReadWriteTimeout = 5 * time.Second

	client, err := NewClient("test", "http://httpbin.org/",
		WithHTTPSetting(setting))
	if err != nil {
		t.Fatal(err)
	}

	resp := &respCarrier{}
	err = client.Get(resp, "/get")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

	n := strings.Index(resp.String(), v)
	if n == -1 {
		t.Fatal(v + " not found in user-agent")
	}
}

func TestOptionWithHeader(t *testing.T) {
	client, err := NewClient("test", "http://httpbin.org/")
	if err != nil {
		t.Fatal(err)
	}
	client.CommonOpts = append(client.CommonOpts, WithHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36"))

	resp := &respCarrier{}
	err = client.Get(resp, "/headers")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

	n := strings.Index(resp.String(), "Mozilla/5.0")
	if n == -1 {
		t.Fatal("Mozilla/5.0 not found in user-agent")
	}
}

func TestOptionWithTokenFactory(t *testing.T) {
	client, err := NewClient("test", "http://httpbin.org/")
	if err != nil {
		t.Fatal(err)
	}
	client.CommonOpts = append(client.CommonOpts,
		WithTokenFactory(func() string {
			return "testauth"
		}))

	resp := &respCarrier{}
	err = client.Get(resp, "/headers")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

	n := strings.Index(resp.String(), "testauth")
	if n == -1 {
		t.Fatal("Auth is not set in request")
	}
}

func TestOptionWithBasicAuth(t *testing.T) {
	client, err := NewClient("test", "http://httpbin.org/")
	if err != nil {
		t.Fatal(err)
	}

	resp := &respCarrier{}
	err = client.Get(resp, "/basic-auth/user/passwd",
		WithBasicAuth(func() (string, string) {
			return "user", "passwd"
		}))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
	n := strings.Index(resp.String(), "authenticated")
	if n == -1 {
		t.Fatal("authenticated not found in response")
	}
}

func TestOptionWithContentType(t *testing.T) {
	client, err := NewClient("test", "http://httpbin.org/")
	if err != nil {
		t.Fatal(err)
	}

	v := "application/json"
	resp := &respCarrier{}
	err = client.Get(resp, "/headers", WithContentType(v))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

	n := strings.Index(resp.String(), v)
	if n == -1 {
		t.Fatal(v + " not found in header")
	}
}

func TestOptionWithParam(t *testing.T) {
	client, err := NewClient("test", "http://httpbin.org/")
	if err != nil {
		t.Fatal(err)
	}

	v := "smallfish"
	resp := &respCarrier{}
	err = client.Get(resp, "/get", WithParam("username", v))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

	n := strings.Index(resp.String(), v)
	if n == -1 {
		t.Fatal(v + " not found in header")
	}
}

func TestOptionWithRetry(t *testing.T) {
	client, err := NewClient("test", "https://goolnk.com/33BD2j",
		WithCheckRedirect(func(redirectReq *http.Request, redirectVia []*http.Request) error {
			return errors.New("Redirect triggered")
		}))
	if err != nil {
		t.Fatal(err)
	}

	retryAmount := 1
	retryDelay := 800 * time.Millisecond
	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	_ = client.Get(nil, "", WithRetry(retryAmount, retryDelay))

	endTime := time.Now().UnixNano() / int64(time.Millisecond)
	elapsedTime := endTime - startTime
	delayedTime := int64(retryAmount) * retryDelay.Milliseconds()
	if elapsedTime < delayedTime {
		t.Errorf("Not enough retries. Took %dms. Delay was meant to take %dms", elapsedTime, delayedTime)
	}
}
