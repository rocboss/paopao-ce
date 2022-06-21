package dao

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/allegro/bigcache/v3"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/sirupsen/logrus"
)

func newBigCacheIndexServant(getIndexPosts indexPostsFunc) *bigCacheIndexServant {
	s := conf.BigCacheIndexSetting

	config := bigcache.DefaultConfig(s.ExpireInSecond)
	config.Shards = s.MaxIndexPage
	config.Verbose = s.Verbose
	config.MaxEntrySize = 10000
	config.Logger = logrus.StandardLogger()
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		logrus.Fatalf("initial bigCahceIndex failure by err: %v", err)
	}

	cacheIndex := &bigCacheIndexServant{
		getIndexPosts: getIndexPosts,
		cache:         cache,
	}

	// indexActionCh capacity custom configure by conf.yaml need in [10, 10000]
	// or re-compile source to adjust min/max capacity
	capacity := conf.CacheIndexSetting.MaxUpdateQPS
	if capacity < 10 {
		capacity = 10
	} else if capacity > 10000 {
		capacity = 10000
	}
	cacheIndex.indexActionCh = make(chan core.IndexActionT, capacity)
	cacheIndex.cachePostsCh = make(chan *postsEntry, capacity)

	// 启动索引更新器
	go cacheIndex.startIndexPosts()

	return cacheIndex
}

func (s *bigCacheIndexServant) IndexPosts(userId int64, offset int, limit int) ([]*model.PostFormated, error) {
	key := s.keyFrom(userId, offset, limit)
	posts, err := s.getPosts(key)
	if err == nil {
		logrus.Debugf("get index posts from cache by key: %s userId: %d offset:%d limit:%d", key, userId, offset, limit)
		return posts, nil
	}

	if posts, err = s.getIndexPosts(userId, offset, limit); err != nil {
		return nil, err
	}
	logrus.Debugf("get index posts from database by userId: %d offset:%d limit:%d", userId, offset, limit)
	s.cachePosts(key, posts)
	return posts, nil
}

func (s *bigCacheIndexServant) getPosts(key string) ([]*model.PostFormated, error) {
	data, err := s.cache.Get(key)
	if err != nil {
		logrus.Debugf("get posts by key: %s from cache err: %v", key, err)
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var posts []*model.PostFormated
	if err := dec.Decode(&posts); err != nil {
		logrus.Debugf("get posts from cache in decode err: %v", err)
		return nil, err
	}
	return posts, nil
}

func (s *bigCacheIndexServant) cachePosts(key string, posts []*model.PostFormated) {
	entry := &postsEntry{key: key, posts: posts}
	select {
	case s.cachePostsCh <- entry:
		logrus.Debugf("cachePosts by chan of key: %s", key)
	default:
		go func(ch chan<- *postsEntry, entry *postsEntry) {
			logrus.Debugf("cachePosts indexAction by goroutine of key: %s", key)
			ch <- entry
		}(s.cachePostsCh, entry)
	}
}

func (s *bigCacheIndexServant) setPosts(entry *postsEntry) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(entry.posts); err != nil {
		logrus.Debugf("setPosts encode post entry err: %v", err)
		return
	}
	if err := s.cache.Set(entry.key, buf.Bytes()); err != nil {
		logrus.Debugf("setPosts set cache err: %v", err)
	}
	logrus.Debugf("setPosts set cache by key: %s", entry.key)
}

func (s *bigCacheIndexServant) keyFrom(userId int64, offset int, limit int) string {
	return fmt.Sprintf("index:%d:%d:%d", userId, offset, limit)
}

func (s *bigCacheIndexServant) SendAction(act core.IndexActionT) {
	select {
	case s.indexActionCh <- act:
		logrus.Debugf("send indexAction by chan: %s", act)
	default:
		go func(ch chan<- core.IndexActionT, act core.IndexActionT) {
			logrus.Debugf("send indexAction by goroutine: %s", act)
			ch <- act
		}(s.indexActionCh, act)
	}
}

func (s *bigCacheIndexServant) startIndexPosts() {
	for {
		select {
		case entry := <-s.cachePostsCh:
			s.setPosts(entry)
		case action := <-s.indexActionCh:
			switch action {
			// TODO: 这里列出来是因为后续可能会精细化处理每种情况
			case core.IdxActCreatePost,
				core.IdxActUpdatePost,
				core.IdxActDeletePost,
				core.IdxActStickPost,
				core.IdxActVisiblePost:
				// TODO: 粗糙处理cache，后续需要针对每一种情况精细化处理
				if time.Since(s.lastCacheResetTime) > time.Minute {
					s.cache.Reset()
					s.lastCacheResetTime = time.Now()
					logrus.Debugf("reset cache by %s", action)
				}
			default:
				// nop
			}
		}
	}
}

func (s *bigCacheIndexServant) Name() string {
	return "BigCacheIndex"
}

func (s *bigCacheIndexServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
