// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/ac"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.IndexPostsServantA = (*shipIndexSrvA)(nil)
	_ core.IndexPostsServantA = (*simpleIndexPostsSrvA)(nil)
)

type shipIndexSrvA struct {
	*sqlxSrv
	ams core.AuthorizationManageService
	ths core.TweetHelpServantA
	q   *ac.ShipIndexA
}

type simpleIndexPostsSrvA struct {
	*sqlxSrv
	ths core.TweetHelpServantA
	q   *ac.SimpleIndexA
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *shipIndexSrvA) IndexPosts(user *ms.User, limit int, offset int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

// IndexPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsSrvA) IndexPosts(_user *ms.User, limit int, offset int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

//lint:ignore U1000 newShipIndexServantA
func newShipIndexServantA(db *sqlx.DB, ams core.AuthorizationManageService, ths core.TweetHelpServantA) core.IndexPostsServantA {
	return &shipIndexSrvA{
		ams:     ams,
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       acBuildFn(db, ac.BuildShipIndexA),
	}
}

//lint:ignore U1000 newSimpleIndexPostsServantA
func newSimpleIndexPostsServantA(db *sqlx.DB, ths core.TweetHelpServantA) core.IndexPostsServantA {
	return &simpleIndexPostsSrvA{
		ths:     ths,
		sqlxSrv: newSqlxSrv(db),
		q:       acBuildFn(db, ac.BuildSimpleIndexA),
	}
}
