// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

var (
	_ api.Admin        = (*adminSrv)(nil)
	_ api.AdminBinding = (*adminBinding)(nil)
	_ api.AdminRender  = (*adminRender)(nil)
)

type adminSrv struct {
	api.UnimplementedAdminServant
	*base.DaoServant
}

type adminBinding struct {
	*api.UnimplementedAdminBinding
}

type adminRender struct {
	*api.UnimplementedAdminRender
}

func (s *adminSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT(), chain.Admin()}
}

func (s *adminSrv) ChangeUserStatus(req *web.ChangeUserStatusReq) mir.Error {
	user, err := s.Ds.GetUserByID(req.ID)
	if err != nil || user.Model == nil || user.ID <= 0 {
		return _errNoExistUsername
	}
	// 执行更新
	user.Status = req.Status
	if err := s.Ds.UpdateUser(user); err != nil {
		return xerror.ServerError
	}
	return nil
}

func newAdminSrv(s *base.DaoServant) api.Admin {
	return &adminSrv{
		DaoServant: s,
	}
}

func newAdminBinding() api.AdminBinding {
	return &adminBinding{
		UnimplementedAdminBinding: &api.UnimplementedAdminBinding{
			BindAny: base.NewBindAnyFn(),
		},
	}
}

func newAdminRender() api.AdminRender {
	return &adminRender{
		UnimplementedAdminRender: &api.UnimplementedAdminRender{
			RenderAny: base.RenderAny,
		},
	}
}
