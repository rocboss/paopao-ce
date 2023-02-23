// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"sync/atomic"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

var (
	_ core.CacheIndexService = (*simpleCacheIndexServant)(nil)
	_ core.VersionInfo       = (*simpleCacheIndexServant)(nil)
)

type simpleCacheIndexServant struct {
	ips core.IndexPostsService

	indexActionCh   chan core.IdxAct
	indexPosts      *core.IndexTweetList
	atomicIndex     atomic.Value
	maxIndexSize    int
	checkTick       *time.Ticker
	expireIndexTick *time.Ticker
}

func (s *simpleCacheIndexServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	cacheResp := s.atomicIndex.Load().(*core.IndexTweetList)
	end := offset + limit
	if cacheResp != nil {
		size := len(cacheResp.Tweets)
		logrus.Debugf("simpleCacheIndexServant.IndexPosts get index posts from cache posts: %d offset:%d limit:%d start:%d, end:%d", size, offset, limit, offset, end)
		if size >= end {
			return &core.IndexTweetList{
				Tweets: cacheResp.Tweets[offset:end],
				Total:  cacheResp.Total,
			}, nil
		}
	}

	logrus.Debugln("simpleCacheIndexServant.IndexPosts get index posts from database")
	return s.ips.IndexPosts(user, offset, limit)
}

func (s *simpleCacheIndexServant) SendAction(act core.IdxAct, _post *core.Post) {
	select {
	case s.indexActionCh <- act:
		logrus.Debugf("simpleCacheIndexServant.SendAction send indexAction by chan: %s", act)
	default:
		go func(ch chan<- core.IdxAct, act core.IdxAct) {
			logrus.Debugf("simpleCacheIndexServant.SendAction send indexAction by goroutine: %s", act)
			ch <- act
		}(s.indexActionCh, act)
	}
}

func (s *simpleCacheIndexServant) startIndexPosts() {
	var err error
	for {
		select {
		case <-s.checkTick.C:
			if s.indexPosts == nil {
				logrus.Debugf("index posts by checkTick")
				if s.indexPosts, err = s.ips.IndexPosts(nil, 0, s.maxIndexSize); err == nil {
					s.atomicIndex.Store(s.indexPosts)
				} else {
					logrus.Errorf("get index posts err: %v", err)
				}
			}
		case <-s.expireIndexTick.C:
			logrus.Debugf("expire index posts by expireIndexTick")
			if s.indexPosts != nil {
				s.indexPosts = nil
				s.atomicIndex.Store(s.indexPosts)
			}
		case action := <-s.indexActionCh:
			switch action {
			// TODO: 这里列出来是因为后续可能会精细化处理每种情况
			case core.IdxActCreatePost,
				core.IdxActUpdatePost,
				core.IdxActDeletePost,
				core.IdxActStickPost,
				core.IdxActVisiblePost:
				// prevent many update post in least time
				if s.indexPosts != nil {
					logrus.Debugf("remove index posts by action %s", action)
					s.indexPosts = nil
					s.atomicIndex.Store(s.indexPosts)
				}
			default:
				// nop
			}
		}
	}
}

func (s *simpleCacheIndexServant) Name() string {
	return "SimpleCacheIndex"
}

func (s *simpleCacheIndexServant) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}
