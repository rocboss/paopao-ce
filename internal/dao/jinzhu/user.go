// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"fmt"
	"strings"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

type userManageSrv struct {
	db  *gorm.DB
	ums core.UserMetricServantA

	_userProfileJoins    string
	_userProfileWhere    string
	_userProfileColoumns []string
}

type userRelationSrv struct {
	db *gorm.DB
}

func newUserManageService(db *gorm.DB, ums core.UserMetricServantA) core.UserManageService {
	return &userManageSrv{
		db:                db,
		ums:               ums,
		_userProfileJoins: fmt.Sprintf("LEFT JOIN %s m ON %s.id=m.user_id", _userMetric_, _user_),
		_userProfileWhere: fmt.Sprintf("%s.username=? AND %s.is_del=0", _user_, _user_),
		_userProfileColoumns: []string{
			fmt.Sprintf("%s.id", _user_),
			fmt.Sprintf("%s.username", _user_),
			fmt.Sprintf("%s.nickname", _user_),
			fmt.Sprintf("%s.phone", _user_),
			fmt.Sprintf("%s.status", _user_),
			fmt.Sprintf("%s.avatar", _user_),
			fmt.Sprintf("%s.balance", _user_),
			fmt.Sprintf("%s.is_admin", _user_),
			fmt.Sprintf("%s.created_on", _user_),
			"m.tweets_count",
		},
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

func (s *userManageSrv) UserProfileByName(username string) (res *cs.UserProfile, err error) {
	err = s.db.Table(_user_).Joins(s._userProfileJoins).
		Where(s._userProfileWhere, username).
		Select(s._userProfileColoumns).
		First(&res).Error
	return
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

func (s *userManageSrv) CreateUser(user *dbr.User) (res *ms.User, err error) {
	if res, err = user.Create(s.db); err == nil {
		// 宽松处理错误
		s.ums.AddUserMetric(res.ID)
	}
	return
}

func (s *userManageSrv) UpdateUser(user *ms.User) error {
	return user.Update(s.db)
}

func (s *userManageSrv) GetRegisterUserCount() (res int64, err error) {
	err = s.db.Model(&dbr.User{}).Count(&res).Error
	return
}

func (s *userRelationSrv) MyFriendIds(userId int64) (res []int64, err error) {
	err = s.db.Table(_contact_).Where("user_id=? AND status=2 AND is_del=0", userId).Select("friend_id").Find(&res).Error
	return
}

func (s *userRelationSrv) MyFollowIds(userId int64) (res []int64, err error) {
	err = s.db.Table(_following_).Where("user_id=? AND is_del=0", userId).Select("follow_id").Find(&res).Error
	return
}

func (s *userRelationSrv) IsMyFriend(userId int64, friendIds ...int64) (map[int64]bool, error) {
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

func (s *userRelationSrv) IsMyFollow(userId int64, followIds ...int64) (map[int64]bool, error) {
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
