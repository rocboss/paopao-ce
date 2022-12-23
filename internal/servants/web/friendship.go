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
	_ api.WebFriendship        = (*webFriendshipSrv)(nil)
	_ api.WebFriendshipBinding = (*webFriendshipBinding)(nil)
	_ api.WebFriendshipRender  = (*webFriendshipRender)(nil)
)

type webFriendshipSrv struct {
	base.BaseServant
	api.UnimplementedWebFriendshipServant
}

type webFriendshipBinding struct {
	base.BaseBinding
	*api.UnimplementedWebFriendshipBinding
}

type webFriendshipRender struct {
	base.BaseRender
	*api.UnimplementedWebFriendshipRender
}

func (s *webFriendshipSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func newWebFriendshipSrv() api.WebFriendship {
	return &webFriendshipSrv{}
}

func newWebFriendshipBinding() api.WebFriendshipBinding {
	return &webFriendshipBinding{
		UnimplementedWebFriendshipBinding: &api.UnimplementedWebFriendshipBinding{
			BindAny: base.BindAny,
		},
	}
}

func newWebFriendshipRender() api.WebFriendshipRender {
	return &webFriendshipRender{
		UnimplementedWebFriendshipRender: &api.UnimplementedWebFriendshipRender{
			RenderAny: base.RenderAny,
		},
	}
}
