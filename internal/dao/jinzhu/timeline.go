// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/pkg/debug"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type shipIndexSrv struct {
	ams core.AuthorizationManageService
	ths core.TweetHelpService
	db  *gorm.DB
}

type simpleIndexPostsSrv struct {
	ths core.TweetHelpService
	db  *gorm.DB
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *shipIndexSrv) IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	predicates := dbr.Predicates{
		"ORDER": []any{"is_top DESC, latest_replied_on DESC"},
	}
	if user == nil {
		predicates["visibility = ?"] = []any{dbr.PostVisitPublic}
	} else if !user.IsAdmin {
		friendIds, _ := s.ams.BeFriendIds(user.ID)
		friendIds = append(friendIds, user.ID)
		args := []any{dbr.PostVisitPublic, dbr.PostVisitPrivate, user.ID, dbr.PostVisitFriend, friendIds}
		predicates["visibility = ? OR (visibility = ? AND user_id = ?) OR (visibility = ? AND user_id IN ?)"] = args
	}

	posts, err := (&dbr.Post{}).Fetch(s.db, predicates, offset, limit)
	if err != nil {
		logrus.Debugf("gormIndexPostsSrv.IndexPosts err: %v", err)
		return nil, err
	}
	formatPosts, err := s.ths.MergePosts(posts)
	if err != nil {
		return nil, err
	}

	total, err := (&dbr.Post{}).CountBy(s.db, predicates)
	if err != nil {
		return nil, err
	}

	return &ms.IndexTweetList{
		Tweets: formatPosts,
		Total:  total,
	}, nil
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsSrv) IndexPosts(_user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	predicates := dbr.Predicates{
		"visibility = ?": []any{dbr.PostVisitPublic},
		"ORDER":          []any{"is_top DESC, latest_replied_on DESC"},
	}

	posts, err := (&dbr.Post{}).Fetch(s.db, predicates, offset, limit)
	if err != nil {
		logrus.Debugf("gormSimpleIndexPostsSrv.IndexPosts err: %v", err)
		return nil, err
	}

	formatPosts, err := s.ths.MergePosts(posts)
	if err != nil {
		return nil, err
	}

	total, err := (&dbr.Post{}).CountBy(s.db, predicates)
	if err != nil {
		return nil, err
	}

	return &ms.IndexTweetList{
		Tweets: formatPosts,
		Total:  total,
	}, nil
}

func (s *simpleIndexPostsSrv) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func newShipIndexService(db *gorm.DB, ams core.AuthorizationManageService, ths core.TweetHelpService) core.IndexPostsService {
	return &shipIndexSrv{
		ams: ams,
		ths: ths,
		db:  db,
	}
}

func newSimpleIndexPostsService(db *gorm.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &simpleIndexPostsSrv{
		ths: ths,
		db:  db,
	}
}
