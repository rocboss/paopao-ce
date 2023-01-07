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
	_ core.UserManageService = (*userManageServant)(nil)
)

type userManageServant struct {
	*sqlxServant
	stmtAddUser    *sqlx.Stmt
	stmtUpdateUser *sqlx.Stmt
	stmtGetUser    *sqlx.Stmt
}

func (s *userManageServant) GetUserByID(id int64) (*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageServant) GetUserByUsername(username string) (*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageServant) GetUserByPhone(phone string) (*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageServant) GetUsersByIDs(ids []int64) ([]*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageServant) GetUsersByKeyword(keyword string) ([]*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageServant) GetTagsByKeyword(keyword string) ([]*core.Tag, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageServant) CreateUser(user *core.User) (*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *userManageServant) UpdateUser(user *core.User) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func newUserManageService(db *sqlx.DB) core.UserManageService {
	return &userManageServant{
		sqlxServant:    newSqlxServant(db),
		stmtAddUser:    c(`SELECT * FROM @person WHERE first_name=?`),
		stmtUpdateUser: c(`SELECT * FROM @person WHERE first_name=?`),
		stmtGetUser:    c(`SELECT * FROM @person WHERE first_name=?`),
	}
}
