// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"strings"
	"time"

	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
)

var (
	_ core.UserManageService = (*userManageSrv)(nil)
)

type userManageSrv struct {
	*sqlxSrv
	q *cc.UserManage
}

func (s *userManageSrv) GetUserByID(id int64) (res *ms.User, err error) {
	err = s.q.GetUserById.Get(res, id)
	return
}

func (s *userManageSrv) GetUserByUsername(username string) (res *ms.User, err error) {
	err = s.q.GetUserByUsername.Get(res, username)
	return
}

func (s *userManageSrv) GetUserByPhone(phone string) (res *ms.User, err error) {
	err = s.q.GetUserByPhone.Get(res, phone)
	return
}

func (s *userManageSrv) GetUsersByIDs(ids []int64) (res []*ms.User, err error) {
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

func newUserManageService(db *sqlx.DB) core.UserManageService {
	return &userManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildUserManage),
	}
}
