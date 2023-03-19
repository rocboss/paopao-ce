// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/pkg/types"
	"gorm.io/gorm"
)

var (
	_ core.AuthorizationManageService = (*authorizationManageSrv)(nil)
)

type authorizationManageSrv struct {
	db *gorm.DB
}

func newAuthorizationManageService(db *gorm.DB) core.AuthorizationManageService {
	return &authorizationManageSrv{
		db: db,
	}
}

func (s *authorizationManageSrv) IsAllow(user *core.User, action *core.Action) bool {
	// user is activation if had bind phone
	isActivation := (len(user.Phone) != 0)
	isFriend := s.isFriend(user.ID, action.UserId)
	// TODO: just use defaut act authorization chek rule now
	return action.Act.IsAllow(user, action.UserId, isFriend, isActivation)
}

func (s *authorizationManageSrv) MyFriendSet(userId int64) core.FriendSet {
	ids, err := (&dbr.Contact{UserId: userId}).MyFriendIds(s.db)
	if err != nil {
		return core.FriendSet{}
	}

	resp := make(core.FriendSet, len(ids))
	for _, id := range ids {
		resp[id] = types.Empty{}
	}
	return resp
}

func (s *authorizationManageSrv) BeFriendFilter(userId int64) core.FriendFilter {
	ids, err := (&dbr.Contact{FriendId: userId}).BeFriendIds(s.db)
	if err != nil {
		return core.FriendFilter{}
	}

	resp := make(core.FriendFilter, len(ids))
	for _, id := range ids {
		resp[id] = types.Empty{}
	}
	return resp
}

func (s *authorizationManageSrv) BeFriendIds(userId int64) ([]int64, error) {
	return (&dbr.Contact{FriendId: userId}).BeFriendIds(s.db)
}

func (s *authorizationManageSrv) isFriend(userId int64, friendId int64) bool {
	contact, err := (&dbr.Contact{UserId: friendId, FriendId: userId}).GetByUserFriend(s.db)
	if err == nil || contact.Status == dbr.ContactStatusAgree {
		return true
	}
	return false
}
