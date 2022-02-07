// Copyright 2014 beego Author. All Rights Reserved.
//

package cache

import (
	"github.com/W3-Engineers-Ltd/Radiant/client/cache"
)

// NewFileCache Create new file cache with no config.
// the level and expiry need set in method StartAndGC as config string.
func NewFileCache() Cache {
	//    return &FileCache{CachePath:FileCachePath, FileSuffix:FileCacheFileSuffix}
	return CreateNewToOldCacheAdapter(cache.NewFileCache())
}

func init() {
	Register("file", NewFileCache)
}
