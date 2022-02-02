// Package memcache for cache provider
//
// depend on github.com/bradfitz/gomemcache/memcache
//
// go install github.com/bradfitz/gomemcache/memcache
//
// Usage:
// import(
//   _ "github.com/W3-Engineers-Ltd/Radiant/client/cache/memcache"
//   "github.com/W3-Engineers-Ltd/Radiant/client/cache"
// )
//
//  bm, err := cache.NewCache("memcache", `{"conn":"127.0.0.1:11211"}`)
//
//  more docs http://radiant.me/docs/module/cache.md
package memcache

import (
	"github.com/W3-Engineers-Ltd/Radiant/adapter/cache"
	"github.com/W3-Engineers-Ltd/Radiant/client/cache/memcache"
)

// NewMemCache create new memcache adapter.
func NewMemCache() cache.Cache {
	return cache.CreateNewToOldCacheAdapter(memcache.NewMemCache())
}

func init() {
	cache.Register("memcache", NewMemCache)
}
