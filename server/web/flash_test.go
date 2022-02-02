package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestFlashController struct {
	Controller
}

func (t *TestFlashController) TestWriteFlash() {
	flash := NewFlash()
	flash.Notice("TestFlashString")
	flash.Store(&t.Controller)
	// we choose to serve json because we don't want to load a template html file
	t.ServeJSON(true)
}

func TestFlashHeader(t *testing.T) {
	// create fake GET request
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// setup the handler
	handler := NewControllerRegister()
	handler.Add("/", &TestFlashController{}, WithRouterMethods(&TestFlashController{}, "get:TestWriteFlash"))
	handler.ServeHTTP(w, r)

	// get the Set-Cookie value
	sc := w.Header().Get("Set-Cookie")
	// match for the expected header
	res := strings.Contains(sc, "radiant_FLASH=%00notice%23radiantFLASH%23TestFlashString%00")
	// validate the assertion
	if !res {
		t.Errorf("TestFlashHeader() unable to validate flash message")
	}
}
