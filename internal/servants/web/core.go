// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

var (
	_ api.WebCore        = (*webCoreSrv)(nil)
	_ api.WebCoreBinding = (*webCoreBinding)(nil)
	_ api.WebCoreRender  = (*webCoreRender)(nil)
)

type webCoreSrv struct {
	base.BaseServant
	api.UnimplementedWebCoreServant
}

type webCoreBinding struct {
	base.BaseBinding
	*api.UnimplementedWebCoreBinding
}

type webCoreRender struct {
	base.BaseRender
	*api.UnimplementedWebCoreRender
}

func newWebCoreSrv() api.WebCore {
	return &webCoreSrv{}
}

func newWebCoreBinding() api.WebCoreBinding {
	return &webCoreBinding{
		UnimplementedWebCoreBinding: &api.UnimplementedWebCoreBinding{
			BindAny: base.BindAny,
		},
	}
}

func newWebCoreRender() api.WebCoreRender {
	return &webCoreRender{
		UnimplementedWebCoreRender: &api.UnimplementedWebCoreRender{
			RenderAny: base.RenderAny,
		},
	}
}
