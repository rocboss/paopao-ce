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
	_ api.WebFollowship        = (*webFollowshipSrv)(nil)
	_ api.WebFollowshipBinding = (*webFollowshipBinding)(nil)
	_ api.WebFollowshipRender  = (*webFollowshipRender)(nil)
)

type webFollowshipSrv struct {
	base.BaseServant
	api.UnimplementedWebFollowshipServant
}

type webFollowshipBinding struct {
	base.BaseBinding
	*api.UnimplementedWebFollowshipBinding
}

type webFollowshipRender struct {
	base.BaseRender
	*api.UnimplementedWebFollowshipRender
}

func (s *webFollowshipSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func newWebFollowshipSrv() api.WebFollowship {
	return &webFollowshipSrv{}
}

func newWebFollowshipBinding() api.WebFollowshipBinding {
	return &webFollowshipBinding{
		UnimplementedWebFollowshipBinding: &api.UnimplementedWebFollowshipBinding{
			BindAny: base.BindAny,
		},
	}
}

func newWebFollowshipRender() api.WebFollowshipRender {
	return &webFollowshipRender{
		UnimplementedWebFollowshipRender: &api.UnimplementedWebFollowshipRender{
			RenderAny: base.RenderAny,
		},
	}
}
