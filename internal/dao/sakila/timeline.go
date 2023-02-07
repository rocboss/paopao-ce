// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.IndexPostsService = (*friendIndexServant)(nil)
	_ core.IndexPostsService = (*followIndexServant)(nil)
	_ core.IndexPostsService = (*lightIndexServant)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsServant)(nil)
)

type friendIndexServant struct {
	*sqlxServant
	ams       core.AuthorizationManageService
	ths       core.TweetHelpService
	stmtIndex *sqlx.Stmt
}

type followIndexServant struct {
	*sqlxServant
	ths       core.TweetHelpService
	stmtIndex *sqlx.Stmt
}

type lightIndexServant struct {
	*sqlxServant
	ths       core.TweetHelpService
	stmtIndex *sqlx.Stmt
}

type simpleIndexPostsServant struct {
	*sqlxServant
	ths       core.TweetHelpService
	stmtIndex *sqlx.Stmt
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *friendIndexServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *friendIndexServant) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *followIndexServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *followIndexServant) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *lightIndexServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *lightIndexServant) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsServant) IndexPosts(_user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *simpleIndexPostsServant) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func newFriendIndexService(db *sqlx.DB, ams core.AuthorizationManageService, ths core.TweetHelpService) core.IndexPostsService {
	return &friendIndexServant{
		ams:         ams,
		sqlxServant: newSqlxServant(db),
		stmtIndex:   c(`SELECT * FROM @person WHERE first_name=?`),
	}
}

func newFollowIndexService(db *sqlx.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &followIndexServant{
		ths:         ths,
		sqlxServant: newSqlxServant(db),
		stmtIndex:   c(`SELECT * FROM @person WHERE first_name=?`),
	}
}

func newLightIndexService(db *sqlx.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &lightIndexServant{
		ths:         ths,
		sqlxServant: newSqlxServant(db),
		stmtIndex:   c(`SELECT * FROM @person WHERE first_name=?`),
	}
}

func newSimpleIndexPostsService(db *sqlx.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &simpleIndexPostsServant{
		ths:         ths,
		sqlxServant: newSqlxServant(db),
		stmtIndex:   c(`SELECT * FROM @person WHERE first_name=?`),
	}
}
