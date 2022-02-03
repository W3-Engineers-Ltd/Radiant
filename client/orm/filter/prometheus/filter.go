// Copyright 2020 radiant
//

package prometheus

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

// FilterChainBuilder is an extension point,
// when we want to support some configuration,
// please use this structure
// this Filter's behavior looks a little bit strange
// for example:
// if we want to records the metrics of QuerySetter
// actually we only records metrics of invoking "QueryTable"
type FilterChainBuilder struct {
	AppName    string
	ServerName string
	RunMode    string
}

var (
	summaryVec     prometheus.ObserverVec
	initSummaryVec sync.Once
)

func (builder *FilterChainBuilder) FilterChain(next orm.Filter) orm.Filter {
	initSummaryVec.Do(func() {
		summaryVec = prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Name:      "radiant",
			Subsystem: "orm_operation",
			ConstLabels: map[string]string{
				"server":  builder.ServerName,
				"env":     builder.RunMode,
				"appname": builder.AppName,
			},
			Help: "The statics info for orm operation",
		}, []string{"method", "name", "insideTx", "txName"})
		prometheus.MustRegister(summaryVec)
	})

	return func(ctx context.Context, inv *orm.Invocation) []interface{} {
		startTime := time.Now()
		res := next(ctx, inv)
		endTime := time.Now()
		dur := (endTime.Sub(startTime)) / time.Millisecond

		// if the TPS is too large, here may be some problem
		// thinking about using goroutine pool
		go builder.report(ctx, inv, dur)
		return res
	}
}

func (builder *FilterChainBuilder) report(ctx context.Context, inv *orm.Invocation, dur time.Duration) {
	// start a transaction, we don't record it
	if strings.HasPrefix(inv.Method, "Begin") {
		return
	}
	if inv.Method == "Commit" || inv.Method == "Rollback" {
		builder.reportTxn(ctx, inv)
		return
	}
	summaryVec.WithLabelValues(inv.Method, inv.GetTableName(),
		strconv.FormatBool(inv.InsideTx), inv.TxName).Observe(float64(dur))
}

func (builder *FilterChainBuilder) reportTxn(ctx context.Context, inv *orm.Invocation) {
	dur := time.Since(inv.TxStartTime) / time.Millisecond
	summaryVec.WithLabelValues(inv.Method, inv.TxName,
		strconv.FormatBool(inv.InsideTx), inv.TxName).Observe(float64(dur))
}
