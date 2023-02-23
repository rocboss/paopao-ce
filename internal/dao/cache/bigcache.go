// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/allegro/bigcache/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/types"
	"github.com/sirupsen/logrus"
)

var (
	_ core.CacheIndexService = (*bigCacheIndexServant)(nil)
	_ core.VersionInfo       = (*bigCacheIndexServant)(nil)
)

type postsEntry struct {
	key    string
	tweets *core.IndexTweetList
}

type bigCacheIndexServant struct {
	ips core.IndexPostsService
	ams core.AuthorizationManageService

	indexActionCh      chan *core.IndexAction
	cachePostsCh       chan *postsEntry
	cache              *bigcache.BigCache
	lastCacheResetTime time.Time
	preventDuration    time.Duration
}

func (s *bigCacheIndexServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
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

func (s *bigCacheIndexServant) getPosts(key string) (*core.IndexTweetList, error) {
	data, err := s.cache.Get(key)
	if err != nil {
		logrus.Debugf("bigCacheIndexServant.getPosts get posts by key: %s from cache err: %v", key, err)
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var resp core.IndexTweetList
	if err := dec.Decode(&resp); err != nil {
		logrus.Debugf("bigCacheIndexServant.getPosts get posts from cache in decode err: %v", err)
		return nil, err
	}
	return &resp, nil
}

func (s *bigCacheIndexServant) cachePosts(key string, tweets *core.IndexTweetList) {
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

func (s *bigCacheIndexServant) keyFrom(user *core.User, offset int, limit int) string {
	var userId int64 = -1
	if user != nil {
		userId = user.ID
	}
	return fmt.Sprintf("index:%d:%d:%d", userId, offset, limit)
}

func (s *bigCacheIndexServant) SendAction(act core.IdxAct, post *core.Post) {
	action := core.NewIndexAction(act, post)
	select {
	case s.indexActionCh <- action:
		logrus.Debugf("bigCacheIndexServant.SendAction send indexAction by chan: %s", act)
	default:
		go func(ch chan<- *core.IndexAction, act *core.IndexAction) {
			logrus.Debugf("bigCacheIndexServant.SendAction send indexAction by goroutine: %s", action.Act)
			ch <- act
		}(s.indexActionCh, action)
	}
}

func (s *bigCacheIndexServant) startIndexPosts() {
	for {
		select {
		case entry := <-s.cachePostsCh:
			s.setPosts(entry)
		case action := <-s.indexActionCh:
			s.handleIndexAction(action)
		}
	}
}

func (s *bigCacheIndexServant) handleIndexAction(action *core.IndexAction) {
	act, post := action.Act, action.Post

	// 创建/删除 私密推文特殊处理
	switch act {
	case core.IdxActCreatePost, core.IdxActDeletePost:
		if post.Visibility == core.PostVisitPrivate {
			s.deleteCacheByUserId(post.UserID, true)
			return
		}
	}

	// 如果在s.preventDuration时间内就清除所有缓存，否则只清除自个儿的缓存
	// TODO: 需要优化只清除受影响的缓存，后续完善
	if time.Since(s.lastCacheResetTime) > s.preventDuration {
		s.deleteCacheByUserId(post.UserID, false)
	} else {
		s.deleteCacheByUserId(post.UserID, true)
	}
}

func (s *bigCacheIndexServant) deleteCacheByUserId(id int64, oneself bool) {
	var keys []string
	userId := strconv.FormatInt(id, 10)
	friendSet := core.FriendSet{}
	if !oneself {
		friendSet = s.ams.MyFriendSet(id)
	}
	friendSet[userId] = types.Empty{}

	// 获取需要删除缓存的key，目前是仅删除自个儿的缓存
	for it := s.cache.Iterator(); it.SetNext(); {
		entry, err := it.Value()
		if err != nil {
			logrus.Debugf("bigCacheIndexServant.deleteCacheByUserId userId: %s err:%s", userId, err)
			return
		}
		key := entry.Key()
		keyParts := strings.Split(key, ":")
		if len(keyParts) > 2 && keyParts[0] == "index" {
			if _, ok := friendSet[keyParts[1]]; ok {
				keys = append(keys, key)
			}
		}
	}

	// 执行删缓存
	for _, k := range keys {
		s.cache.Delete(k)
	}
	s.lastCacheResetTime = time.Now()
	logrus.Debugf("bigCacheIndexServant.deleteCacheByUserId userId:%s oneself:%t keys:%d", userId, oneself, len(keys))
}

func (s *bigCacheIndexServant) Name() string {
	return "BigCacheIndex"
}

func (s *bigCacheIndexServant) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}
