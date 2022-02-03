// Copyright 2020
//

package session

import (
	"context"
	"net/http"

	"github.com/W3-Engineers-Ltd/Radiant/server/web/session"
)

type NewToOldStoreAdapter struct {
	delegate session.Store
}

func CreateNewToOldStoreAdapter(s session.Store) Store {
	return &NewToOldStoreAdapter{
		delegate: s,
	}
}

func (n *NewToOldStoreAdapter) Set(key, value interface{}) error {
	return n.delegate.Set(context.Background(), key, value)
}

func (n *NewToOldStoreAdapter) Get(key interface{}) interface{} {
	return n.delegate.Get(context.Background(), key)
}

func (n *NewToOldStoreAdapter) Delete(key interface{}) error {
	return n.delegate.Delete(context.Background(), key)
}

func (n *NewToOldStoreAdapter) SessionID() string {
	return n.delegate.SessionID(context.Background())
}

func (n *NewToOldStoreAdapter) SessionRelease(w http.ResponseWriter) {
	n.delegate.SessionRelease(context.Background(), w)
}

func (n *NewToOldStoreAdapter) Flush() error {
	return n.delegate.Flush(context.Background())
}

type oldToNewStoreAdapter struct {
	delegate Store
}

func (o *oldToNewStoreAdapter) Set(ctx context.Context, key, value interface{}) error {
	return o.delegate.Set(key, value)
}

func (o *oldToNewStoreAdapter) Get(ctx context.Context, key interface{}) interface{} {
	return o.delegate.Get(key)
}

func (o *oldToNewStoreAdapter) Delete(ctx context.Context, key interface{}) error {
	return o.delegate.Delete(key)
}

func (o *oldToNewStoreAdapter) SessionID(ctx context.Context) string {
	return o.delegate.SessionID()
}

func (o *oldToNewStoreAdapter) SessionRelease(ctx context.Context, w http.ResponseWriter) {
	o.delegate.SessionRelease(w)
}

func (o *oldToNewStoreAdapter) Flush(ctx context.Context) error {
	return o.delegate.Flush()
}
