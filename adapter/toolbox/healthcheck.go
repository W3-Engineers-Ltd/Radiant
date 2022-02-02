// Package toolbox healthcheck
//
// type DatabaseCheck struct {
// }
//
// func (dc *DatabaseCheck) Check() error {
//	if dc.isConnected() {
//		return nil
//	} else {
//		return errors.New("can't connect database")
// 	 }
// }
//
// AddHealthCheck("database",&DatabaseCheck{})
//
// more docs: http://radiant.me/docs/module/toolbox.md
package toolbox

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/admin"
)

// AdminCheckList holds health checker map
// Deprecated using admin.AdminCheckList
var AdminCheckList map[string]HealthChecker

// HealthChecker health checker interface
type HealthChecker admin.HealthChecker

// AddHealthCheck add health checker with name string
func AddHealthCheck(name string, hc HealthChecker) {
	admin.AddHealthCheck(name, hc)
	AdminCheckList[name] = hc
}

func init() {
	AdminCheckList = make(map[string]HealthChecker)
}
