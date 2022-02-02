package ssdb

import (
	"github.com/W3-Engineers-Ltd/Radiant/adapter/cache"
	ssdb2 "github.com/W3-Engineers-Ltd/Radiant/client/cache/ssdb"
)

// NewSsdbCache create new ssdb adapter.
func NewSsdbCache() cache.Cache {
	return cache.CreateNewToOldCacheAdapter(ssdb2.NewSsdbCache())
}

func init() {
	cache.Register("ssdb", NewSsdbCache)
}
