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
	_ api.AlipayPub         = (*alipayPubSrv)(nil)
	_ api.AlipayPubBinding  = (*alipayPubBinding)(nil)
	_ api.AlipayPubRender   = (*alipayPubRender)(nil)
	_ api.AlipayPriv        = (*alipayPrivSrv)(nil)
	_ api.AlipayPrivBinding = (*alipayPrivBinding)(nil)
	_ api.AlipayPrivRender  = (*alipayPrivRender)(nil)
)

type alipayPubSrv struct {
	base.BaseServant
	api.UnimplementedAlipayPubServant
}

type alipayPubBinding struct {
	base.BaseBinding
	*api.UnimplementedAlipayPubBinding
}

type alipayPubRender struct {
	base.BaseRender
	*api.UnimplementedAlipayPubRender
}

type alipayPrivSrv struct {
	base.BaseServant
	api.UnimplementedAlipayPrivServant
}

type alipayPrivBinding struct {
	base.BaseBinding
	*api.UnimplementedAlipayPrivBinding
}

type alipayPrivRender struct {
	base.BaseRender
	*api.UnimplementedAlipayPrivRender
}

func (s *alipayPrivSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func newAlipayPubSrv() api.AlipayPub {
	return &alipayPubSrv{}
}

func newAlipayPubBinding() api.AlipayPubBinding {
	return &alipayPubBinding{
		UnimplementedAlipayPubBinding: &api.UnimplementedAlipayPubBinding{
			BindAny: base.BindAny,
		},
	}
}

func newAlipayPubRender() api.AlipayPubRender {
	return &alipayPubRender{
		UnimplementedAlipayPubRender: &api.UnimplementedAlipayPubRender{
			RenderAny: base.RenderAny,
		},
	}
}

func newAlipayPrivSrv() api.AlipayPriv {
	return &alipayPrivSrv{}
}

func newAlipayPrivBinding() api.AlipayPrivBinding {
	return &alipayPrivBinding{
		UnimplementedAlipayPrivBinding: &api.UnimplementedAlipayPrivBinding{
			BindAny: base.BindAny,
		},
	}
}

func newAlipayPrivRender() api.AlipayPrivRender {
	return &alipayPrivRender{
		UnimplementedAlipayPrivRender: &api.UnimplementedAlipayPrivRender{
			RenderAny: base.RenderAny,
		},
	}
}
