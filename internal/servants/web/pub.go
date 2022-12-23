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
	_ api.WebPub        = (*webPubSrv)(nil)
	_ api.WebPubBinding = (*webPubBinding)(nil)
	_ api.WebPubRender  = (*webPubRender)(nil)
)

type webPubSrv struct {
	base.BaseServant
	api.UnimplementedWebPubServant
}

type webPubBinding struct {
	base.BaseBinding
	*api.UnimplementedWebPubBinding
}

type webPubRender struct {
	base.BaseRender
	*api.UnimplementedWebPubRender
}

func (s *webPubSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtLoose()}
}

func newWebPubSrv() api.WebPub {
	return &webPubSrv{}
}

func newWebPubBinding() api.WebPubBinding {
	return &webPubBinding{
		UnimplementedWebPubBinding: &api.UnimplementedWebPubBinding{
			BindAny: base.BindAny,
		},
	}
}

func newWebPubRender() api.WebPubRender {
	return &webPubRender{
		UnimplementedWebPubRender: &api.UnimplementedWebPubRender{
			RenderAny: base.RenderAny,
		},
	}
}
