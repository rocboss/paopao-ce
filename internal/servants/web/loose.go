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
	_ api.Loose        = (*looseSrv)(nil)
	_ api.LooseBinding = (*looseBinding)(nil)
	_ api.LooseRender  = (*looseRender)(nil)
)

type looseSrv struct {
	base.BaseServant
	api.UnimplementedLooseServant
}

type looseBinding struct {
	base.BaseBinding
	*api.UnimplementedLooseBinding
}

type looseRender struct {
	base.BaseRender
	*api.UnimplementedLooseRender
}

func (s *looseSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtLoose()}
}

func newLooseSrv() api.Loose {
	return &looseSrv{}
}

func newLooseBinding() api.LooseBinding {
	return &looseBinding{
		UnimplementedLooseBinding: &api.UnimplementedLooseBinding{
			BindAny: base.BindAny,
		},
	}
}

func newLooseRender() api.LooseRender {
	return &looseRender{
		UnimplementedLooseRender: &api.UnimplementedLooseRender{
			RenderAny: base.RenderAny,
		},
	}
}
