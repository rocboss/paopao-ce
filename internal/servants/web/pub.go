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
	_ api.Pub        = (*pubSrv)(nil)
	_ api.PubBinding = (*pubBinding)(nil)
	_ api.PubRender  = (*pubRender)(nil)
)

type pubSrv struct {
	base.BaseServant
	api.UnimplementedPubServant
}

type pubBinding struct {
	base.BaseBinding
	*api.UnimplementedPubBinding
}

type pubRender struct {
	base.BaseRender
	*api.UnimplementedPubRender
}

func (s *pubSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtLoose()}
}

func newPubSrv() api.Pub {
	return &pubSrv{}
}

func newPubBinding() api.PubBinding {
	return &pubBinding{
		UnimplementedPubBinding: &api.UnimplementedPubBinding{
			BindAny: base.BindAny,
		},
	}
}

func newPubRender() api.PubRender {
	return &pubRender{
		UnimplementedPubRender: &api.UnimplementedPubRender{
			RenderAny: base.RenderAny,
		},
	}
}
