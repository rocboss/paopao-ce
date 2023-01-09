// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/pkg/debug"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	_ core.IndexPostsService = (*friendIndexServant)(nil)
	_ core.IndexPostsService = (*followIndexServant)(nil)
	_ core.IndexPostsService = (*lightIndexServant)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsServant)(nil)
)

type friendIndexServant struct {
	ams core.AuthorizationManageService
	ths core.TweetHelpService
	db  *gorm.DB
}

type followIndexServant struct {
	ths core.TweetHelpService
	db  *gorm.DB
}

type lightIndexServant struct {
	ths core.TweetHelpService
	db  *gorm.DB
}

type simpleIndexPostsServant struct {
	ths core.TweetHelpService
	db  *gorm.DB
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *friendIndexServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
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
		logrus.Debugf("gormIndexPostsServant.IndexPosts err: %v", err)
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

	return &core.IndexTweetList{
		Tweets: formatPosts,
		Total:  total,
	}, nil
}

// IndexPosts 根据userId查询广场推文列表
func (s *followIndexServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，获取公开可见Tweet或者所属用户的私有Tweet
func (s *lightIndexServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	predicates := dbr.Predicates{
		"ORDER": []any{"is_top DESC, latest_replied_on DESC"},
	}
	if user == nil {
		predicates["visibility = ?"] = []any{dbr.PostVisitPublic}
	} else if !user.IsAdmin {
		args := []any{dbr.PostVisitPublic, dbr.PostVisitPrivate, user.ID}
		predicates["visibility = ? OR (visibility = ? AND user_id = ?)"] = args
	}

	posts, err := (&dbr.Post{}).Fetch(s.db, predicates, offset, limit)
	if err != nil {
		logrus.Debugf("gormIndexPostsServant.IndexPosts err: %v", err)
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

	return &core.IndexTweetList{
		Tweets: formatPosts,
		Total:  total,
	}, nil
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsServant) IndexPosts(_user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	predicates := dbr.Predicates{
		"visibility = ?": []any{dbr.PostVisitPublic},
		"ORDER":          []any{"is_top DESC, latest_replied_on DESC"},
	}

	posts, err := (&dbr.Post{}).Fetch(s.db, predicates, offset, limit)
	if err != nil {
		logrus.Debugf("gormSimpleIndexPostsServant.IndexPosts err: %v", err)
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

	return &core.IndexTweetList{
		Tweets: formatPosts,
		Total:  total,
	}, nil
}

func newFriendIndexService(db *gorm.DB, ams core.AuthorizationManageService, ths core.TweetHelpService) core.IndexPostsService {
	return &friendIndexServant{
		ams: ams,
		ths: ths,
		db:  db,
	}
}

func newFollowIndexService(db *gorm.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &followIndexServant{
		ths: ths,
		db:  db,
	}
}

func newLightIndexService(db *gorm.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &lightIndexServant{
		ths: ths,
		db:  db,
	}
}

func newSimpleIndexPostsService(db *gorm.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &simpleIndexPostsServant{
		ths: ths,
		db:  db,
	}
}
