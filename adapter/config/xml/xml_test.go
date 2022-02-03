package xml

import (
	"fmt"
	"os"
	"testing"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/config"
)

func TestXML(t *testing.T) {
	var (
		// xml parse should incluce in <config></config> tags
		xmlcontext = `<?xml version="1.0" encoding="UTF-8"?>
<config>
<appname>radicalapi</appname>
<httpport>8080</httpport>
<mysqlport>3600</mysqlport>
<PI>3.1415976</PI>
<runmode>dev</runmode>
<autorender>false</autorender>
<copyrequestbody>true</copyrequestbody>
<path1>${GOPATH}</path1>
<path2>${GOPATH||/home/go}</path2>
<mysection>
<id>1</id>
<name>MySection</name>
</mysection>
</config>
`
		keyValue = map[string]interface{}{
			"appname":         "radicalapi",
			"httpport":        8080,
			"mysqlport":       int64(3600),
			"PI":              3.1415976,
			"runmode":         "dev",
			"autorender":      false,
			"copyrequestbody": true,
			"path1":           os.Getenv("GOPATH"),
			"path2":           os.Getenv("GOPATH"),
			"error":           "",
			"emptystrings":    []string{},
		}
	)

	cfgFileName := "testxml.conf"
	f, err := os.Create(cfgFileName)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(xmlcontext)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove(cfgFileName)

	xmlconf, err := config.NewConfig("xml", cfgFileName)
	if err != nil {
		t.Fatal(err)
	}

	var xmlsection map[string]string
	xmlsection, err = xmlconf.GetSection("mysection")
	if err != nil {
		t.Fatal(err)
	}

	if len(xmlsection) == 0 {
		t.Error("section should not be empty")
	}

	for k, v := range keyValue {

		var (
			value interface{}
			err   error
		)

		switch v.(type) {
		case int:
			value, err = xmlconf.Int(k)
		case int64:
			value, err = xmlconf.Int64(k)
		case float64:
			value, err = xmlconf.Float(k)
		case bool:
			value, err = xmlconf.Bool(k)
		case []string:
			value = xmlconf.Strings(k)
		case string:
			value = xmlconf.String(k)
		default:
			value, err = xmlconf.DIY(k)
		}
		if err != nil {
			t.Errorf("get key %q value fatal,%v err %s", k, v, err)
		} else if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", value) {
			t.Errorf("get key %q value, want %v got %v .", k, v, value)
		}

	}

	if err = xmlconf.Set("name", "astaxie"); err != nil {
		t.Fatal(err)
	}
	if xmlconf.String("name") != "astaxie" {
		t.Fatal("get name error")
	}
}
