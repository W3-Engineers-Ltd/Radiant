// Copyright 2020
//

package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHttpServerWithCfg(t *testing.T) {
	BConfig.AppName = "Before"
	svr := NewHttpServerWithCfg(BConfig)
	svr.Cfg.AppName = "hello"
	assert.Equal(t, "hello", BConfig.AppName)
}

func TestServerCtrlGet(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "/user", nil)
	w := httptest.NewRecorder()

	CtrlGet("/user", ExampleController.Ping)
	RadicalApp.Handlers.ServeHTTP(w, r)
	if w.Body.String() != exampleBody {
		t.Errorf("TestServerCtrlGet can't run")
	}
}

func TestServerCtrlPost(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPost, "/user", nil)
	w := httptest.NewRecorder()

	CtrlPost("/user", ExampleController.Ping)
	RadicalApp.Handlers.ServeHTTP(w, r)
	if w.Body.String() != exampleBody {
		t.Errorf("TestServerCtrlPost can't run")
	}
}

func TestServerCtrlHead(t *testing.T) {
	r, _ := http.NewRequest(http.MethodHead, "/user", nil)
	w := httptest.NewRecorder()

	CtrlHead("/user", ExampleController.Ping)
	RadicalApp.Handlers.ServeHTTP(w, r)
	if w.Body.String() != exampleBody {
		t.Errorf("TestServerCtrlHead can't run")
	}
}

func TestServerCtrlPut(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPut, "/user", nil)
	w := httptest.NewRecorder()

	CtrlPut("/user", ExampleController.Ping)
	RadicalApp.Handlers.ServeHTTP(w, r)
	if w.Body.String() != exampleBody {
		t.Errorf("TestServerCtrlPut can't run")
	}
}

func TestServerCtrlPatch(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPatch, "/user", nil)
	w := httptest.NewRecorder()

	CtrlPatch("/user", ExampleController.Ping)
	RadicalApp.Handlers.ServeHTTP(w, r)
	if w.Body.String() != exampleBody {
		t.Errorf("TestServerCtrlPatch can't run")
	}
}

func TestServerCtrlDelete(t *testing.T) {
	r, _ := http.NewRequest(http.MethodDelete, "/user", nil)
	w := httptest.NewRecorder()

	CtrlDelete("/user", ExampleController.Ping)
	RadicalApp.Handlers.ServeHTTP(w, r)
	if w.Body.String() != exampleBody {
		t.Errorf("TestServerCtrlDelete can't run")
	}
}

func TestServerCtrlAny(t *testing.T) {
	CtrlAny("/user", ExampleController.Ping)

	for method := range HTTPMETHOD {
		r, _ := http.NewRequest(method, "/user", nil)
		w := httptest.NewRecorder()
		RadicalApp.Handlers.ServeHTTP(w, r)
		if w.Body.String() != exampleBody {
			t.Errorf("TestServerCtrlAny can't run")
		}
	}
}
