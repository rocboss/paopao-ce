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
	_ api.WebAlipayPub         = (*webAlipayPubSrv)(nil)
	_ api.WebAlipayPubBinding  = (*webAlipayPubBinding)(nil)
	_ api.WebAlipayPubRender   = (*webAlipayPubRender)(nil)
	_ api.WebAlipayPriv        = (*webAlipayPrivSrv)(nil)
	_ api.WebAlipayPrivBinding = (*webAlipayPrivBinding)(nil)
	_ api.WebAlipayPrivRender  = (*webAlipayPrivRender)(nil)
)

type webAlipayPubSrv struct {
	base.BaseServant
	api.UnimplementedWebAlipayPubServant
}

type webAlipayPubBinding struct {
	base.BaseBinding
	*api.UnimplementedWebAlipayPubBinding
}

type webAlipayPubRender struct {
	base.BaseRender
	*api.UnimplementedWebAlipayPubRender
}

type webAlipayPrivSrv struct {
	base.BaseServant
	api.UnimplementedWebAlipayPrivServant
}

type webAlipayPrivBinding struct {
	base.BaseBinding
	*api.UnimplementedWebAlipayPrivBinding
}

type webAlipayPrivRender struct {
	base.BaseRender
	*api.UnimplementedWebAlipayPrivRender
}

func (s *webAlipayPrivSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func newWebAlipayPubSrv() api.WebAlipayPub {
	return &webAlipayPubSrv{}
}

func newWebAlipayPubBinding() api.WebAlipayPubBinding {
	return &webAlipayPubBinding{
		UnimplementedWebAlipayPubBinding: &api.UnimplementedWebAlipayPubBinding{
			BindAny: base.BindAny,
		},
	}
}

func newWebAlipayPubRender() api.WebAlipayPubRender {
	return &webAlipayPubRender{
		UnimplementedWebAlipayPubRender: &api.UnimplementedWebAlipayPubRender{
			RenderAny: base.RenderAny,
		},
	}
}

func newWebAlipayPrivSrv() api.WebAlipayPriv {
	return &webAlipayPrivSrv{}
}

func newWebAlipayPrivBinding() api.WebAlipayPrivBinding {
	return &webAlipayPrivBinding{
		UnimplementedWebAlipayPrivBinding: &api.UnimplementedWebAlipayPrivBinding{
			BindAny: base.BindAny,
		},
	}
}

func newWebAlipayPrivRender() api.WebAlipayPrivRender {
	return &webAlipayPrivRender{
		UnimplementedWebAlipayPrivRender: &api.UnimplementedWebAlipayPrivRender{
			RenderAny: base.RenderAny,
		},
	}
}
