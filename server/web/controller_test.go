package web

import (
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

func TestGetInt(t *testing.T) {
	i := context.NewInput()
	i.SetParam("age", "40")
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}
	val, _ := ctrlr.GetInt("age")
	if val != 40 {
		t.Errorf("TestGetInt expect 40,get %T,%v", val, val)
	}
}

func TestGetInt8(t *testing.T) {
	i := context.NewInput()
	i.SetParam("age", "40")
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}
	val, _ := ctrlr.GetInt8("age")
	if val != 40 {
		t.Errorf("TestGetInt8 expect 40,get %T,%v", val, val)
	}
	// Output: int8
}

func TestGetInt16(t *testing.T) {
	i := context.NewInput()
	i.SetParam("age", "40")
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}
	val, _ := ctrlr.GetInt16("age")
	if val != 40 {
		t.Errorf("TestGetInt16 expect 40,get %T,%v", val, val)
	}
}

func TestGetInt32(t *testing.T) {
	i := context.NewInput()
	i.SetParam("age", "40")
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}
	val, _ := ctrlr.GetInt32("age")
	if val != 40 {
		t.Errorf("TestGetInt32 expect 40,get %T,%v", val, val)
	}
}

func TestGetInt64(t *testing.T) {
	i := context.NewInput()
	i.SetParam("age", "40")
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}
	val, _ := ctrlr.GetInt64("age")
	if val != 40 {
		t.Errorf("TestGeetInt64 expect 40,get %T,%v", val, val)
	}
}

func TestGetUint8(t *testing.T) {
	i := context.NewInput()
	i.SetParam("age", strconv.FormatUint(math.MaxUint8, 10))
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}
	val, _ := ctrlr.GetUint8("age")
	if val != math.MaxUint8 {
		t.Errorf("TestGetUint8 expect %v,get %T,%v", math.MaxUint8, val, val)
	}
}

func TestGetUint16(t *testing.T) {
	i := context.NewInput()
	i.SetParam("age", strconv.FormatUint(math.MaxUint16, 10))
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}
	val, _ := ctrlr.GetUint16("age")
	if val != math.MaxUint16 {
		t.Errorf("TestGetUint16 expect %v,get %T,%v", math.MaxUint16, val, val)
	}
}

func TestGetUint32(t *testing.T) {
	i := context.NewInput()
	i.SetParam("age", strconv.FormatUint(math.MaxUint32, 10))
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}
	val, _ := ctrlr.GetUint32("age")
	if val != math.MaxUint32 {
		t.Errorf("TestGetUint32 expect %v,get %T,%v", math.MaxUint32, val, val)
	}
}

func TestGetUint64(t *testing.T) {
	i := context.NewInput()
	i.SetParam("age", strconv.FormatUint(math.MaxUint64, 10))
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}
	val, _ := ctrlr.GetUint64("age")
	if val != math.MaxUint64 {
		t.Errorf("TestGetUint64 expect %v,get %T,%v", uint64(math.MaxUint64), val, val)
	}
}

func TestAdditionalViewPaths(t *testing.T) {
	tmpDir := os.TempDir()
	dir1 := filepath.Join(tmpDir, "_beeTmp", "TestAdditionalViewPaths")
	dir2 := filepath.Join(tmpDir, "_beeTmp2", "TestAdditionalViewPaths")
	defer os.RemoveAll(dir1)
	defer os.RemoveAll(dir2)

	dir1file := "file1.tpl"
	dir2file := "file2.tpl"

	genFile := func(dir string, name string, content string) {
		os.MkdirAll(filepath.Dir(filepath.Join(dir, name)), 0o777)
		if f, err := os.Create(filepath.Join(dir, name)); err != nil {
			t.Fatal(err)
		} else {
			defer f.Close()
			f.WriteString(content)
			f.Close()
		}
	}
	genFile(dir1, dir1file, `<div>{{.Content}}</div>`)
	genFile(dir2, dir2file, `<html>{{.Content}}</html>`)

	AddViewPath(dir1)
	AddViewPath(dir2)

	ctrl := Controller{
		TplName:  "file1.tpl",
		ViewPath: dir1,
	}
	ctrl.Data = map[interface{}]interface{}{
		"Content": "value2",
	}
	if result, err := ctrl.RenderString(); err != nil {
		t.Fatal(err)
	} else {
		if result != "<div>value2</div>" {
			t.Fatalf("TestAdditionalViewPaths expect %s got %s", "<div>value2</div>", result)
		}
	}

	func() {
		ctrl.TplName = "file2.tpl"
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("TestAdditionalViewPaths expected error")
			}
		}()
		ctrl.RenderString()
	}()

	ctrl.TplName = "file2.tpl"
	ctrl.ViewPath = dir2
	ctrl.RenderString()
}

func TestBindJson(t *testing.T) {
	var s struct {
		Foo string `json:"foo"`
	}
	header := map[string][]string{"Content-Type": {"application/json"}}
	request := &http.Request{Header: header}
	input := &context.radiantInput{RequestBody: []byte(`{"foo": "FOO"}`)}
	ctx := &context.Context{Request: request, Input: input}
	ctrlr := Controller{Ctx: ctx}
	err := ctrlr.Bind(&s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}

func TestBindNoContentType(t *testing.T) {
	var s struct {
		Foo string `json:"foo"`
	}
	header := map[string][]string{}
	request := &http.Request{Header: header}
	input := &context.radiantInput{RequestBody: []byte(`{"foo": "FOO"}`)}
	ctx := &context.Context{Request: request, Input: input}
	ctrlr := Controller{Ctx: ctx}
	err := ctrlr.Bind(&s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}

func TestBindXML(t *testing.T) {
	var s struct {
		Foo string `xml:"foo"`
	}
	xmlBody := `<?xml version="1.0" encoding="UTF-8"?>
<root>
   <foo>FOO</foo>
</root>`
	header := map[string][]string{"Content-Type": {"text/xml"}}
	request := &http.Request{Header: header}
	input := &context.radiantInput{RequestBody: []byte(xmlBody)}
	ctx := &context.Context{Request: request, Input: input}
	ctrlr := Controller{Ctx: ctx}
	err := ctrlr.Bind(&s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}

func TestBindYAML(t *testing.T) {
	var s struct {
		Foo string `yaml:"foo"`
	}
	header := map[string][]string{"Content-Type": {"application/x-yaml"}}
	request := &http.Request{Header: header}
	input := &context.radiantInput{RequestBody: []byte("foo: FOO")}
	ctx := &context.Context{Request: request, Input: input}
	ctrlr := Controller{Ctx: ctx}
	err := ctrlr.Bind(&s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}

type TestRespController struct {
	Controller
}

func (t *TestRespController) TestResponse() {
	type S struct {
		Foo string `json:"foo" xml:"foo" yaml:"foo"`
	}

	bar := S{Foo: "bar"}

	_ = t.Resp(bar)
}

type respTestCase struct {
	Accept                string
	ExpectedContentLength int64
	ExpectedResponse      string
}

func TestControllerResp(t *testing.T) {
	// test cases
	tcs := []respTestCase{
		{Accept: context.ApplicationJSON, ExpectedContentLength: 13, ExpectedResponse: `{"foo":"bar"}`},
		{Accept: context.ApplicationXML, ExpectedContentLength: 21, ExpectedResponse: `<S><foo>bar</foo></S>`},
		{Accept: context.ApplicationYAML, ExpectedContentLength: 9, ExpectedResponse: "foo: bar\n"},
		{Accept: "OTHER", ExpectedContentLength: 13, ExpectedResponse: `{"foo":"bar"}`},
	}

	for _, tc := range tcs {
		testControllerRespTestCases(t, tc)
	}
}

func testControllerRespTestCases(t *testing.T, tc respTestCase) {
	// create fake GET request
	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("Accept", tc.Accept)
	w := httptest.NewRecorder()

	// setup the handler
	handler := NewControllerRegister()
	handler.Add("/", &TestRespController{}, WithRouterMethods(&TestRespController{}, "get:TestResponse"))
	handler.ServeHTTP(w, r)

	response := w.Result()
	if response.ContentLength != tc.ExpectedContentLength {
		t.Errorf("TestResponse() unable to validate content length %d for %s", response.ContentLength, tc.Accept)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("TestResponse() failed to validate response code for %s", tc.Accept)
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("TestResponse() failed to parse response body for %s", tc.Accept)
	}
	bodyString := string(bodyBytes)
	if bodyString != tc.ExpectedResponse {
		t.Errorf("TestResponse() failed to validate response body '%s' for %s", bodyString, tc.Accept)
	}
}
