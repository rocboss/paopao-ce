// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	dbr "github.com/rocboss/paopao-ce/internal/dao/slonik/ce/postgres"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.IndexPostsService = (*indexPostsServant)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsServant)(nil)
)

type indexPostsServant struct {
	db *pgx.Conn
	q  dbr.Querier
}

type simpleIndexPostsServant struct {
	db *pgx.Conn
	q  dbr.Querier
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *indexPostsServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsServant) IndexPosts(_user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func newIndexPostsService(db *pgx.Conn) core.IndexPostsService {
	return &indexPostsServant{
		db: db,
		q:  dbr.New(db),
	}
}

func newSimpleIndexPostsService(db *pgx.Conn) core.IndexPostsService {
	return &simpleIndexPostsServant{
		db: db,
		q:  dbr.New(db),
	}
}
