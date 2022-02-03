package yaml

import (
	"fmt"
	"os"
	"testing"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/config"
)

func TestYaml(t *testing.T) {
	var (
		yamlcontext = `
"appname": radicalapi
"httpport": 8080
"mysqlport": 3600
"PI": 3.1415976
"runmode": dev
"autorender": false
"copyrequestbody": true
"PATH": GOPATH
"path1": ${GOPATH}
"path2": ${GOPATH||/home/go}
"empty": "" 
`

		keyValue = map[string]interface{}{
			"appname":         "radicalapi",
			"httpport":        8080,
			"mysqlport":       int64(3600),
			"PI":              3.1415976,
			"runmode":         "dev",
			"autorender":      false,
			"copyrequestbody": true,
			"PATH":            "GOPATH",
			"path1":           os.Getenv("GOPATH"),
			"path2":           os.Getenv("GOPATH"),
			"error":           "",
			"emptystrings":    []string{},
		}
	)
	cfgFileName := "testyaml.conf"
	f, err := os.Create(cfgFileName)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(yamlcontext)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove(cfgFileName)
	yamlconf, err := config.NewConfig("yaml", cfgFileName)
	if err != nil {
		t.Fatal(err)
	}

	if yamlconf.String("appname") != "radicalapi" {
		t.Fatal("appname not equal to radicalapi")
	}

	for k, v := range keyValue {

		var (
			value interface{}
			err   error
		)

		switch v.(type) {
		case int:
			value, err = yamlconf.Int(k)
		case int64:
			value, err = yamlconf.Int64(k)
		case float64:
			value, err = yamlconf.Float(k)
		case bool:
			value, err = yamlconf.Bool(k)
		case []string:
			value = yamlconf.Strings(k)
		case string:
			value = yamlconf.String(k)
		default:
			value, err = yamlconf.DIY(k)
		}
		if err != nil {
			t.Errorf("get key %q value fatal,%v err %s", k, v, err)
		} else if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", value) {
			t.Errorf("get key %q value, want %v got %v .", k, v, value)
		}

	}

	if err = yamlconf.Set("name", "astaxie"); err != nil {
		t.Fatal(err)
	}
	if yamlconf.String("name") != "astaxie" {
		t.Fatal("get name error")
	}
}
