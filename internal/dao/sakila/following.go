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
)

var (
	_ core.FollowingManageService = (*followingManageSrv)(nil)
)

type followingManageSrv struct {
	*sqlxSrv
	q *cc.FollowingManager
}

func (s *followingManageSrv) FollowUser(userId int64, followId int64) error {
	// TODO
	return cs.ErrNotImplemented
}

func (s *followingManageSrv) UnfollowUser(userId int64, followId int64) error {
	// TDOO
	return cs.ErrNotImplemented
}

func (s *followingManageSrv) ListFollows(userId int64, limit, offset int) (*ms.ContactList, error) {
	// TODO
	return nil, cs.ErrNotImplemented
}

func (s *followingManageSrv) ListFollowings(userId int64, limit, offset int) (*ms.ContactList, error) {
	// TODO
	return nil, cs.ErrNotImplemented
}

func (s *followingManageSrv) GetFollowCount(userId int64) (int64, int64, error) {
	// TODO
	return 0, 0, cs.ErrNotImplemented
}

func (s *followingManageSrv) IsFollow(userId int64, followId int64) bool {
	// TODO
	return false
}

func newFollowingManageService(db *sqlx.DB) core.FollowingManageService {
	return &followingManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildFollowingManager),
	}
}
