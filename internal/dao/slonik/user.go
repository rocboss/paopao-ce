// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	dbr "github.com/rocboss/paopao-ce/internal/dao/slonik/ce/postgres"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.UserManageService = (*userManageServant)(nil)
)

type userManageServant struct {
	db *pgx.Conn
	q  dbr.Querier
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

func newUserManageService(db *pgx.Conn) core.UserManageService {
	return &userManageServant{
		db: db,
		q:  dbr.New(db),
	}
}
