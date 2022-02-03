// Package redis for session provider
//
// depend on github.com/gomodule/redigo/redis
//
// go install github.com/gomodule/redigo/redis
//
// Usage:
// import(
//   _ "github.com/W3-Engineers-Ltd/Radiant/server/web/session/redis"
//   "github.com/W3-Engineers-Ltd/Radiant/server/web/session"
// )
//
// 	func init() {
// 		globalSessions, _ = session.NewManager("redis", ``{"cookieName":"gosessionid","gclifetime":3600,"ProviderConfig":"127.0.0.1:7070"}``)
// 		go globalSessions.GC()
// 	}
//
// more docs: http://radiant.me/docs/module/session.md
package redis

import (
	"context"
	"net/http"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/session"
	radicalRedis "github.com/W3-Engineers-Ltd/Radiant/server/web/session/redis"
)

// MaxPoolSize redis max pool size
var MaxPoolSize = radicalRedis.MaxPoolSize

// SessionStore redis session store
type SessionStore radicalRedis.SessionStore

// Set value in redis session
func (rs *SessionStore) Set(key, value interface{}) error {
	return (*radicalRedis.SessionStore)(rs).Set(context.Background(), key, value)
}

// Get value in redis session
func (rs *SessionStore) Get(key interface{}) interface{} {
	return (*radicalRedis.SessionStore)(rs).Get(context.Background(), key)
}

// Delete value in redis session
func (rs *SessionStore) Delete(key interface{}) error {
	return (*radicalRedis.SessionStore)(rs).Delete(context.Background(), key)
}

// Flush clear all values in redis session
func (rs *SessionStore) Flush() error {
	return (*radicalRedis.SessionStore)(rs).Flush(context.Background())
}

// SessionID get redis session id
func (rs *SessionStore) SessionID() string {
	return (*radicalRedis.SessionStore)(rs).SessionID(context.Background())
}

// SessionRelease save session values to redis
func (rs *SessionStore) SessionRelease(w http.ResponseWriter) {
	(*radicalRedis.SessionStore)(rs).SessionRelease(context.Background(), w)
}

// Provider redis session provider
type Provider radicalRedis.Provider

// SessionInit init redis session
// savepath like redis server addr,pool size,password,dbnum,IdleTimeout second
// e.g. 127.0.0.1:6379,100,astaxie,0,30
func (rp *Provider) SessionInit(maxlifetime int64, savePath string) error {
	return (*radicalRedis.Provider)(rp).SessionInit(context.Background(), maxlifetime, savePath)
}

// SessionRead read redis session by sid
func (rp *Provider) SessionRead(sid string) (session.Store, error) {
	s, err := (*radicalRedis.Provider)(rp).SessionRead(context.Background(), sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionExist check redis session exist by sid
func (rp *Provider) SessionExist(sid string) bool {
	res, _ := (*radicalRedis.Provider)(rp).SessionExist(context.Background(), sid)
	return res
}

// SessionRegenerate generate new sid for redis session
func (rp *Provider) SessionRegenerate(oldsid, sid string) (session.Store, error) {
	s, err := (*radicalRedis.Provider)(rp).SessionRegenerate(context.Background(), oldsid, sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionDestroy delete redis session by id
func (rp *Provider) SessionDestroy(sid string) error {
	return (*radicalRedis.Provider)(rp).SessionDestroy(context.Background(), sid)
}

// SessionGC Impelment method, no used.
func (rp *Provider) SessionGC() {
	(*radicalRedis.Provider)(rp).SessionGC(context.Background())
}

// SessionAll return all activeSession
func (rp *Provider) SessionAll() int {
	return (*radicalRedis.Provider)(rp).SessionAll(context.Background())
}
