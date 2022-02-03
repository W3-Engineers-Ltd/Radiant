// Copyright 2020 astaxie
//

package metric

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/context"
)

func TestPrometheusMiddleWare(t *testing.T) {
	middleware := PrometheusMiddleWare(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		fmt.Print("you are coming")
	}))
	writer := &context.Response{}
	request := &http.Request{
		URL: &url.URL{
			Host:    "localhost",
			RawPath: "/a/b/c",
		},
		Method: "POST",
	}
	vec := prometheus.NewSummaryVec(prometheus.SummaryOpts{}, []string{"pattern", "method", "status"})

	report(time.Second, writer, request, vec)
	middleware.ServeHTTP(writer, request)
}
