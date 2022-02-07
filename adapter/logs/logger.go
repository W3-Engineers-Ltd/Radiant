// Copyright 2014 beego Author. All Rights Reserved.
//

package logs

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

// ColorByStatus return color by http code
// 2xx return Green
// 3xx return White
// 4xx return Yellow
// 5xx return Red
func ColorByStatus(code int) string {
	return logs.ColorByStatus(code)
}

// ColorByMethod return color by http code
func ColorByMethod(method string) string {
	return logs.ColorByMethod(method)
}

// ResetColor return reset color
func ResetColor() string {
	return logs.ResetColor()
}
