// Package memcache for session provider
//
// depend on github.com/bradfitz/gomemcache/memcache
//
// go install github.com/bradfitz/gomemcache/memcache
//
// Usage:
// import(
//   _ "github.com/W3-Engineers-Ltd/Radiant/server/web/session/memcache"
//   "github.com/W3-Engineers-Ltd/Radiant/server/web/session"
// )
//
//	func init() {
//		globalSessions, _ = session.NewManager("memcache", ``{"cookieName":"gosessionid","gclifetime":3600,"ProviderConfig":"127.0.0.1:11211"}``)
//		go globalSessions.GC()
//	}
//
// more docs: http://radiant.me/docs/module/session.md
package memcache

import (
	"context"
	"net/http"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/session"
	radicalmem "github.com/W3-Engineers-Ltd/Radiant/server/web/session/memcache"
)

// SessionStore memcache session store
type SessionStore radicalmem.SessionStore

// Set value in memcache session
func (rs *SessionStore) Set(key, value interface{}) error {
	return (*radicalmem.SessionStore)(rs).Set(context.Background(), key, value)
}

// Get value in memcache session
func (rs *SessionStore) Get(key interface{}) interface{} {
	return (*radicalmem.SessionStore)(rs).Get(context.Background(), key)
}

// Delete value in memcache session
func (rs *SessionStore) Delete(key interface{}) error {
	return (*radicalmem.SessionStore)(rs).Delete(context.Background(), key)
}

// Flush clear all values in memcache session
func (rs *SessionStore) Flush() error {
	return (*radicalmem.SessionStore)(rs).Flush(context.Background())
}

// SessionID get memcache session id
func (rs *SessionStore) SessionID() string {
	return (*radicalmem.SessionStore)(rs).SessionID(context.Background())
}

// SessionRelease save session values to memcache
func (rs *SessionStore) SessionRelease(w http.ResponseWriter) {
	(*radicalmem.SessionStore)(rs).SessionRelease(context.Background(), w)
}

// MemProvider memcache session provider
type MemProvider radicalmem.MemProvider

// SessionInit init memcache session
// savepath like
// e.g. 127.0.0.1:9090
func (rp *MemProvider) SessionInit(maxlifetime int64, savePath string) error {
	return (*radicalmem.MemProvider)(rp).SessionInit(context.Background(), maxlifetime, savePath)
}

// SessionRead read memcache session by sid
func (rp *MemProvider) SessionRead(sid string) (session.Store, error) {
	s, err := (*radicalmem.MemProvider)(rp).SessionRead(context.Background(), sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionExist check memcache session exist by sid
func (rp *MemProvider) SessionExist(sid string) bool {
	res, _ := (*radicalmem.MemProvider)(rp).SessionExist(context.Background(), sid)
	return res
}

// SessionRegenerate generate new sid for memcache session
func (rp *MemProvider) SessionRegenerate(oldsid, sid string) (session.Store, error) {
	s, err := (*radicalmem.MemProvider)(rp).SessionRegenerate(context.Background(), oldsid, sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionDestroy delete memcache session by id
func (rp *MemProvider) SessionDestroy(sid string) error {
	return (*radicalmem.MemProvider)(rp).SessionDestroy(context.Background(), sid)
}

// SessionGC Impelment method, no used.
func (rp *MemProvider) SessionGC() {
	(*radicalmem.MemProvider)(rp).SessionGC(context.Background())
}

// SessionAll return all activeSession
func (rp *MemProvider) SessionAll() int {
	return (*radicalmem.MemProvider)(rp).SessionAll(context.Background())
}
