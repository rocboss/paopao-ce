// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.AuthorizationManageService = (*authorizationManageSrv)(nil)
)

type authorizationManageSrv struct {
	*pgxSrv
}

func (s *authorizationManageSrv) IsAllow(user *ms.User, action *ms.Action) bool {
	// TODO
	debug.NotImplemented()
	return false
}

func (s *authorizationManageSrv) MyFriendSet(userId int64) ms.FriendSet {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *authorizationManageSrv) BeFriendFilter(userId int64) ms.FriendFilter {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *authorizationManageSrv) BeFriendIds(userId int64) ([]int64, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *authorizationManageSrv) isFriend(userId int64, friendId int64) bool {
	// TODO
	debug.NotImplemented()
	return false
}

func newAuthorizationManageService(db *pgx.Conn) core.AuthorizationManageService {
	return &authorizationManageSrv{
		pgxSrv: newPgxSrv(db),
	}
}
