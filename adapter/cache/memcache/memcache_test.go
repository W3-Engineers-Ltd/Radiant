package memcache

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/cache"
)

func TestMemcacheCache(t *testing.T) {
	addr := os.Getenv("MEMCACHE_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11211"
	}

	bm, err := cache.NewCache("memcache", fmt.Sprintf(`{"conn": "%s"}`, addr))
	assert.Nil(t, err)
	timeoutDuration := 5 * time.Second

	assert.Nil(t, bm.Put("astaxie", "1", timeoutDuration))

	assert.True(t, bm.IsExist("astaxie"))

	time.Sleep(11 * time.Second)

	assert.False(t, bm.IsExist("astaxie"))

	assert.Nil(t, bm.Put("astaxie", "1", timeoutDuration))
	v, err := strconv.Atoi(string(bm.Get("astaxie").([]byte)))
	assert.Nil(t, err)
	assert.Equal(t, 1, v)

	assert.Nil(t, bm.Incr("astaxie"))

	v, err = strconv.Atoi(string(bm.Get("astaxie").([]byte)))
	assert.Nil(t, err)
	assert.Equal(t, 2, v)

	assert.Nil(t, bm.Decr("astaxie"))

	v, err = strconv.Atoi(string(bm.Get("astaxie").([]byte)))
	assert.Nil(t, err)
	assert.Equal(t, 1, v)

	assert.Nil(t, bm.Delete("astaxie"))

	assert.False(t, bm.IsExist("astaxie"))

	assert.Nil(t, bm.Put("astaxie", "author", timeoutDuration))

	assert.True(t, bm.IsExist("astaxie"))

	assert.Equal(t, []byte("author"), bm.Get("astaxie"))

	assert.Nil(t, bm.Put("astaxie1", "author1", timeoutDuration))

	assert.True(t, bm.IsExist("astaxie1"))

	vv := bm.GetMulti([]string{"astaxie", "astaxie1"})
	assert.Equal(t, 2, len(vv))
	assert.Equal(t, []byte("author"), vv[0])
	assert.Equal(t, []byte("author1"), vv[1])

	assert.Nil(t, bm.ClearAll())
}
