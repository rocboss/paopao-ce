// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.AuthorizationManageService = (*authorizationManageServant)(nil)
)

type authorizationManageServant struct {
	db               *sqlx.DB
	stmtIdx          *sqlx.Stmt
	stmtUpdateFriend *sqlx.Stmt
}

func (s *authorizationManageServant) IsAllow(user *core.User, action *core.Action) bool {
	// TODO
	debug.NotImplemented()
	return false
}

func (s *authorizationManageServant) MyFriendSet(userId int64) core.FriendSet {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *authorizationManageServant) BeFriendFilter(userId int64) core.FriendFilter {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *authorizationManageServant) BeFriendIds(userId int64) ([]int64, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *authorizationManageServant) isFriend(userId int64, friendId int64) bool {
	// TODO
	debug.NotImplemented()
	return false
}

func newAuthorizationManageService(db *sqlx.DB) core.AuthorizationManageService {
	return &authorizationManageServant{
		db:               db,
		stmtIdx:          c(`SELECT * FROM @person WHERE first_name=?`),
		stmtUpdateFriend: c(`SELECT * FROM @person WHERE first_name=?`),
	}
}
