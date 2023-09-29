// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/slonik/sqlc/auto/pgc"
	"github.com/rocboss/paopao-ce/pkg/types"
	"github.com/sirupsen/logrus"
)

var (
	_ core.AuthorizationManageService = (*authorizationManageSrv)(nil)
)

type authorizationManageSrv struct {
	*pgcSrv
}

func (s *authorizationManageSrv) IsAllow(user *ms.User, action *ms.Action) bool {
	// user is activation if had bind phone
	isActivation := (len(user.Phone) != 0)
	isFriend := s.isFriend(user.ID, action.UserId)
	// TODO: just use defaut act authorization chek rule now
	return action.Act.IsAllow(user, action.UserId, isFriend, isActivation)
}

func (s *authorizationManageSrv) MyFriendSet(userId int64) ms.FriendSet {
	ids, err := s.q.MyFriendSet(context.Background(), userId)
	if err != nil {
		logrus.Warnf("get my FriendSet error: %s", err)
		return ms.FriendSet{}
	}
	resp := make(ms.FriendSet, len(ids))
	for _, id := range ids {
		resp[strconv.FormatInt(id, 10)] = types.Empty{}
	}
	return resp
}

func (s *authorizationManageSrv) BeFriendFilter(userId int64) ms.FriendFilter {
	ids, err := s.q.BeFriendIds(context.Background(), userId)
	if err != nil {
		logrus.Warnf("get my BeFriendFilter error: %s", err)
		return ms.FriendFilter{}
	}
	resp := make(ms.FriendFilter, len(ids))
	for _, id := range ids {
		resp[id] = types.Empty{}
	}
	return resp
}

func (s *authorizationManageSrv) BeFriendIds(userId int64) ([]int64, error) {
	return s.q.BeFriendIds(context.Background(), userId)
}

func (s *authorizationManageSrv) isFriend(userId int64, friendId int64) bool {
	ok, _ := s.q.IsFriend(context.Background(), &pgc.IsFriendParams{
		UserID:   userId,
		FriendID: friendId,
	})
	return ok
}

func newAuthorizationManageService(db *pgx.Conn) core.AuthorizationManageService {
	return &authorizationManageSrv{
		pgcSrv: newPgcSrv(db),
	}
}
