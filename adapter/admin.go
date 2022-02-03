package adapter

import (
	"time"

	_ "github.com/W3-Engineers-Ltd/Radiant/core/admin"
	"github.com/W3-Engineers-Ltd/Radiant/server/web"
)

// FilterMonitorFunc is default monitor filter when admin module is enable.
// if this func returns, admin module records qps for this request by condition of this function logic.
// usage:
// 	func MyFilterMonitor(method, requestPath string, t time.Duration, pattern string, statusCode int) bool {
//	 	if method == "POST" {
//			return false
//	 	}
//	 	if t.Nanoseconds() < 100 {
//			return false
//	 	}
//	 	if strings.HasPrefix(requestPath, "/astaxie") {
//			return false
//	 	}
//	 	return true
// 	}
// 	radiant.FilterMonitorFunc = MyFilterMonitor.
var FilterMonitorFunc func(string, string, time.Duration, string, int) bool

// PrintTree prints all registered routers.
func PrintTree() M {
	return (M)(web.RadicalApp.PrintTree())
}
