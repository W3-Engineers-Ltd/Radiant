// Package couchbase for session provider
//
// depend on github.com/couchbaselabs/go-couchbasee
//
// go install github.com/couchbaselabs/go-couchbase
//
// Usage:
// import(
//   _ "github.com/W3-Engineers-Ltd/Radiant/server/web/session/couchbase"
//   "github.com/W3-Engineers-Ltd/Radiant/server/web/session"
// )
//
//	func init() {
//		globalSessions, _ = session.NewManager("couchbase", ``{"cookieName":"gosessionid","gclifetime":3600,"ProviderConfig":"http://host:port/, Pool, Bucket"}``)
//		go globalSessions.GC()
//	}
//
// more docs: http://radiant.me/docs/module/session.md
package couchbase

import (
	"context"
	"net/http"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/session"
	radicalcb "github.com/W3-Engineers-Ltd/Radiant/server/web/session/couchbase"
)

// SessionStore store each session
type SessionStore radicalcb.SessionStore

// Provider couchabse provided
type Provider radicalcb.Provider

// Set value to couchabse session
func (cs *SessionStore) Set(key, value interface{}) error {
	return (*radicalcb.SessionStore)(cs).Set(context.Background(), key, value)
}

// Get value from couchabse session
func (cs *SessionStore) Get(key interface{}) interface{} {
	return (*radicalcb.SessionStore)(cs).Get(context.Background(), key)
}

// Delete value in couchbase session by given key
func (cs *SessionStore) Delete(key interface{}) error {
	return (*radicalcb.SessionStore)(cs).Delete(context.Background(), key)
}

// Flush Clean all values in couchbase session
func (cs *SessionStore) Flush() error {
	return (*radicalcb.SessionStore)(cs).Flush(context.Background())
}

// SessionID Get couchbase session store id
func (cs *SessionStore) SessionID() string {
	return (*radicalcb.SessionStore)(cs).SessionID(context.Background())
}

// SessionRelease Write couchbase session with Gob string
func (cs *SessionStore) SessionRelease(w http.ResponseWriter) {
	(*radicalcb.SessionStore)(cs).SessionRelease(context.Background(), w)
}

// SessionInit init couchbase session
// savepath like couchbase server REST/JSON URL
// e.g. http://host:port/, Pool, Bucket
func (cp *Provider) SessionInit(maxlifetime int64, savePath string) error {
	return (*radicalcb.Provider)(cp).SessionInit(context.Background(), maxlifetime, savePath)
}

// SessionRead read couchbase session by sid
func (cp *Provider) SessionRead(sid string) (session.Store, error) {
	s, err := (*radicalcb.Provider)(cp).SessionRead(context.Background(), sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionExist Check couchbase session exist.
// it checkes sid exist or not.
func (cp *Provider) SessionExist(sid string) bool {
	res, _ := (*radicalcb.Provider)(cp).SessionExist(context.Background(), sid)
	return res
}

// SessionRegenerate remove oldsid and use sid to generate new session
func (cp *Provider) SessionRegenerate(oldsid, sid string) (session.Store, error) {
	s, err := (*radicalcb.Provider)(cp).SessionRegenerate(context.Background(), oldsid, sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionDestroy Remove bucket in this couchbase
func (cp *Provider) SessionDestroy(sid string) error {
	return (*radicalcb.Provider)(cp).SessionDestroy(context.Background(), sid)
}

// SessionGC Recycle
func (cp *Provider) SessionGC() {
	(*radicalcb.Provider)(cp).SessionGC(context.Background())
}

// SessionAll return all active session
func (cp *Provider) SessionAll() int {
	return (*radicalcb.Provider)(cp).SessionAll(context.Background())
}
