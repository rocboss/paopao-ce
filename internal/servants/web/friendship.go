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
	_ api.Friendship        = (*friendshipSrv)(nil)
	_ api.FriendshipBinding = (*friendshipBinding)(nil)
	_ api.FriendshipRender  = (*friendshipRender)(nil)
)

type friendshipSrv struct {
	base.BaseServant
	api.UnimplementedFriendshipServant
}

type friendshipBinding struct {
	base.BaseBinding
	*api.UnimplementedFriendshipBinding
}

type friendshipRender struct {
	base.BaseRender
	*api.UnimplementedFriendshipRender
}

func (s *friendshipSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func newFriendshipSrv() api.Friendship {
	return &friendshipSrv{}
}

func newFriendshipBinding() api.FriendshipBinding {
	return &friendshipBinding{
		UnimplementedFriendshipBinding: &api.UnimplementedFriendshipBinding{
			BindAny: base.BindAny,
		},
	}
}

func newFriendshipRender() api.FriendshipRender {
	return &friendshipRender{
		UnimplementedFriendshipRender: &api.UnimplementedFriendshipRender{
			RenderAny: base.RenderAny,
		},
	}
}
