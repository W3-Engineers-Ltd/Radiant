// Package redis for cache provider
//
// depend on github.com/gomodule/redigo/redis
//
// go install github.com/gomodule/redigo/redis
//
// Usage:
// import(
//   _ "github.com/W3-Engineers-Ltd/Radiant/client/cache/redis"
//   "github.com/W3-Engineers-Ltd/Radiant/client/cache"
// )
//
//  bm, err := cache.NewCache("redis", `{"conn":"127.0.0.1:11211"}`)
//
//  more docs http://radiant.me/docs/module/cache.md
package redis

import (
	"github.com/W3-Engineers-Ltd/Radiant/adapter/cache"
	redis2 "github.com/W3-Engineers-Ltd/Radiant/client/cache/redis"
)

// DefaultKey the collection name of redis for cache adapter.
var DefaultKey = "beecacheRedis"

// NewRedisCache create new redis cache with default collection name.
func NewRedisCache() cache.Cache {
	return cache.CreateNewToOldCacheAdapter(redis2.NewRedisCache())
}

func init() {
	cache.Register("redis", NewRedisCache)
}
