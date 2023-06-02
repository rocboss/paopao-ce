// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.UserManageService = (*userManageSrv)(nil)
)

type userManageSrv struct {
	*sqlxSrv
	q *cc.UserManage
}

func (s *userManageSrv) GetUserByID(id int64) (*ms.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) GetUserByUsername(username string) (*ms.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) GetUserByPhone(phone string) (*ms.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) GetUsersByIDs(ids []int64) ([]*ms.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) GetUsersByKeyword(keyword string) ([]*ms.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) CreateUser(user *ms.User) (*ms.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) UpdateUser(user *ms.User) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func newUserManageService(db *sqlx.DB) core.UserManageService {
	return &userManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildUserManage),
	}
}
