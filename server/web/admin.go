package web

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

// RadicalAdminApp is the default adminApp used by admin module.
var radicalAdminApp *adminApp

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

func init() {
	FilterMonitorFunc = func(string, string, time.Duration, string, int) bool { return true }
}

func list(root string, p interface{}, m M) {
	pt := reflect.TypeOf(p)
	pv := reflect.ValueOf(p)
	if pt.Kind() == reflect.Ptr {
		pt = pt.Elem()
		pv = pv.Elem()
	}
	for i := 0; i < pv.NumField(); i++ {
		var key string
		if root == "" {
			key = pt.Field(i).Name
		} else {
			key = root + "." + pt.Field(i).Name
		}
		if pv.Field(i).Kind() == reflect.Struct {
			list(key, pv.Field(i).Interface(), m)
		} else {
			m[key] = pv.Field(i).Interface()
		}
	}
}

func writeJSON(rw http.ResponseWriter, jsonData []byte) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}

// adminApp is an http.HandlerFunc map used as radicalAdminApp.
type adminApp struct {
	*HttpServer
}

// Run start Radiant admin
func (admin *adminApp) Run() {
	logs.Debug("now we don't start tasks here, if you use task module," +
		" please invoke task.StartTask, or task will not be executed")
	addr := BConfig.Listen.AdminAddr
	if BConfig.Listen.AdminPort != 0 {
		addr = fmt.Sprintf("%s:%d", BConfig.Listen.AdminAddr, BConfig.Listen.AdminPort)
	}
	logs.Info("Admin server Running on %s", addr)
	admin.HttpServer.Run(addr)
}

func registerAdmin() error {
	if BConfig.Listen.EnableAdmin {

		c := &adminController{
			servers: make([]*HttpServer, 0, 2),
		}

		// copy config to avoid conflict
		adminCfg := *BConfig
		radicalAdminApp = &adminApp{
			HttpServer: NewHttpServerWithCfg(&adminCfg),
		}
		// keep in mind that all data should be html escaped to avoid XSS attack
		radicalAdminApp.Router("/", c, "get:AdminIndex")
		radicalAdminApp.Router("/qps", c, "get:QpsIndex")
		radicalAdminApp.Router("/prof", c, "get:ProfIndex")
		radicalAdminApp.Router("/healthcheck", c, "get:Healthcheck")
		radicalAdminApp.Router("/task", c, "get:TaskStatus")
		radicalAdminApp.Router("/listconf", c, "get:ListConf")
		radicalAdminApp.Router("/metrics", c, "get:PrometheusMetrics")

		go radicalAdminApp.Run()
	}
	return nil
}
