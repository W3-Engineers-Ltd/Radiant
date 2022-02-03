// Copyright 2020 radiant
//

package prometheus

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/W3-Engineers-Ltd/Radiant/client/httplib"
)

type FilterChainBuilder struct {
	AppName    string
	ServerName string
	RunMode    string
}

var (
	summaryVec     prometheus.ObserverVec
	initSummaryVec sync.Once
)

func (builder *FilterChainBuilder) FilterChain(next httplib.Filter) httplib.Filter {
	initSummaryVec.Do(func() {
		summaryVec = prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Name:      "radiant",
			Subsystem: "remote_http_request",
			ConstLabels: map[string]string{
				"server":  builder.ServerName,
				"env":     builder.RunMode,
				"appname": builder.AppName,
			},
			Help: "The statics info for remote http requests",
		}, []string{"proto", "scheme", "method", "host", "path", "status", "isError"})

		prometheus.MustRegister(summaryVec)
	})

	return func(ctx context.Context, req *httplib.RadiantHTTPRequest) (*http.Response, error) {
		startTime := time.Now()
		resp, err := next(ctx, req)
		endTime := time.Now()
		go builder.report(startTime, endTime, ctx, req, resp, err)
		return resp, err
	}
}

func (builder *FilterChainBuilder) report(startTime time.Time, endTime time.Time,
	ctx context.Context, req *httplib.RadiantHTTPRequest, resp *http.Response, err error) {

	proto := req.GetRequest().Proto

	scheme := req.GetRequest().URL.Scheme
	method := req.GetRequest().Method

	host := req.GetRequest().URL.Host
	path := req.GetRequest().URL.Path

	status := -1
	if resp != nil {
		status = resp.StatusCode
	}

	dur := int(endTime.Sub(startTime) / time.Millisecond)

	summaryVec.WithLabelValues(proto, scheme, method, host, path,
		strconv.Itoa(status), strconv.FormatBool(err != nil)).Observe(float64(dur))
}
