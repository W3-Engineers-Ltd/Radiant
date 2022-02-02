// Package xml for config provider.
//
// depend on github.com/beego/x2j.
//
// go install github.com/beego/x2j.
//
// Usage:
//  import(
//    _ "github.com/W3-Engineers-Ltd/Radiant/core/config/xml"
//      "github.com/W3-Engineers-Ltd/Radiant/core/config"
//  )
//
//  cnf, err := config.NewConfig("xml", "config.xml")
//
// More docs http://radiant.me/docs/module/config.md
package xml

import (
	_ "github.com/W3-Engineers-Ltd/Radiant/core/config/xml"
)
