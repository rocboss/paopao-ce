// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"strings"
	"time"

	"github.com/alimy/cfg"
	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pgc"
)

var (
	_ core.UserManageService = (*userManageSrv)(nil)
)

type userManageSrv struct {
	*sqlxSrv
	q *cc.UserManage
}

func (s *userManageSrv) GetUserByID(id int64) (res *ms.User, err error) {
	err = stmtGet(s.q.GetUserById, &res, id)
	return
}

func (s *userManageSrv) GetUserByUsername(username string) (res *ms.User, err error) {
	err = stmtGet(s.q.GetUserByUsername, &res, username)
	return
}

func (s *userManageSrv) GetUserByPhone(phone string) (res *ms.User, err error) {
	err = stmtGet(s.q.GetUserByPhone, &res, phone)
	return
}

func (s *userManageSrv) GetUsersByIDs(ids []int64) (res []*ms.User, err error) {
	if len(ids) == 0 {
		return nil, nil
	}
	err = s.db.InSelect(&res, s.q.GetUsersByIds, ids)
	return
}

func (s *userManageSrv) GetUsersByKeyword(keyword string) (res []*ms.User, err error) {
	keyword = strings.Trim(keyword, " ") + "%"
	if keyword == "%" {
		err = s.q.GetAnyUsers.Get(&res)
	} else {
		err = s.q.GetUsersByKeyword.Select(&res, keyword)
	}
	return
}

func (s *userManageSrv) CreateUser(r *ms.User) (*ms.User, error) {
	r.Model = &ms.Model{
		CreatedOn: time.Now().Unix(),
	}
	res, err := s.q.CreateUser.Exec(r)
	if err != nil {
		return nil, err
	}
	r.ID, err = res.LastInsertId()
	return r, err
}

func (s *userManageSrv) UpdateUser(r *ms.User) error {
	_, err := s.q.UpdateUser.Exec(r)
	return err
}

func newUserManageService(db *sqlx.DB) (s core.UserManageService) {
	ums := &userManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildUserManage),
	}
	s = ums
	if cfg.Any("PostgreSQL", "PgSQL", "Postgres") {
		s = &pgcUserManageSrv{
			userManageSrv: ums,
			p:             pgcBuild(db, pgc.BuildUserManage),
		}
	}
	return
}
