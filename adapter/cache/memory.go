package cache

import (
	"github.com/W3-Engineers-Ltd/Radiant/client/cache"
)

// NewMemoryCache returns a new MemoryCache.
func NewMemoryCache() Cache {
	return CreateNewToOldCacheAdapter(cache.NewMemoryCache())
}

func init() {
	Register("memory", NewMemoryCache)
}
