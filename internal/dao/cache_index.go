package dao

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/sirupsen/logrus"
)

func newSimpleCacheIndexServant(getIndexPosts func(offset, limit int) ([]*model.PostFormated, error)) *simpleCacheIndexServant {
	s := conf.SimpleCacheIndexSetting
	cacheIndex := &simpleCacheIndexServant{
		getIndexPosts:   getIndexPosts,
		maxIndexSize:    s.MaxIndexSize,
		indexPosts:      make([]*model.PostFormated, 0),
		indexActionCh:   make(chan core.IndexActionT, 100),                                 // optimize: size need configure by custom
		checkTick:       time.NewTicker(time.Duration(s.CheckTickDuration) * time.Second),  // check whether need update index every 1 minute
		expireIndexTick: time.NewTicker(time.Duration(s.ExpireTickDuration) * time.Second), // force expire index every 5 minute
	}

	// start index posts
	cacheIndex.atomicIndex.Store(cacheIndex.indexPosts)
	go cacheIndex.startIndexPosts()

	return cacheIndex
}

func (s *simpleCacheIndexServant) IndexPosts(offset int, limit int) ([]*model.PostFormated, bool) {
	posts := s.atomicIndex.Load().([]*model.PostFormated)
	end := offset + limit
	size := len(posts)
	logrus.Debugf("get index posts from posts: %d offset:%d limit:%d start:%d, end:%d", size, offset, limit, offset, end)
	if size >= end {
		return posts[offset:end], true
	}
	return nil, false
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
				if s.indexPosts, err = s.getIndexPosts(0, s.maxIndexSize); err == nil {
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
			case core.IdxActCreatePost, core.IdxActUpdatePost, core.IdxActDeletePost, core.IdxActStickPost:
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
