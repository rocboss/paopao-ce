// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.UserManageService = (*userManageSrv)(nil)
)

type userManageSrv struct {
	*sqlxSrv
	stmtAddUser    *sqlx.Stmt
	stmtUpdateUser *sqlx.Stmt
	stmtGetUser    *sqlx.Stmt
}

func (s *userManageSrv) GetUserByID(id int64) (*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) GetUserByUsername(username string) (*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) GetUserByPhone(phone string) (*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) GetUsersByIDs(ids []int64) ([]*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) GetUsersByKeyword(keyword string) ([]*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) CreateUser(user *core.User) (*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageSrv) UpdateUser(user *core.User) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func newUserManageService(db *sqlx.DB) core.UserManageService {
	return &userManageSrv{
		sqlxSrv:        newSqlxSrv(db),
		stmtAddUser:    c(`SELECT * FROM @user WHERE username=?`),
		stmtUpdateUser: c(`SELECT * FROM @user WHERE username=?`),
		stmtGetUser:    c(`SELECT * FROM @user WHERE username=?`),
	}
}
