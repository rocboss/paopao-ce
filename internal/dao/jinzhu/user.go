// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"strings"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.UserManageService = (*userManageSrv)(nil)
)

type userManageSrv struct {
	db *gorm.DB
}

type userRelationSrv struct {
	db *gorm.DB
}

func newUserManageService(db *gorm.DB) core.UserManageService {
	return &userManageSrv{
		db: db,
	}
}

func newUserRelationService(db *gorm.DB) core.UserRelationService {
	return &userRelationSrv{
		db: db,
	}
}

func (s *userManageSrv) GetUserByID(id int64) (*ms.User, error) {
	user := &dbr.User{
		Model: &dbr.Model{
			ID: id,
		},
	}
	return user.Get(s.db)
}

func (s *userManageSrv) GetUserByUsername(username string) (*ms.User, error) {
	user := &dbr.User{
		Username: username,
	}
	return user.Get(s.db)
}

func (s *userManageSrv) GetUserByPhone(phone string) (*ms.User, error) {
	user := &dbr.User{
		Phone: phone,
	}
	return user.Get(s.db)
}

func (s *userManageSrv) GetUsersByIDs(ids []int64) ([]*ms.User, error) {
	user := &dbr.User{}
	return user.List(s.db, &dbr.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}

func (s *userManageSrv) GetUsersByKeyword(keyword string) ([]*ms.User, error) {
	user := &dbr.User{}
	keyword = strings.Trim(keyword, " ") + "%"
	if keyword == "%" {
		return user.List(s.db, &dbr.ConditionsT{
			"ORDER": "id ASC",
		}, 0, 6)
	} else {
		return user.List(s.db, &dbr.ConditionsT{
			"username LIKE ?": keyword,
		}, 0, 6)
	}
}

func (s *userManageSrv) CreateUser(user *dbr.User) (*ms.User, error) {
	return user.Create(s.db)
}

func (s *userManageSrv) UpdateUser(user *ms.User) error {
	return user.Update(s.db)
}

func (s *userManageSrv) GetRegisterUserCount() (res int64, err error) {
	err = s.db.Model(&dbr.User{}).Count(&res).Error
	return
}

func (s *userRelationSrv) MyFriendIds(userId int64) (res []int64, err error) {
	err = s.db.Table(_contact_).Where("user_id=?", userId).Select("friend_id").Find(&res).Error
	return
}

func (s *userRelationSrv) MyFollowIds(userId int64) (res []int64, err error) {
	err = s.db.Table(_following_).Where("user_id=?", userId).Select("follow_id").Find(&res).Error
	return
}

func (s *userRelationSrv) IsMyFriend(userId int64, friendIds []int64) (map[int64]bool, error) {
	size := len(friendIds)
	res := make(map[int64]bool, size)
	if size == 0 {
		return res, nil
	}
	myFriendIds, err := s.MyFriendIds(userId)
	if err != nil {
		return nil, err
	}
	for _, friendId := range friendIds {
		res[friendId] = false
		for _, myFriendId := range myFriendIds {
			if friendId == myFriendId {
				res[friendId] = true
				break
			}
		}
	}
	return res, nil
}

func (s *userRelationSrv) IsMyFollow(userId int64, followIds []int64) (map[int64]bool, error) {
	size := len(followIds)
	res := make(map[int64]bool, size)
	if size == 0 {
		return res, nil
	}
	myFollowIds, err := s.MyFollowIds(userId)
	if err != nil {
		return nil, err
	}
	for _, followId := range followIds {
		res[followId] = false
		for _, myFollowId := range myFollowIds {
			if followId == myFollowId {
				res[followId] = true
				break
			}
		}
	}
	return res, nil
}
