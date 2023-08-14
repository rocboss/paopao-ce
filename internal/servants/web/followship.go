// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Followship = (*followshipSrv)(nil)
)

type followshipSrv struct {
	api.UnimplementedFollowshipServant
	*base.DaoServant
}

func (s *followshipSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (s *followshipSrv) ListFollowings(r *web.ListFollowingsReq) (*web.ListFollowingsResp, mir.Error) {
	he, err := s.Ds.GetUserByUsername(r.Username)
	if err != nil {
		logrus.Errorf("Ds.GetUserByUsername err: %s", err)
		return nil, web.ErrNoExistUsername
	}
	res, err := s.Ds.ListFollowings(he.ID, r.PageSize, r.Page-1)
	if err != nil {
		logrus.Errorf("Ds.ListFollowings err: %s", err)
		return nil, web.ErrListFollowingsFailed
	}
	resp := base.PageRespFrom(res.Contacts, r.Page, r.PageSize, res.Total)
	return (*web.ListFollowingsResp)(resp), nil
}

func (s *followshipSrv) ListFollows(r *web.ListFollowsReq) (*web.ListFollowsResp, mir.Error) {
	he, err := s.Ds.GetUserByUsername(r.Username)
	if err != nil {
		logrus.Errorf("Ds.GetUserByUsername err: %s", err)
		return nil, web.ErrNoExistUsername
	}
	res, err := s.Ds.ListFollows(he.ID, r.PageSize, r.Page-1)
	if err != nil {
		logrus.Errorf("Ds.ListFollows err: %s", err)
		return nil, web.ErrListFollowsFailed
	}
	resp := base.PageRespFrom(res.Contacts, r.Page, r.PageSize, res.Total)
	return (*web.ListFollowsResp)(resp), nil
}

func (s *followshipSrv) UnfollowUser(r *web.UnfollowUserReq) mir.Error {
	if err := s.Ds.UnfollowUser(r.User.ID, r.UserId); err != nil {
		logrus.Errorf("Ds.UnfollowUser err: %s userId: %d followId: %d", err, r.User.ID, r.UserId)
		return web.ErrUnfollowUserFailed
	}
	return nil
}

func (s *followshipSrv) FollowUser(r *web.FollowUserReq) mir.Error {
	if err := s.Ds.FollowUser(r.User.ID, r.UserId); err != nil {
		logrus.Errorf("Ds.FollowUser err: %s userId: %d followId: %d", err, r.User.ID, r.UserId)
		return web.ErrUnfollowUserFailed
	}
	return nil
}

func newFollowshipSrv(s *base.DaoServant) api.Followship {
	return &followshipSrv{
		DaoServant: s,
	}
}
