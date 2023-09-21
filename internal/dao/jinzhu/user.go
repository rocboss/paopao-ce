// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
	"strings"
	"time"
)

var (
	_ core.UserManageService = (*userManageSrv)(nil)
)

type userManageSrv struct {
	db *gorm.DB
}

func newUserManageService(db *gorm.DB) core.UserManageService {
	return &userManageSrv{
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

// GetRegisterUserCount 获取注册用户数量
func (s *userManageSrv) GetRegisterUserCount() (res int64, err error) {
	err = s.db.Model(&dbr.User{}).Count(&res).Error
	return
}

// AddUserSubscribe 增加用户订阅
func (s *userManageSrv) AddUserSubscribe(user *ms.User, days int, reason string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if user.Balance == 0 || user.Balance <= time.Now().Unix() {
			//更新当前用户的balance时间戳，增加days天
			user.Balance = time.Now().Unix() + int64(days)*24*60*60
		} else {
			//更新当前用户的balance时间戳，增加14天
			user.Balance = user.Balance + int64(days)*24*60*60
		}

		err := user.Update(s.db)
		if err != nil {
			return err
		}

		// 新增账单
		if err := tx.Create(&dbr.WalletStatement{
			UserID:          user.ID,
			ChangeAmount:    int64(days * 100),
			BalanceSnapshot: user.Balance,
			Reason:          reason,
		}).Error; err != nil {
			return err
		}

		//返回nil，事务会被提交
		return nil
	})
}
