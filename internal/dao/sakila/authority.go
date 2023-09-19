// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
	"github.com/rocboss/paopao-ce/pkg/types"
	"github.com/sirupsen/logrus"
)

var (
	_ core.AuthorizationManageService = (*authorizationManageSrv)(nil)
)

type authorizationManageSrv struct {
	*sqlxSrv
	q *cc.AuthorizationManage
}

func (s *authorizationManageSrv) IsAllow(user *ms.User, action *ms.Action) bool {
	// user is activation if had bind phone
	isActivation := (len(user.Phone) != 0)
	isFriend := s.isFriend(user.ID, action.UserId)
	// TODO: just use defaut act authorization chek rule now
	return action.Act.IsAllow(user, action.UserId, isFriend, isActivation)
}

func (s *authorizationManageSrv) MyFriendSet(userId int64) ms.FriendSet {
	var ids []string
	if err := s.q.MyFriendSet.Get(&ids, userId); err != nil {
		logrus.Warnf("get my FriendSet error: %s", err)
		return ms.FriendSet{}
	}
	resp := make(ms.FriendSet, len(ids))
	for _, id := range ids {
		resp[id] = types.Empty{}
	}
	return resp
}

func (s *authorizationManageSrv) BeFriendFilter(userId int64) ms.FriendFilter {
	var ids []int64
	if err := s.q.BeFriendIds.Get(&ids, userId); err != nil {
		logrus.Warnf("get my BeFriendFilter error: %s", err)
		return ms.FriendFilter{}
	}
	resp := make(ms.FriendFilter, len(ids))
	for _, id := range ids {
		resp[id] = types.Empty{}
	}
	return resp
}

func (s *authorizationManageSrv) BeFriendIds(userId int64) (res []int64, err error) {
	err = s.q.BeFriendIds.Get(&res, userId)
	return
}

func (s *authorizationManageSrv) isFriend(userId int64, friendId int64) bool {
	var status int8
	err := s.q.IsFriend.Get(&status, userId, friendId)
	if err == nil || status == cs.ContactStatusAgree {
		return true
	}
	return false
}

func newAuthorizationManageService(db *sqlx.DB) core.AuthorizationManageService {
	return &authorizationManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildAuthorizationManage),
	}
}
