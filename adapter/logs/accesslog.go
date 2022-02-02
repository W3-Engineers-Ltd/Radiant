package logs

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

// AccessLogRecord struct for holding access log data.
type AccessLogRecord logs.AccessLogRecord

// AccessLog - Format and print access log.
func AccessLog(r *AccessLogRecord, format string) {
	logs.AccessLog((*logs.AccessLogRecord)(r), format)
}
