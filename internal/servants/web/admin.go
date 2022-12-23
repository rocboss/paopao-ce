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
	_ api.WebAdmin        = (*webAdminSrv)(nil)
	_ api.WebAdminBinding = (*webAdminBinding)(nil)
	_ api.WebAdminRender  = (*webAdminRender)(nil)
)

type webAdminSrv struct {
	base.BaseServant
	api.UnimplementedWebAdminServant
}

type webAdminBinding struct {
	base.BaseBinding
	*api.UnimplementedWebAdminBinding
}

type webAdminRender struct {
	base.BaseRender
	*api.UnimplementedWebAdminRender
}

func (s *webAdminSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT(), chain.Admin()}
}

func newWebAdminSrv() api.WebAdmin {
	return &webAdminSrv{}
}

func newWebAdminBinding() api.WebAdminBinding {
	return &webAdminBinding{
		UnimplementedWebAdminBinding: &api.UnimplementedWebAdminBinding{
			BindAny: base.BindAny,
		},
	}
}

func newWebAdminRender() api.WebAdminRender {
	return &webAdminRender{
		UnimplementedWebAdminRender: &api.UnimplementedWebAdminRender{
			RenderAny: base.RenderAny,
		},
	}
}
