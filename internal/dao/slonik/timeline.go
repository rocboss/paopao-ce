// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.IndexPostsService = (*friendIndexSrv)(nil)
	_ core.IndexPostsService = (*followIndexSrv)(nil)
	_ core.IndexPostsService = (*lightIndexSrv)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsSrv)(nil)
)

type friendIndexSrv struct {
	*pgxSrv
	ams core.AuthorizationManageService
	ths core.TweetHelpService
}

type followIndexSrv struct {
	*pgxSrv
	ths core.TweetHelpService
}

type lightIndexSrv struct {
	*pgxSrv
	ths core.TweetHelpService
}

type simpleIndexPostsSrv struct {
	*pgxSrv
	ths core.TweetHelpService
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *friendIndexSrv) IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *friendIndexSrv) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *followIndexSrv) IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *followIndexSrv) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *lightIndexSrv) IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *lightIndexSrv) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsSrv) IndexPosts(_user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *simpleIndexPostsSrv) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func newFriendIndexService(db *pgx.Conn, ams core.AuthorizationManageService, ths core.TweetHelpService) core.IndexPostsService {
	return &friendIndexSrv{
		ams:    ams,
		pgxSrv: newPgxSrv(db),
	}
}

func newFollowIndexService(db *pgx.Conn, ths core.TweetHelpService) core.IndexPostsService {
	return &followIndexSrv{
		ths:    ths,
		pgxSrv: newPgxSrv(db),
	}
}

func newLightIndexService(db *pgx.Conn, ths core.TweetHelpService) core.IndexPostsService {
	return &lightIndexSrv{
		ths:    ths,
		pgxSrv: newPgxSrv(db),
	}
}

func newSimpleIndexPostsService(db *pgx.Conn, ths core.TweetHelpService) core.IndexPostsService {
	return &simpleIndexPostsSrv{
		ths:    ths,
		pgxSrv: newPgxSrv(db),
	}
}
