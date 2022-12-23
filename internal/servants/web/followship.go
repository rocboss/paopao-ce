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
	_ api.Followship        = (*followshipSrv)(nil)
	_ api.FollowshipBinding = (*followshipBinding)(nil)
	_ api.FollowshipRender  = (*followshipRender)(nil)
)

type followshipSrv struct {
	base.BaseServant
	api.UnimplementedFollowshipServant
}

type followshipBinding struct {
	base.BaseBinding
	*api.UnimplementedFollowshipBinding
}

type followshipRender struct {
	base.BaseRender
	*api.UnimplementedFollowshipRender
}

func (s *followshipSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func newFollowshipSrv() api.Followship {
	return &followshipSrv{}
}

func newFollowshipBinding() api.FollowshipBinding {
	return &followshipBinding{
		UnimplementedFollowshipBinding: &api.UnimplementedFollowshipBinding{
			BindAny: base.BindAny,
		},
	}
}

func newFollowshipRender() api.FollowshipRender {
	return &followshipRender{
		UnimplementedFollowshipRender: &api.UnimplementedFollowshipRender{
			RenderAny: base.RenderAny,
		},
	}
}
