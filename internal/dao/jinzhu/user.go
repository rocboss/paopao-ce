// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"strings"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.UserManageService = (*userManageServant)(nil)
)

type userManageServant struct {
	db *gorm.DB
}

func newUserManageService(db *gorm.DB) core.UserManageService {
	return &userManageServant{
		db: db,
	}
}

func (s *userManageServant) GetUserByID(id int64) (*core.User, error) {
	user := &dbr.User{
		Model: &dbr.Model{
			ID: id,
		},
	}
	return user.Get(s.db)
}

func (s *userManageServant) GetUserByUsername(username string) (*core.User, error) {
	user := &dbr.User{
		Username: username,
	}
	return user.Get(s.db)
}

func (s *userManageServant) GetUserByPhone(phone string) (*core.User, error) {
	user := &dbr.User{
		Phone: phone,
	}
	return user.Get(s.db)
}

func (s *userManageServant) GetUsersByIDs(ids []int64) ([]*core.User, error) {
	user := &dbr.User{}
	return user.List(s.db, &dbr.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}

func (s *userManageServant) GetUsersByKeyword(keyword string) ([]*core.User, error) {
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

func (s *userManageServant) GetTagsByKeyword(keyword string) ([]*core.Tag, error) {
	tag := &dbr.Tag{}
	keyword = "%" + strings.Trim(keyword, " ") + "%"
	if keyword == "%%" {
		return tag.List(s.db, &dbr.ConditionsT{
			"ORDER": "quote_num DESC",
		}, 0, 6)
	} else {
		return tag.List(s.db, &dbr.ConditionsT{
			"tag LIKE ?": keyword,
			"ORDER":      "quote_num DESC",
		}, 0, 6)
	}
}

func (s *userManageServant) CreateUser(user *dbr.User) (*core.User, error) {
	return user.Create(s.db)
}

func (s *userManageServant) UpdateUser(user *core.User) error {
	return user.Update(s.db)
}
