// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
)

var (
	_ core.FollowingManageService = (*followingManageSrv)(nil)
)

type followItem struct {
	UserId   int64
	Username string
	Nickname string
	Avatar   string
	CreateOn int64
}

type followingManageSrv struct {
	*sqlxSrv
	q *cc.FollowingManager
}

func (s *followingManageSrv) FollowUser(userId int64, followId int64) (err error) {
	exist := false
	if err = s.q.ExistFollowing.Get(&exist, userId, followId); err == nil && exist {
		return
	}
	_, err = s.q.CreateFollowing.Exec(userId, followId, time.Now().Unix())
	return
}

func (s *followingManageSrv) UnfollowUser(userId int64, followId int64) (err error) {
	_, err = s.q.DeleteFollowing.Exec(time.Now().Unix(), userId, followId)
	return
}

func (s *followingManageSrv) ListFollows(userId int64, limit, offset int) (res *ms.ContactList, err error) {
	follows := []followItem{}
	res = &ms.ContactList{}
	if err = s.q.ListFollows.Select(&follows, userId, limit, offset); err == nil {
		err = s.q.CountFollows.Get(&res.Total, userId)
	}
	for _, f := range follows {
		res.Contacts = append(res.Contacts, ms.ContactItem{
			UserId:    f.UserId,
			Username:  f.Username,
			Nickname:  f.Nickname,
			Avatar:    f.Avatar,
			CreatedOn: f.CreateOn,
		})
	}
	return
}

func (s *followingManageSrv) ListFollowings(userId int64, limit, offset int) (res *ms.ContactList, err error) {
	followings := []followItem{}
	res = &ms.ContactList{}
	if err = s.q.ListFollowings.Select(&followings, userId, limit, offset); err == nil {
		err = s.q.CountFollowings.Get(&res.Total, userId)
	}
	for _, f := range followings {
		res.Contacts = append(res.Contacts, ms.ContactItem{
			UserId:    f.UserId,
			Username:  f.Username,
			Nickname:  f.Nickname,
			Avatar:    f.Avatar,
			CreatedOn: f.CreateOn,
		})
	}
	return
}

func (s *followingManageSrv) GetFollowCount(userId int64) (follows int64, followings int64, err error) {
	if err = s.q.CountFollows.Get(&follows); err == nil {
		err = s.q.CountFollowings.Get(&followings)
	}
	return
}

func (s *followingManageSrv) IsFollow(userId int64, followId int64) (yn bool) {
	if err := s.q.ExistFollowing.Get(&yn, userId, followId); err == nil {
		return
	}
	return false
}

func newFollowingManageService(db *sqlx.DB) core.FollowingManageService {
	return &followingManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildFollowingManager),
	}
}
