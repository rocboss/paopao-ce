// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
)

var (
	_ api.Admin        = (*adminSrv)(nil)
	_ api.AdminBinding = (*adminBinding)(nil)
	_ api.AdminRender  = (*adminRender)(nil)
)

type adminSrv struct {
	base.BaseServant
	api.UnimplementedAdminServant
}

type adminBinding struct {
	base.BaseBinding
	*api.UnimplementedAdminBinding
}

type adminRender struct {
	base.BaseRender
	*api.UnimplementedAdminRender
}

func (s *adminSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT(), chain.Admin()}
}

func newAdminSrv() api.Admin {
	return &adminSrv{}
}

func newAdminBinding() api.AdminBinding {
	return &adminBinding{
		UnimplementedAdminBinding: &api.UnimplementedAdminBinding{
			BindAny: base.BindAny,
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
