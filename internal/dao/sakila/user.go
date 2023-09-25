// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"strings"
	"time"

	"github.com/alimy/tryst/cfg"
	"github.com/bitbus/sqlx"
	"github.com/bitbus/sqlx/db"
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
	q   *cc.UserManage
	ums core.UserMetricServantA
}

type userRelationSrv struct {
	*sqlxSrv
	q *cc.UserRelation
}

func (s *userManageSrv) GetUserByID(id int64) (*ms.User, error) {
	return db.Get[ms.User](s.q.GetUserById, id)
}

func (s *userManageSrv) GetUserByUsername(username string) (*ms.User, error) {
	return db.Get[ms.User](s.q.GetUserByUsername, username)
}

func (s *userManageSrv) GetUserByPhone(phone string) (*ms.User, error) {
	return db.Get[ms.User](s.q.GetUserByPhone, phone)
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
	if r.ID, err = res.LastInsertId(); err == nil {
		// 宽松处理错误
		s.ums.AddUserMetric(r.ID)
	}
	return r, err
}

func (s *userManageSrv) UpdateUser(r *ms.User) error {
	_, err := s.q.UpdateUser.Exec(r)
	return err
}

func (s *userManageSrv) GetRegisterUserCount() (res int64, err error) {
	err = s.q.GetRegisterUserCount.Get(&res)
	return
}

func (s *userRelationSrv) MyFriendIds(userId int64) (res []int64, err error) {
	err = s.q.MyFriendIds.Select(&res, userId)
	return
}

func (s *userRelationSrv) MyFollowIds(userId int64) (res []int64, err error) {
	err = s.q.MyFollowIds.Select(&res, userId)
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

func newUserManageService(db *sqlx.DB, userMetric core.UserMetricServantA) (s core.UserManageService) {
	ums := &userManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildUserManage),
		ums:     userMetric,
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

func newUserRelationService(db *sqlx.DB) core.UserRelationService {
	return &userRelationSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildUserRelation),
	}
}
