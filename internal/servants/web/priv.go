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
	_ api.WebPriv        = (*webPrivSrv)(nil)
	_ api.WebPrivBinding = (*webPrivBinding)(nil)
	_ api.WebPrivRender  = (*webPrivRender)(nil)
)

type webPrivSrv struct {
	base.BaseServant
	api.UnimplementedWebPrivServant
}

type webPrivBinding struct {
	base.BaseBinding
	*api.UnimplementedWebPrivBinding
}

type webPrivRender struct {
	base.BaseRender
	*api.UnimplementedWebPrivRender
}

func (s *webPrivSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT(), chain.Priv()}
}

func newWebPrivSrv() api.WebPriv {
	return &webPrivSrv{}
}

func newWebPrivBinding() api.WebPrivBinding {
	return &webPrivBinding{
		UnimplementedWebPrivBinding: &api.UnimplementedWebPrivBinding{
			BindAny: base.BindAny,
		},
	}
}

func newWebPrivRender() api.WebPrivRender {
	return &webPrivRender{
		UnimplementedWebPrivRender: &api.UnimplementedWebPrivRender{
			RenderAny: base.RenderAny,
		},
	}
}
