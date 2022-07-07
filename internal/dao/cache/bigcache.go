package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/allegro/bigcache/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/model/rest"
	"github.com/sirupsen/logrus"
)

var (
	_ core.CacheIndexService = (*bigCacheIndexServant)(nil)
	_ core.VersionInfo       = (*bigCacheIndexServant)(nil)
)

type postsEntry struct {
	key    string
	tweets *rest.IndexTweetsResp
}

type bigCacheIndexServant struct {
	ips core.IndexPostsService

	indexActionCh      chan core.IndexActionT
	cachePostsCh       chan *postsEntry
	cache              *bigcache.BigCache
	lastCacheResetTime time.Time
}

func (s *bigCacheIndexServant) IndexPosts(user *model.User, offset int, limit int) (*rest.IndexTweetsResp, error) {
	key := s.keyFrom(user, offset, limit)
	posts, err := s.getPosts(key)
	if err == nil {
		logrus.Debugf("bigCacheIndexServant.IndexPosts get index posts from cache by key: %s", key)
		return posts, nil
	}

	if posts, err = s.ips.IndexPosts(user, offset, limit); err != nil {
		return nil, err
	}
	logrus.Debugf("bigCacheIndexServant.IndexPosts get index posts from database by key: %s", key)
	s.cachePosts(key, posts)
	return posts, nil
}

func (s *bigCacheIndexServant) getPosts(key string) (*rest.IndexTweetsResp, error) {
	data, err := s.cache.Get(key)
	if err != nil {
		logrus.Debugf("bigCacheIndexServant.getPosts get posts by key: %s from cache err: %v", key, err)
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var resp rest.IndexTweetsResp
	if err := dec.Decode(&resp); err != nil {
		logrus.Debugf("bigCacheIndexServant.getPosts get posts from cache in decode err: %v", err)
		return nil, err
	}
	return &resp, nil
}

func (s *bigCacheIndexServant) cachePosts(key string, tweets *rest.IndexTweetsResp) {
	entry := &postsEntry{key: key, tweets: tweets}
	select {
	case s.cachePostsCh <- entry:
		logrus.Debugf("bigCacheIndexServant.cachePosts cachePosts by chan of key: %s", key)
	default:
		go func(ch chan<- *postsEntry, entry *postsEntry) {
			logrus.Debugf("bigCacheIndexServant.cachePosts cachePosts indexAction by goroutine of key: %s", key)
			ch <- entry
		}(s.cachePostsCh, entry)
	}
}

func (s *bigCacheIndexServant) setPosts(entry *postsEntry) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(entry.tweets); err != nil {
		logrus.Debugf("bigCacheIndexServant.setPosts setPosts encode post entry err: %v", err)
		return
	}
	if err := s.cache.Set(entry.key, buf.Bytes()); err != nil {
		logrus.Debugf("bigCacheIndexServant.setPosts setPosts set cache err: %v", err)
	}
	logrus.Debugf("bigCacheIndexServant.setPosts setPosts set cache by key: %s", entry.key)
}

func (s *bigCacheIndexServant) keyFrom(user *model.User, offset int, limit int) string {
	var userId int64 = -1
	if user != nil {
		userId = user.ID
	}
	return fmt.Sprintf("index:%d:%d:%d", userId, offset, limit)
}

func (s *bigCacheIndexServant) SendAction(act core.IndexActionT) {
	select {
	case s.indexActionCh <- act:
		logrus.Debugf("bigCacheIndexServant.SendAction send indexAction by chan: %s", act)
	default:
		go func(ch chan<- core.IndexActionT, act core.IndexActionT) {
			logrus.Debugf("bigCacheIndexServant.SendAction send indexAction by goroutine: %s", act)
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
