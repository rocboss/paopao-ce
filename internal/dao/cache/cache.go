// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

func NewRedisCache() core.RedisCache {
	return &redisCache{
		c: conf.MustRedisClient(),
	}
}

func NewBigCacheIndexService(ips core.IndexPostsService, ams core.AuthorizationManageService) (core.CacheIndexService, core.VersionInfo) {
	s := conf.BigCacheIndexSetting
	c := bigcache.DefaultConfig(s.ExpireInSecond)
	c.Shards = s.MaxIndexPage
	c.HardMaxCacheSize = s.HardMaxCacheSize
	c.Verbose = s.Verbose
	c.MaxEntrySize = 10000
	c.Logger = logrus.StandardLogger()

	bc, err := bigcache.NewBigCache(c)
	if err != nil {
		logrus.Fatalf("initial bigCahceIndex failure by err: %v", err)
	}
	cacheIndex := newCacheIndexSrv(ips, ams, &bigCacheTweetsCache{
		bc: bc,
	})
	return cacheIndex, cacheIndex
}

func NewRedisCacheIndexService(ips core.IndexPostsService, ams core.AuthorizationManageService) (core.CacheIndexService, core.VersionInfo) {
	cacheIndex := newCacheIndexSrv(ips, ams, &redisCacheTweetsCache{
		expireDuration: conf.RedisCacheIndexSetting.ExpireInSecond,
		expireInSecond: int64(conf.RedisCacheIndexSetting.ExpireInSecond / time.Second),
		c:              conf.MustRedisClient(),
	})
	return cacheIndex, cacheIndex
}

func NewSimpleCacheIndexService(indexPosts core.IndexPostsService) (core.CacheIndexService, core.VersionInfo) {
	s := conf.SimpleCacheIndexSetting
	cacheIndex := &simpleCacheIndexServant{
		ips:             indexPosts,
		maxIndexSize:    s.MaxIndexSize,
		indexPosts:      nil,
		checkTick:       time.NewTicker(s.CheckTickDuration), // check whether need update index every 1 minute
		expireIndexTick: time.NewTicker(time.Second),
	}

	// force expire index every ExpireTickDuration second
	if s.ExpireTickDuration != 0 {
		cacheIndex.expireIndexTick.Reset(s.CheckTickDuration)
	} else {
		cacheIndex.expireIndexTick.Stop()
	}

	// indexActionCh capacity custom configure by conf.yaml need in [10, 10000]
	// or re-compile source to adjust min/max capacity
	capacity := conf.CacheIndexSetting.MaxUpdateQPS
	if capacity < 10 {
		capacity = 10
	} else if capacity > 10000 {
		capacity = 10000
	}
	cacheIndex.indexActionCh = make(chan core.IdxAct, capacity)

	// start index posts
	cacheIndex.atomicIndex.Store(cacheIndex.indexPosts)
	go cacheIndex.startIndexPosts()

	return cacheIndex, cacheIndex
}

func NewNoneCacheIndexService(indexPosts core.IndexPostsService) (core.CacheIndexService, core.VersionInfo) {
	obj := &noneCacheIndexServant{
		ips: indexPosts,
	}
	return obj, obj
}
