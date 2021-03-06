// Copyright 2020 radiant
//

package log

import (
	"context"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/W3-Engineers-Ltd/Radiant/client/httplib"
	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

// FilterChainBuilder can build a log filter
type FilterChainBuilder struct {
	printableContentTypes []string                              // only print the body of included mime types of request and response
	log                   func(f interface{}, v ...interface{}) // custom log function
}

// BuilderOption option constructor
type BuilderOption func(*FilterChainBuilder)

type logInfo struct {
	req  []byte
	resp []byte
	err  error
}

var defaultprintableContentTypes = []string{
	"text/plain", "text/xml", "text/html", "text/csv",
	"text/calendar", "text/javascript", "text/javascript",
	"text/css",
}

// NewFilterChainBuilder initialize a filterChainBuilder, pass options to customize
func NewFilterChainBuilder(opts ...BuilderOption) *FilterChainBuilder {
	res := &FilterChainBuilder{
		printableContentTypes: defaultprintableContentTypes,
		log:                   logs.Debug,
	}
	for _, o := range opts {
		o(res)
	}

	return res
}

// WithLog return option constructor modify log function
func WithLog(f func(f interface{}, v ...interface{})) BuilderOption {
	return func(h *FilterChainBuilder) {
		h.log = f
	}
}

// WithprintableContentTypes return option constructor modify printableContentTypes
func WithprintableContentTypes(types []string) BuilderOption {
	return func(h *FilterChainBuilder) {
		h.printableContentTypes = types
	}
}

// FilterChain can print the request after FilterChain processing and response before processsing
func (builder *FilterChainBuilder) FilterChain(next httplib.Filter) httplib.Filter {
	return func(ctx context.Context, req *httplib.RadiantHTTPRequest) (*http.Response, error) {
		info := &logInfo{}
		defer info.print(builder.log)
		resp, err := next(ctx, req)
		info.err = err
		contentType := req.GetRequest().Header.Get("Content-Type")
		shouldPrintBody := builder.shouldPrintBody(contentType, req.GetRequest().Body)
		dump, err := httputil.DumpRequest(req.GetRequest(), shouldPrintBody)
		info.req = dump
		if err != nil {
			logs.Error(err)
		}
		if resp != nil {
			contentType = resp.Header.Get("Content-Type")
			shouldPrintBody = builder.shouldPrintBody(contentType, resp.Body)
			dump, err = httputil.DumpResponse(resp, shouldPrintBody)
			info.resp = dump
			if err != nil {
				logs.Error(err)
			}
		}
		return resp, err
	}
}

func (builder *FilterChainBuilder) shouldPrintBody(contentType string, body io.ReadCloser) bool {
	if contains(builder.printableContentTypes, contentType) {
		return true
	}
	if body != nil {
		logs.Warn("printableContentTypes do not contain %s, if you want to print request and response body please add it.", contentType)
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (info *logInfo) print(log func(f interface{}, v ...interface{})) {
	log("Request: ====================")
	log("%q", info.req)
	log("Response: ===================")
	log("%q", info.resp)
	if info.err != nil {
		log("Error: ======================")
		log("%q", info.err)
	}
}
