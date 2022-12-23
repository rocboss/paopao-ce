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
	_ api.WebLoose        = (*webLooseSrv)(nil)
	_ api.WebLooseBinding = (*webLooseBinding)(nil)
	_ api.WebLooseRender  = (*webLooseRender)(nil)
)

type webLooseSrv struct {
	base.BaseServant
	api.UnimplementedWebLooseServant
}

type webLooseBinding struct {
	base.BaseBinding
	*api.UnimplementedWebLooseBinding
}

type webLooseRender struct {
	base.BaseRender
	*api.UnimplementedWebLooseRender
}

func (s *webLooseSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtLoose()}
}

func newWebLooseSrv() api.WebLoose {
	return &webLooseSrv{}
}

func newWebLooseBinding() api.WebLooseBinding {
	return &webLooseBinding{
		UnimplementedWebLooseBinding: &api.UnimplementedWebLooseBinding{
			BindAny: base.BindAny,
		},
	}
}

func newWebLooseRender() api.WebLooseRender {
	return &webLooseRender{
		UnimplementedWebLooseRender: &api.UnimplementedWebLooseRender{
			RenderAny: base.RenderAny,
		},
	}
}
