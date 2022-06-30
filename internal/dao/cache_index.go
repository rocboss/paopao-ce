package dao

import (
	"sync/atomic"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
)

var (
	_ core.CacheIndexService = (*simpleCacheIndexServant)(nil)
	_ core.CacheIndexService = (*bigCacheIndexServant)(nil)
	_ core.CacheIndexService = (*noneCacheIndexServant)(nil)
)

type postsEntry struct {
	key   string
	posts []*model.PostFormated
}

type indexPostsFunc func(*model.User, int, int) ([]*model.PostFormated, error)

type bigCacheIndexServant struct {
	getIndexPosts      indexPostsFunc
	indexActionCh      chan core.IndexActionT
	cachePostsCh       chan *postsEntry
	cache              *bigcache.BigCache
	lastCacheResetTime time.Time
}

type simpleCacheIndexServant struct {
	getIndexPosts   indexPostsFunc
	indexActionCh   chan core.IndexActionT
	indexPosts      []*model.PostFormated
	atomicIndex     atomic.Value
	maxIndexSize    int
	checkTick       *time.Ticker
	expireIndexTick *time.Ticker
}

type noneCacheIndexServant struct {
	getIndexPosts indexPostsFunc
}
