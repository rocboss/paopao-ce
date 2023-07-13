// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
)

var (
	_ core.UserManageService = (*userManageSrv)(nil)
)

type userManageSrv struct {
	*sqlxSrv
	q *cc.UserManage
}

func (s *userManageSrv) GetUserByID(id int64) (*ms.User, error) {
	res := &ms.User{}
	err := s.q.GetUserById.Get(res, id)
	return res, err
}

func (s *userManageSrv) GetUserByUsername(username string) (*ms.User, error) {
	res := &ms.User{}
	err := s.q.GetUserByUsername.Get(res, username)
	return res, err
}

func (s *userManageSrv) GetUserByPhone(phone string) (*ms.User, error) {
	res := &ms.User{}
	err := s.q.GetUserByPhone.Get(res, phone)
	return res, err
}

func (s *userManageSrv) GetUsersByIDs(ids []int64) ([]*ms.User, error) {
	res := []*ms.User{}
	err := s.inSelect(s.db, &res, s.q.GetUsersByIds, ids)
	return res, err
}

func (s *userManageSrv) GetUsersByKeyword(keyword string) ([]*ms.User, error) {
	res := []*ms.User{}
	err := s.q.GetUsersByKeyword.Select(&res, keyword)
	return res, err
}

func (s *userManageSrv) CreateUser(r *ms.User) (*ms.User, error) {
	res, err := s.q.AddUser.Exec(r.Username, r.Nickname, r.Password, r.Salt, r.Avatar)
	if err != nil {
		return nil, err
	}
	r.Model = &dbr.Model{}
	r.ID, err = res.LastInsertId()
	return r, err
}

func (s *userManageSrv) UpdateUser(r *ms.User) error {
	_, err := s.q.UpdateUser.Exec(r.Nickname, r.Password, r.Salt, r.Phone, r.Status)
	return err
}

func newUserManageService(db *sqlx.DB) core.UserManageService {
	return &userManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildUserManage),
	}
}
