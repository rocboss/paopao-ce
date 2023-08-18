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
	_ core.IndexPostsService = (*shipIndexSrv)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsSrv)(nil)
)

type shipIndexSrv struct {
	*pgxSrv
	ams core.AuthorizationManageService
	ths core.TweetHelpService
}

type simpleIndexPostsSrv struct {
	*pgxSrv
	ths core.TweetHelpService
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *shipIndexSrv) IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
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

func newShipIndexService(db *pgx.Conn, ams core.AuthorizationManageService, ths core.TweetHelpService) core.IndexPostsService {
	return &shipIndexSrv{
		ams:    ams,
		pgxSrv: newPgxSrv(db),
	}
}

func newSimpleIndexPostsService(db *pgx.Conn, ths core.TweetHelpService) core.IndexPostsService {
	return &simpleIndexPostsSrv{
		ths:    ths,
		pgxSrv: newPgxSrv(db),
	}
}
