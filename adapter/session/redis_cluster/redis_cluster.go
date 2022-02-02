// Package redis for session provider
//
// depend on github.com/go-redis/redis
//
// go install github.com/go-redis/redis
//
// Usage:
// import(
//   _ "github.com/W3-Engineers-Ltd/Radiant/server/web/session/redis_cluster"
//   "github.com/W3-Engineers-Ltd/Radiant/server/web/session"
// )
//
//	func init() {
//		globalSessions, _ = session.NewManager("redis_cluster", ``{"cookieName":"gosessionid","gclifetime":3600,"ProviderConfig":"127.0.0.1:7070;127.0.0.1:7071"}``)
//		go globalSessions.GC()
//	}
//
// more docs: http://radiant.me/docs/module/session.md
package redis_cluster

import (
	"context"
	"net/http"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/session"
	cluster "github.com/W3-Engineers-Ltd/Radiant/server/web/session/redis_cluster"
)

// MaxPoolSize redis_cluster max pool size
var MaxPoolSize = cluster.MaxPoolSize

// SessionStore redis_cluster session store
type SessionStore cluster.SessionStore

// Set value in redis_cluster session
func (rs *SessionStore) Set(key, value interface{}) error {
	return (*cluster.SessionStore)(rs).Set(context.Background(), key, value)
}

// Get value in redis_cluster session
func (rs *SessionStore) Get(key interface{}) interface{} {
	return (*cluster.SessionStore)(rs).Get(context.Background(), key)
}

// Delete value in redis_cluster session
func (rs *SessionStore) Delete(key interface{}) error {
	return (*cluster.SessionStore)(rs).Delete(context.Background(), key)
}

// Flush clear all values in redis_cluster session
func (rs *SessionStore) Flush() error {
	return (*cluster.SessionStore)(rs).Flush(context.Background())
}

// SessionID get redis_cluster session id
func (rs *SessionStore) SessionID() string {
	return (*cluster.SessionStore)(rs).SessionID(context.Background())
}

// SessionRelease save session values to redis_cluster
func (rs *SessionStore) SessionRelease(w http.ResponseWriter) {
	(*cluster.SessionStore)(rs).SessionRelease(context.Background(), w)
}

// Provider redis_cluster session provider
type Provider cluster.Provider

// SessionInit init redis_cluster session
// savepath like redis server addr,pool size,password,dbnum
// e.g. 127.0.0.1:6379;127.0.0.1:6380,100,test,0
func (rp *Provider) SessionInit(maxlifetime int64, savePath string) error {
	return (*cluster.Provider)(rp).SessionInit(context.Background(), maxlifetime, savePath)
}

// SessionRead read redis_cluster session by sid
func (rp *Provider) SessionRead(sid string) (session.Store, error) {
	s, err := (*cluster.Provider)(rp).SessionRead(context.Background(), sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionExist check redis_cluster session exist by sid
func (rp *Provider) SessionExist(sid string) bool {
	res, _ := (*cluster.Provider)(rp).SessionExist(context.Background(), sid)
	return res
}

// SessionRegenerate generate new sid for redis_cluster session
func (rp *Provider) SessionRegenerate(oldsid, sid string) (session.Store, error) {
	s, err := (*cluster.Provider)(rp).SessionRegenerate(context.Background(), oldsid, sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionDestroy delete redis session by id
func (rp *Provider) SessionDestroy(sid string) error {
	return (*cluster.Provider)(rp).SessionDestroy(context.Background(), sid)
}

// SessionGC Impelment method, no used.
func (rp *Provider) SessionGC() {
	(*cluster.Provider)(rp).SessionGC(context.Background())
}

// SessionAll return all activeSession
func (rp *Provider) SessionAll() int {
	return (*cluster.Provider)(rp).SessionAll(context.Background())
}
