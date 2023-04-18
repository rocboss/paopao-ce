// Copyright 2023 ROC. All rights reserved.
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
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/types"
	"github.com/sirupsen/logrus"
)

const (
	_cacheIndexKey = "paopao_index"
)

var (
	_ core.CacheIndexService = (*cacheIndexSrv)(nil)
	_ core.VersionInfo       = (*cacheIndexSrv)(nil)
)

type postsEntry struct {
	key    string
	tweets *core.IndexTweetList
}

type tweetsCache interface {
	core.VersionInfo
	getTweetsBytes(key string) ([]byte, error)
	setTweetsBytes(key string, bs []byte) error
	delTweets(keys []string) error
	allKeys() ([]string, error)
}

type cacheIndexSrv struct {
	ips core.IndexPostsService
	ams core.AuthorizationManageService

	name               string
	version            *semver.Version
	indexActionCh      chan *core.IndexAction
	cachePostsCh       chan *postsEntry
	cache              tweetsCache
	lastCacheResetTime time.Time
	preventDuration    time.Duration
}

func (s *cacheIndexSrv) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	key := s.keyFrom(user, offset, limit)
	posts, err := s.getPosts(key)
	if err == nil {
		logrus.Debugf("cacheIndexSrv.IndexPosts get index posts from cache by key: %s", key)
		return posts, nil
	}

	if posts, err = s.ips.IndexPosts(user, offset, limit); err != nil {
		return nil, err
	}
	logrus.Debugf("cacheIndexSrv.IndexPosts get index posts from database by key: %s", key)
	s.cachePosts(key, posts)
	return posts, nil
}

func (s *cacheIndexSrv) getPosts(key string) (*core.IndexTweetList, error) {
	data, err := s.cache.getTweetsBytes(key)
	if err != nil {
		logrus.Debugf("cacheIndexSrv.getPosts get posts by key: %s from cache err: %v", key, err)
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var resp core.IndexTweetList
	if err := dec.Decode(&resp); err != nil {
		logrus.Debugf("cacheIndexSrv.getPosts get posts from cache in decode err: %v", err)
		return nil, err
	}
	return &resp, nil
}

func (s *cacheIndexSrv) cachePosts(key string, tweets *core.IndexTweetList) {
	entry := &postsEntry{key: key, tweets: tweets}
	select {
	case s.cachePostsCh <- entry:
		logrus.Debugf("cacheIndexSrv.cachePosts cachePosts by chan of key: %s", key)
	default:
		go func(ch chan<- *postsEntry, entry *postsEntry) {
			logrus.Debugf("cacheIndexSrv.cachePosts cachePosts indexAction by goroutine of key: %s", key)
			ch <- entry
		}(s.cachePostsCh, entry)
	}
}

func (s *cacheIndexSrv) setPosts(entry *postsEntry) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(entry.tweets); err != nil {
		logrus.Debugf("cacheIndexSrv.setPosts setPosts encode post entry err: %v", err)
		return
	}
	if err := s.cache.setTweetsBytes(entry.key, buf.Bytes()); err != nil {
		logrus.Debugf("cacheIndexSrv.setPosts setPosts set cache err: %v", err)
	}
	logrus.Debugf("cacheIndexSrv.setPosts setPosts set cache by key: %s", entry.key)
}

func (s *cacheIndexSrv) keyFrom(user *core.User, offset int, limit int) string {
	var userId int64 = -1
	if user != nil {
		userId = user.ID
	}
	return fmt.Sprintf("%s:%d:%d:%d", _cacheIndexKey, userId, offset, limit)
}

func (s *cacheIndexSrv) SendAction(act core.IdxAct, post *core.Post) {
	action := core.NewIndexAction(act, post)
	select {
	case s.indexActionCh <- action:
		logrus.Debugf("cacheIndexSrv.SendAction send indexAction by chan: %s", act)
	default:
		go func(ch chan<- *core.IndexAction, act *core.IndexAction) {
			logrus.Debugf("cacheIndexSrv.SendAction send indexAction by goroutine: %s", action.Act)
			ch <- act
		}(s.indexActionCh, action)
	}
}

func (s *cacheIndexSrv) startIndexPosts() {
	for {
		select {
		case entry := <-s.cachePostsCh:
			s.setPosts(entry)
		case action := <-s.indexActionCh:
			s.handleIndexAction(action)
		}
	}
}

func (s *cacheIndexSrv) handleIndexAction(action *core.IndexAction) {
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

func (s *cacheIndexSrv) deleteCacheByUserId(id int64, oneself bool) {
	var keys []string
	userId := strconv.FormatInt(id, 10)
	friendSet := core.FriendSet{}
	if !oneself {
		friendSet = s.ams.MyFriendSet(id)
	}
	friendSet[userId] = types.Empty{}

	// 获取需要删除缓存的key，目前是仅删除自个儿的缓存
	allKeys, err := s.cache.allKeys()
	if err != nil {
		logrus.Debugf("cacheIndexSrv.deleteCacheByUserId userId: %s err:%s", userId, err)
	}
	for _, key := range allKeys {
		keyParts := strings.Split(key, ":")
		if len(keyParts) > 2 && keyParts[0] == _cacheIndexKey {
			if _, ok := friendSet[keyParts[1]]; ok {
				keys = append(keys, key)
			}
		}
	}

	// 执行删缓存
	s.cache.delTweets(keys)
	s.lastCacheResetTime = time.Now()
	logrus.Debugf("cacheIndexSrv.deleteCacheByUserId userId:%s oneself:%t keys:%d", userId, oneself, len(keys))
}

func (s *cacheIndexSrv) Name() string {
	return s.name
}

func (s *cacheIndexSrv) Version() *semver.Version {
	return s.version
}

func newCacheIndexSrv(ips core.IndexPostsService, ams core.AuthorizationManageService, tc tweetsCache) *cacheIndexSrv {
	cacheIndex := &cacheIndexSrv{
		ips:             ips,
		ams:             ams,
		cache:           tc,
		name:            tc.Name(),
		version:         tc.Version(),
		preventDuration: 10 * time.Second,
	}

	// indexActionCh capacity custom configure by conf.yaml need in [10, 10000]
	// or re-compile source to adjust min/max capacity
	capacity := conf.CacheIndexSetting.MaxUpdateQPS
	if capacity < 10 {
		capacity = 10
	} else if capacity > 10000 {
		capacity = 10000
	}
	cacheIndex.indexActionCh = make(chan *core.IndexAction, capacity)
	cacheIndex.cachePostsCh = make(chan *postsEntry, capacity)
	// 启动索引更新器
	go cacheIndex.startIndexPosts()

	return cacheIndex
}
