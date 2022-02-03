// Package ledis provide session Provider
package ledis

import (
	"context"
	"net/http"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/session"
	radicalLedis "github.com/W3-Engineers-Ltd/Radiant/server/web/session/ledis"
)

// SessionStore ledis session store
type SessionStore radicalLedis.SessionStore

// Set value in ledis session
func (ls *SessionStore) Set(key, value interface{}) error {
	return (*radicalLedis.SessionStore)(ls).Set(context.Background(), key, value)
}

// Get value in ledis session
func (ls *SessionStore) Get(key interface{}) interface{} {
	return (*radicalLedis.SessionStore)(ls).Get(context.Background(), key)
}

// Delete value in ledis session
func (ls *SessionStore) Delete(key interface{}) error {
	return (*radicalLedis.SessionStore)(ls).Delete(context.Background(), key)
}

// Flush clear all values in ledis session
func (ls *SessionStore) Flush() error {
	return (*radicalLedis.SessionStore)(ls).Flush(context.Background())
}

// SessionID get ledis session id
func (ls *SessionStore) SessionID() string {
	return (*radicalLedis.SessionStore)(ls).SessionID(context.Background())
}

// SessionRelease save session values to ledis
func (ls *SessionStore) SessionRelease(w http.ResponseWriter) {
	(*radicalLedis.SessionStore)(ls).SessionRelease(context.Background(), w)
}

// Provider ledis session provider
type Provider radicalLedis.Provider

// SessionInit init ledis session
// savepath like ledis server saveDataPath,pool size
// e.g. 127.0.0.1:6379,100,astaxie
func (lp *Provider) SessionInit(maxlifetime int64, savePath string) error {
	return (*radicalLedis.Provider)(lp).SessionInit(context.Background(), maxlifetime, savePath)
}

// SessionRead read ledis session by sid
func (lp *Provider) SessionRead(sid string) (session.Store, error) {
	s, err := (*radicalLedis.Provider)(lp).SessionRead(context.Background(), sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionExist check ledis session exist by sid
func (lp *Provider) SessionExist(sid string) bool {
	res, _ := (*radicalLedis.Provider)(lp).SessionExist(context.Background(), sid)
	return res
}

// SessionRegenerate generate new sid for ledis session
func (lp *Provider) SessionRegenerate(oldsid, sid string) (session.Store, error) {
	s, err := (*radicalLedis.Provider)(lp).SessionRegenerate(context.Background(), oldsid, sid)
	return session.CreateNewToOldStoreAdapter(s), err
}

// SessionDestroy delete ledis session by id
func (lp *Provider) SessionDestroy(sid string) error {
	return (*radicalLedis.Provider)(lp).SessionDestroy(context.Background(), sid)
}

// SessionGC Impelment method, no used.
func (lp *Provider) SessionGC() {
	(*radicalLedis.Provider)(lp).SessionGC(context.Background())
}

// SessionAll return all active session
func (lp *Provider) SessionAll() int {
	return (*radicalLedis.Provider)(lp).SessionAll(context.Background())
}
