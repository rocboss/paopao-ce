// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/sirupsen/logrus"
)

type eventCacheIndexSrv struct {
	tms core.TweetMetricServantA
}

func (s *eventCacheIndexSrv) SendAction(act core.IdxAct, post *ms.Post) {
	err := error(nil)
	switch act {
	case core.IdxActUpdatePost:
		err = s.tms.UpdateRankScore(&cs.TweetMetric{
			PostId:          post.ID,
			CommentCount:    post.CommentCount,
			UpvoteCount:     post.UpvoteCount,
			CollectionCount: post.CollectionCount,
			ShareCount:      post.ShareCount,
		})
		OnExpireIndexTweetEvent(post.UserID)
	case core.IdxActCreatePost:
		err = s.tms.AddTweetMetric(post.ID)
		OnExpireIndexTweetEvent(post.UserID)
	case core.IdxActDeletePost:
		err = s.tms.DeleteTweetMetric(post.ID)
		OnExpireIndexTweetEvent(post.UserID)
	case core.IdxActStickPost, core.IdxActVisiblePost:
		OnExpireIndexTweetEvent(post.UserID)
	}
	if err != nil {
		logrus.Errorf("eventCacheIndexSrv.SendAction(%s) occurs error: %s", act, err)
	}
}

func NewEventCacheIndexSrv(tms core.TweetMetricServantA) core.CacheIndexService {
	lazyInitial()
	return &eventCacheIndexSrv{
		tms: tms,
	}
}
