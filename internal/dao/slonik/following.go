// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/slonik/sqlc/auto/pgc"
)

var (
	_ core.FollowingManageService = (*followingManageSrv)(nil)
)

type followingManageSrv struct {
	*pgcSrv
}

func (s *followingManageSrv) FollowUser(userId int64, followId int64) error {
	ctx := context.Background()
	if exist, err := s.q.ExistFollowing(ctx, &pgc.ExistFollowingParams{
		UserID:   userId,
		FollowID: followId,
	}); err == nil && exist {
		return nil
	}
	return s.q.CreateFollowing(ctx, &pgc.CreateFollowingParams{
		UserID:    userId,
		FollowID:  followId,
		CreatedOn: time.Now().Unix(),
	})
}

func (s *followingManageSrv) UnfollowUser(userId int64, followId int64) error {
	return s.q.DeleteFollowing(context.Background(), &pgc.DeleteFollowingParams{
		DeletedOn: time.Now().Unix(),
		UserID:    userId,
		FollowID:  followId,
	})
}

func (s *followingManageSrv) ListFollows(userId int64, limit, offset int) (res *ms.ContactList, err error) {
	var follows []*pgc.ListFollowsRow
	res = &ms.ContactList{}
	ctx := context.Background()
	if res.Total, err = s.q.CountFollows(ctx, userId); err != nil || res.Total == 0 {
		return
	}
	if follows, err = s.q.ListFollows(ctx, &pgc.ListFollowsParams{
		UserID: userId,
		Limit:  int32(limit),
		Offset: int32(offset),
	}); err != nil {
		return
	}
	for _, f := range follows {
		res.Contacts = append(res.Contacts, ms.ContactItem{
			UserId:    f.UserID,
			Username:  f.Username,
			Nickname:  f.Nickname,
			Avatar:    f.Avatar,
			CreatedOn: f.CreatedOn,
		})
	}
	return
}

func (s *followingManageSrv) ListFollowings(userId int64, limit, offset int) (res *ms.ContactList, err error) {
	var followings []*pgc.ListFollowingsRow
	res = &ms.ContactList{}
	ctx := context.Background()
	if res.Total, err = s.q.CountFollowings(ctx, userId); err != nil || res.Total == 0 {
		return
	}
	if followings, err = s.q.ListFollowings(ctx, &pgc.ListFollowingsParams{
		FollowID: userId,
		Limit:    int32(limit),
		Offset:   int32(offset),
	}); err != nil {
		return
	}
	for _, f := range followings {
		res.Contacts = append(res.Contacts, ms.ContactItem{
			UserId:    f.UserID,
			Username:  f.Username,
			Nickname:  f.Nickname,
			Avatar:    f.Avatar,
			CreatedOn: f.CreatedOn,
		})
	}
	return
}

func (s *followingManageSrv) GetFollowCount(userId int64) (follows int64, followings int64, err error) {
	ctx := context.Background()
	if follows, err = s.q.CountFollows(ctx, userId); err == nil {
		followings, err = s.q.CountFollowings(ctx, userId)
	}
	return
}

func (s *followingManageSrv) IsFollow(userId int64, followId int64) bool {
	if exist, err := s.q.ExistFollowing(context.TODO(), &pgc.ExistFollowingParams{
		UserID:   userId,
		FollowID: followId,
	}); err == nil {
		return exist
	}
	return false
}

func newFollowingManageService(db *pgx.Conn) core.FollowingManageService {
	return &followingManageSrv{
		pgcSrv: newPgcSrv(db),
	}
}
