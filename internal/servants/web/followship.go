// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
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
	// TODO
	res := ms.ContactList{}
	resp := base.PageRespFrom(res.Contacts, r.Page, r.PageSize, res.Total)
	return (*web.ListFollowingsResp)(resp), nil
}

func (s *followshipSrv) ListFollows(r *web.ListFollowsReq) (*web.ListFollowsResp, mir.Error) {
	// TODO
	res := ms.ContactList{}
	resp := base.PageRespFrom(res.Contacts, r.Page, r.PageSize, res.Total)
	return (*web.ListFollowsResp)(resp), nil
}

func (s *followshipSrv) UnfollowUser(r *web.UnfollowUserReq) mir.Error {
	// TODO
	return web.ErrNotImplemented
}

func (s *followshipSrv) FollowUser(r *web.FollowUserReq) mir.Error {
	// TODO
	return web.ErrNotImplemented
}

func newFollowshipSrv(s *base.DaoServant) api.Followship {
	return &followshipSrv{
		DaoServant: s,
	}
}
