// Package yaml for config provider
//
// depend on github.com/radiant/goyaml2
//
// go install github.com/radiant/goyaml2
//
// Usage:
//  import(
//   _ "github.com/W3-Engineers-Ltd/Radiant/core/config/yaml"
//     "github.com/W3-Engineers-Ltd/Radiant/core/config"
//  )
//
//  cnf, err := config.NewConfig("yaml", "config.yaml")
//
// More docs http://radiant.me/docs/module/config.md
package yaml

import (
	_ "github.com/W3-Engineers-Ltd/Radiant/core/config/yaml"
)
