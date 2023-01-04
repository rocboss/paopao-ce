// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
)

var (
	_ core.IndexPostsService = (*indexPostsServant)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsServant)(nil)
)

type indexPostsServant struct {
	db        *sqlx.DB
	stmtIndex *sqlx.Stmt
}

type simpleIndexPostsServant struct {
	db        *sqlx.DB
	stmtIndex *sqlx.Stmt
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *indexPostsServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, nil
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsServant) IndexPosts(_user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	// TODO
	return nil, nil
}

func newIndexPostsService(db *sqlx.DB) core.IndexPostsService {
	return &indexPostsServant{
		db: db,
		stmtIndex: c(`
			SELECT * FROM @person WHERE first_name=?
		`),
	}
}

func newSimpleIndexPostsService(db *sqlx.DB) core.IndexPostsService {
	return &simpleIndexPostsServant{
		db: db,
		stmtIndex: c(`
			SELECT * FROM @person WHERE first_name=?
		`),
	}
}
