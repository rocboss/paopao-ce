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
	_ api.Priv        = (*privSrv)(nil)
	_ api.PrivBinding = (*privBinding)(nil)
	_ api.PrivRender  = (*privRender)(nil)
)

type privSrv struct {
	base.BaseServant
	api.UnimplementedPrivServant
}

type privBinding struct {
	base.BaseBinding
	*api.UnimplementedPrivBinding
}

type privRender struct {
	base.BaseRender
	*api.UnimplementedPrivRender
}

func (s *privSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT(), chain.Priv()}
}

func newPrivSrv() api.Priv {
	return &privSrv{}
}

func newPrivBinding() api.PrivBinding {
	return &privBinding{
		UnimplementedPrivBinding: &api.UnimplementedPrivBinding{
			BindAny: base.BindAny,
		},
	}
}

func newPrivRender() api.PrivRender {
	return &privRender{
		UnimplementedPrivRender: &api.UnimplementedPrivRender{
			RenderAny: base.RenderAny,
		},
	}
}
