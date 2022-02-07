// Copyright 2014 beego Author. All Rights Reserved.
//

package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

var FilterUser = func(ctx *context.Context) {
	ctx.Output.Body([]byte("i am " + ctx.Input.Param(":last") + ctx.Input.Param(":first")))
}

func TestFilter(t *testing.T) {
	r, _ := http.NewRequest("GET", "/person/asta/Xie", nil)
	w := httptest.NewRecorder()
	handler := NewControllerRegister()
	handler.InsertFilter("/person/:last/:first", BeforeRouter, FilterUser)
	handler.Add("/person/:last/:first", &TestController{})
	handler.ServeHTTP(w, r)
	if w.Body.String() != "i am astaXie" {
		t.Errorf("user define func can't run")
	}
}

var FilterAdminUser = func(ctx *context.Context) {
	ctx.Output.Body([]byte("i am admin"))
}

// Filter pattern /admin/:all
// all url like    /admin/    /admin/xie    will all get filter

func TestPatternTwo(t *testing.T) {
	r, _ := http.NewRequest("GET", "/admin/", nil)
	w := httptest.NewRecorder()
	handler := NewControllerRegister()
	handler.InsertFilter("/admin/?:all", BeforeRouter, FilterAdminUser)
	handler.ServeHTTP(w, r)
	if w.Body.String() != "i am admin" {
		t.Errorf("filter /admin/ can't run")
	}
}

func TestPatternThree(t *testing.T) {
	r, _ := http.NewRequest("GET", "/admin/astaxie", nil)
	w := httptest.NewRecorder()
	handler := NewControllerRegister()
	handler.InsertFilter("/admin/:all", BeforeRouter, FilterAdminUser)
	handler.ServeHTTP(w, r)
	if w.Body.String() != "i am admin" {
		t.Errorf("filter /admin/astaxie can't run")
	}
}
