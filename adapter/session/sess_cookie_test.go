package session

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const setCookieKey = "Set-Cookie"

func TestCookie(t *testing.T) {
	config := `{"cookieName":"gosessionid","enableSetCookie":false,"gclifetime":3600,"ProviderConfig":"{\"cookieName\":\"gosessionid\",\"securityKey\":\"radiantcookiehashkey\"}"}`
	conf := new(ManagerConfig)
	if err := json.Unmarshal([]byte(config), conf); err != nil {
		t.Fatal("json decode error", err)
	}
	globalSessions, err := NewManager("cookie", conf)
	if err != nil {
		t.Fatal("init cookie session err", err)
	}
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	sess, err := globalSessions.SessionStart(w, r)
	if err != nil {
		t.Fatal("set error,", err)
	}
	err = sess.Set("username", "astaxie")
	if err != nil {
		t.Fatal("set error,", err)
	}
	if username := sess.Get("username"); username != "astaxie" {
		t.Fatal("get username error")
	}
	sess.SessionRelease(w)

	if cookiestr := w.Header().Get(setCookieKey); cookiestr == "" {
		t.Fatal("setcookie error")
	} else {
		parts := strings.Split(strings.TrimSpace(cookiestr), ";")
		for k, v := range parts {
			nameval := strings.Split(v, "=")
			if k == 0 && nameval[0] != "gosessionid" {
				t.Fatal("error")
			}
		}
	}
}

func TestDestorySessionCookie(t *testing.T) {
	config := `{"cookieName":"gosessionid","enableSetCookie":true,"gclifetime":3600,"ProviderConfig":"{\"cookieName\":\"gosessionid\",\"securityKey\":\"radiantcookiehashkey\"}"}`
	conf := new(ManagerConfig)
	if err := json.Unmarshal([]byte(config), conf); err != nil {
		t.Fatal("json decode error", err)
	}
	globalSessions, err := NewManager("cookie", conf)
	if err != nil {
		t.Fatal("init cookie session err", err)
	}

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	session, err := globalSessions.SessionStart(w, r)
	if err != nil {
		t.Fatal("session start err,", err)
	}

	// request again ,will get same sesssion id .
	r1, _ := http.NewRequest("GET", "/", nil)
	r1.Header.Set("Cookie", w.Header().Get(setCookieKey))
	w = httptest.NewRecorder()
	newSession, err := globalSessions.SessionStart(w, r1)
	if err != nil {
		t.Fatal("session start err,", err)
	}
	if newSession.SessionID() != session.SessionID() {
		t.Fatal("get cookie session id is not the same again.")
	}

	// After destroy session , will get a new session id .
	globalSessions.SessionDestroy(w, r1)
	r2, _ := http.NewRequest("GET", "/", nil)
	r2.Header.Set("Cookie", w.Header().Get(setCookieKey))

	w = httptest.NewRecorder()
	newSession, err = globalSessions.SessionStart(w, r2)
	if err != nil {
		t.Fatal("session start error")
	}
	if newSession.SessionID() == session.SessionID() {
		t.Fatal("after destroy session and reqeust again ,get cookie session id is same.")
	}
}
