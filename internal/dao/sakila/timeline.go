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
	_ core.IndexPostsService = (*friendIndexSrv)(nil)
	_ core.IndexPostsService = (*followIndexSrv)(nil)
	_ core.IndexPostsService = (*lightIndexSrv)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsSrv)(nil)
)

type friendIndexSrv struct {
	*sqlxSrv
	ams       core.AuthorizationManageService
	ths       core.TweetHelpService
	stmtIndex *sqlx.Stmt
}

type followIndexSrv struct {
	*sqlxSrv
	ths       core.TweetHelpService
	stmtIndex *sqlx.Stmt
}

type lightIndexSrv struct {
	*sqlxSrv
	ths       core.TweetHelpService
	stmtIndex *sqlx.Stmt
}

type simpleIndexPostsSrv struct {
	*sqlxSrv
	ths       core.TweetHelpService
	stmtIndex *sqlx.Stmt
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *friendIndexSrv) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *friendIndexSrv) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *followIndexSrv) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *followIndexSrv) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *lightIndexSrv) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *lightIndexSrv) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsSrv) IndexPosts(_user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *simpleIndexPostsSrv) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func newFriendIndexService(db *sqlx.DB, ams core.AuthorizationManageService, ths core.TweetHelpService) core.IndexPostsService {
	return &friendIndexSrv{
		ams:       ams,
		sqlxSrv:   newSqlxSrv(db),
		stmtIndex: c(`SELECT * FROM @user WHERE username=?`),
	}
}

func newFollowIndexService(db *sqlx.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &followIndexSrv{
		ths:       ths,
		sqlxSrv:   newSqlxSrv(db),
		stmtIndex: c(`SELECT * FROM @user WHERE username=?`),
	}
}

func newLightIndexService(db *sqlx.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &lightIndexSrv{
		ths:       ths,
		sqlxSrv:   newSqlxSrv(db),
		stmtIndex: c(`SELECT * FROM @user WHERE username=?`),
	}
}

func newSimpleIndexPostsService(db *sqlx.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &simpleIndexPostsSrv{
		ths:       ths,
		sqlxSrv:   newSqlxSrv(db),
		stmtIndex: c(`SELECT * FROM @user WHERE username=?`),
	}
}
