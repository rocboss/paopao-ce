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
	_ core.AuthorizationManageService = (*authorizationManageServant)(nil)
)

type authorizationManageServant struct {
	db *gorm.DB
}

func newAuthorizationManageService(db *gorm.DB) core.AuthorizationManageService {
	return &authorizationManageServant{
		db: db,
	}
}

func (s *authorizationManageServant) IsAllow(user *core.User, action *core.Action) bool {
	// user is activation if had bind phone
	isActivation := (len(user.Phone) != 0)
	isFriend := s.isFriend(user.ID, action.UserId)
	// TODO: just use defaut act authorization chek rule now
	return action.Act.IsAllow(user, action.UserId, isFriend, isActivation)
}

func (s *authorizationManageServant) MyFriendSet(userId int64) core.FriendSet {
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

func (s *authorizationManageServant) BeFriendFilter(userId int64) core.FriendFilter {
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

func (s *authorizationManageServant) BeFriendIds(userId int64) ([]int64, error) {
	return (&dbr.Contact{FriendId: userId}).BeFriendIds(s.db)
}

func (s *authorizationManageServant) isFriend(userId int64, friendId int64) bool {
	contact, err := (&dbr.Contact{UserId: friendId, FriendId: userId}).GetByUserFriend(s.db)
	if err == nil || contact.Status == dbr.ContactStatusAgree {
		return true
	}
	return false
}
