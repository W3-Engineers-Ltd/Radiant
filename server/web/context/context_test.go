// Copyright 2021 radiant Author. All Rights Reserved.
//

package context

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/W3-Engineers-Ltd/Radiant/server/web/session"
)

func TestXsrfReset_01(t *testing.T) {
	r := &http.Request{}
	c := NewContext()
	c.Request = r
	c.ResponseWriter = &Response{}
	c.ResponseWriter.reset(httptest.NewRecorder())
	c.Output.Reset(c)
	c.Input.Reset(c)
	c.XSRFToken("key", 16)
	if c._xsrfToken == "" {
		t.FailNow()
	}
	token := c._xsrfToken
	c.Reset(&Response{ResponseWriter: httptest.NewRecorder()}, r)
	if c._xsrfToken != "" {
		t.FailNow()
	}
	c.XSRFToken("key", 16)
	if c._xsrfToken == "" {
		t.FailNow()
	}
	if token == c._xsrfToken {
		t.FailNow()
	}
}

func TestContext_Session(t *testing.T) {
	c := NewContext()
	if store, err := c.Session(); store != nil || err == nil {
		t.FailNow()
	}
}

func TestContext_Session1(t *testing.T) {
	c := Context{}
	if store, err := c.Session(); store != nil || err == nil {
		t.FailNow()
	}
}

func TestContext_Session2(t *testing.T) {
	c := NewContext()
	c.Input.CruSession = &session.MemSessionStore{}

	if store, err := c.Session(); store == nil || err != nil {
		t.FailNow()
	}
}

func TestSetCookie(t *testing.T) {
	type cookie struct {
		Name     string
		Value    string
		MaxAge   int64
		Path     string
		Domain   string
		Secure   bool
		HttpOnly bool
		SameSite string
	}
	type testItem struct {
		item cookie
		want string
	}
	cases := []struct {
		request string
		valueGp []testItem
	}{
		{"/", []testItem{{cookie{"name", "value", -1, "/", "", false, false, "Strict"}, "name=value; Max-Age=0; Path=/; SameSite=Strict"}}},
		{"/", []testItem{{cookie{"name", "value", -1, "/", "", false, false, "Lax"}, "name=value; Max-Age=0; Path=/; SameSite=Lax"}}},
		{"/", []testItem{{cookie{"name", "value", -1, "/", "", false, false, "None"}, "name=value; Max-Age=0; Path=/; SameSite=None"}}},
		{"/", []testItem{{cookie{"name", "value", -1, "/", "", false, false, ""}, "name=value; Max-Age=0; Path=/"}}},
	}
	for _, c := range cases {
		r, _ := http.NewRequest("GET", c.request, nil)
		output := NewOutput()
		output.Context = NewContext()
		output.Context.Reset(httptest.NewRecorder(), r)
		for _, item := range c.valueGp {
			params := item.item
			var others = []interface{}{params.MaxAge, params.Path, params.Domain, params.Secure, params.HttpOnly, params.SameSite}
			output.Context.SetCookie(params.Name, params.Value, others...)
			got := output.Context.ResponseWriter.Header().Get("Set-Cookie")
			if got != item.want {
				t.Fatalf("SetCookie error,should be:\n%v \ngot:\n%v", item.want, got)
			}
		}
	}
}
