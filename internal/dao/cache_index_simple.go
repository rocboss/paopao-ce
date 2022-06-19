package dao

import (
	"errors"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/sirupsen/logrus"
)

var (
	errNotExist = errors.New("index posts cache not exist")
)

func newSimpleCacheIndexServant(getIndexPosts indexPostsFunc) *simpleCacheIndexServant {
	s := conf.SimpleCacheIndexSetting
	cacheIndex := &simpleCacheIndexServant{
		getIndexPosts:   getIndexPosts,
		maxIndexSize:    s.MaxIndexSize,
		indexPosts:      make([]*model.PostFormated, 0),
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
	capacity := s.ActionQPS
	if capacity < 10 {
		capacity = 10
	} else if capacity > 10000 {
		capacity = 10000
	}
	cacheIndex.indexActionCh = make(chan core.IndexActionT, capacity)

	// start index posts
	cacheIndex.atomicIndex.Store(cacheIndex.indexPosts)
	go cacheIndex.startIndexPosts()

	return cacheIndex
}

func (s *simpleCacheIndexServant) IndexPosts(_userId int64, offset int, limit int) ([]*model.PostFormated, error) {
	posts := s.atomicIndex.Load().([]*model.PostFormated)
	end := offset + limit
	size := len(posts)
	logrus.Debugf("get index posts from posts: %d offset:%d limit:%d start:%d, end:%d", size, offset, limit, offset, end)
	if size >= end {
		return posts[offset:end], nil
	}
	return nil, errNotExist
}

func (s *simpleCacheIndexServant) SendAction(act core.IndexActionT) {
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

func (s *simpleCacheIndexServant) startIndexPosts() {
	var err error
	for {
		select {
		case <-s.checkTick.C:
			if len(s.indexPosts) == 0 {
				logrus.Debugf("index posts by checkTick")
				if s.indexPosts, err = s.getIndexPosts(0, 0, s.maxIndexSize); err == nil {
					s.atomicIndex.Store(s.indexPosts)
				} else {
					logrus.Errorf("get index posts err: %v", err)
				}
			}
		case <-s.expireIndexTick.C:
			logrus.Debugf("expire index posts by expireIndexTick")
			if len(s.indexPosts) != 0 {
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
				if len(s.indexPosts) != 0 {
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
	return semver.MustParse("v0.1.0")
}
