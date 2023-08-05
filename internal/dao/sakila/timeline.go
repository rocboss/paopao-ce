// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.IndexPostsService = (*friendIndexSrv)(nil)
	_ core.IndexPostsService = (*followIndexSrv)(nil)
	_ core.IndexPostsService = (*lightIndexSrv)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsSrv)(nil)

	_ core.IndexPostsServantA = (*friendIndexSrvA)(nil)
	_ core.IndexPostsServantA = (*followIndexSrvA)(nil)
	_ core.IndexPostsServantA = (*lightIndexSrvA)(nil)
	_ core.IndexPostsServantA = (*simpleIndexPostsSrvA)(nil)
)

type friendIndexSrv struct {
	*sqlxSrv
	ams core.AuthorizationManageService
	ths core.TweetHelpService
	q   *cc.FriendIndex
}

type followIndexSrv struct {
	*sqlxSrv
	ths core.TweetHelpService
	q   *cc.FollowIndex
}

type lightIndexSrv struct {
	*sqlxSrv
	ths core.TweetHelpService
	q   *cc.LightIndex
}

type simpleIndexPostsSrv struct {
	*sqlxSrv
	ths core.TweetHelpService
	q   *cc.SimpleIndex
}

type friendIndexSrvA struct {
	*sqlxSrv
	ams core.AuthorizationManageService
	ths core.TweetHelpServantA
	q   *cc.FriendIndexA
}

type followIndexSrvA struct {
	*sqlxSrv
	ths core.TweetHelpServantA
	q   *cc.FollowIndexA
}

type lightIndexSrvA struct {
	*sqlxSrv
	ths core.TweetHelpServantA
	q   *cc.LightIndexA
}

type simpleIndexPostsSrvA struct {
	*sqlxSrv
	ths core.TweetHelpServantA
	q   *cc.SimpleIndexA
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *friendIndexSrv) IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *followIndexSrv) IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *lightIndexSrv) IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsSrv) IndexPosts(_user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *friendIndexSrvA) IndexPosts(user *ms.User, limit int, offset int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *followIndexSrvA) IndexPosts(user *ms.User, limit int, offset int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *lightIndexSrvA) IndexPosts(user *ms.User, limit int, offset int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsSrvA) IndexPosts(_user *ms.User, limit int, offset int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func newFriendIndexService(db *sqlx.DB, ams core.AuthorizationManageService, ths core.TweetHelpService) core.IndexPostsService {
	return &friendIndexSrv{
		ams:     ams,
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildFriendIndex),
	}
}

func newFollowIndexService(db *sqlx.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &followIndexSrv{
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildFollowIndex),
	}
}

func newLightIndexService(db *sqlx.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &lightIndexSrv{
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildLightIndex),
	}
}

func newSimpleIndexPostsService(db *sqlx.DB, ths core.TweetHelpService) core.IndexPostsService {
	return &simpleIndexPostsSrv{
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildSimpleIndex),
	}
}

func newFriendIndexServantA(db *sqlx.DB, ams core.AuthorizationManageService, ths core.TweetHelpServantA) core.IndexPostsServantA {
	return &friendIndexSrvA{
		ams:     ams,
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildFriendIndexA),
	}
}

func newFollowIndexServantA(db *sqlx.DB, ths core.TweetHelpServantA) core.IndexPostsServantA {
	return &followIndexSrvA{
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildFollowIndexA),
	}
}

func newLightIndexServantA(db *sqlx.DB, ths core.TweetHelpServantA) core.IndexPostsServantA {
	return &lightIndexSrvA{
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildLightIndexA),
	}
}

func newSimpleIndexPostsServantA(db *sqlx.DB, ths core.TweetHelpServantA) core.IndexPostsServantA {
	return &simpleIndexPostsSrvA{
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildSimpleIndexA),
	}
}
