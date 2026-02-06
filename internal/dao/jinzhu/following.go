// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var _ core.FollowingManageService = (*followingManageSrv)(nil)

type followingManageSrv struct {
	db *gorm.DB
	f  *dbr.Following
	u  *dbr.User
}

func newFollowingManageService(db *gorm.DB) core.FollowingManageService {
	return &followingManageSrv{
		db: db,
		f:  &dbr.Following{},
		u:  &dbr.User{},
	}
}

func (s *followingManageSrv) FollowUser(userId int64, followId int64) error {
	if _, err := s.f.GetFollowing(s.db, userId, followId); err != nil {
		following := &dbr.Following{
			UserId:   userId,
			FollowId: followId,
		}
		if _, err = following.Create(s.db); err != nil {
			logrus.Errorf("contactManageSrv.fetchOrNewContact create new contact err:%s", err)
			return err
		}
	}
	return nil
}

func (s *followingManageSrv) UnfollowUser(userId int64, followId int64) error {
	return s.f.DelFollowing(s.db, userId, followId)
}

func (s *followingManageSrv) ListFollows(userId int64, limit, offset int) (*ms.ContactList, error) {
	follows, totoal, err := s.f.ListFollows(s.db, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	res := &ms.ContactList{
		Total: totoal,
	}
	for _, f := range follows {
		res.Contacts = append(res.Contacts, ms.ContactItem{
			UserId:    f.User.ID,
			Username:  f.User.Username,
			Nickname:  f.User.Nickname,
			Avatar:    f.User.Avatar,
			CreatedOn: f.User.CreatedOn,
		})
	}
	return res, nil
}

func (s *followingManageSrv) ListFollowings(userId int64, limit, offset int) (*ms.ContactList, error) {
	followingIds, totoal, err := s.f.ListFollowingIds(s.db, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	followings, err := s.u.ListUserInfoById(s.db, followingIds)
	if err != nil {
		return nil, err
	}
	res := &ms.ContactList{
		Total: totoal,
	}
	for _, user := range followings {
		res.Contacts = append(res.Contacts, ms.ContactItem{
			UserId:    user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			CreatedOn: user.CreatedOn,
		})
	}
	return res, nil
}

func (s *followingManageSrv) GetFollowCount(userId int64) (int64, int64, error) {
	return s.f.FollowCount(s.db, userId)
}

func (s *followingManageSrv) IsFollow(userId int64, followId int64) bool {
	if _, err := s.f.GetFollowing(s.db, userId, followId); err == nil {
		return true
	}
	return false
}
