// Copyright 2020 radiant
//

package httplib

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"
)

// RadiantHTTPSettings is the http.Client setting
type RadiantHTTPSettings struct {
	UserAgent        string
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	TLSClientConfig  *tls.Config
	Proxy            func(*http.Request) (*url.URL, error)
	Transport        http.RoundTripper
	CheckRedirect    func(req *http.Request, via []*http.Request) error
	EnableCookie     bool
	Gzip             bool
	Retries          int // if set to -1 means will retry forever
	RetryDelay       time.Duration
	FilterChains     []FilterChain
	EscapeHTML       bool // if set to false means will not escape escape HTML special characters during processing, default true
}

// createDefaultCookie creates a global cookiejar to store cookies.
func createDefaultCookie() {
	settingMutex.Lock()
	defer settingMutex.Unlock()
	defaultCookieJar, _ = cookiejar.New(nil)
}

// SetDefaultSetting overwrites default settings
// Keep in mind that when you invoke the SetDefaultSetting
// some methods invoked before SetDefaultSetting
func SetDefaultSetting(setting RadiantHTTPSettings) {
	settingMutex.Lock()
	defer settingMutex.Unlock()
	defaultSetting = setting
}

// GetDefaultSetting return current default setting
func GetDefaultSetting() RadiantHTTPSettings {
	return defaultSetting
}

var defaultSetting = RadiantHTTPSettings{
	UserAgent:        "radiantServer",
	ConnectTimeout:   60 * time.Second,
	ReadWriteTimeout: 60 * time.Second,
	Gzip:             true,
	FilterChains:     make([]FilterChain, 0, 4),
	EscapeHTML:       true,
}

var (
	defaultCookieJar http.CookieJar
	settingMutex     sync.Mutex
)

// AddDefaultFilter add a new filter into defaultSetting
// Be careful about using this method if you invoke SetDefaultSetting somewhere
func AddDefaultFilter(fc FilterChain) {
	settingMutex.Lock()
	defer settingMutex.Unlock()
	if defaultSetting.FilterChains == nil {
		defaultSetting.FilterChains = make([]FilterChain, 0, 4)
	}
	defaultSetting.FilterChains = append(defaultSetting.FilterChains, fc)
}
