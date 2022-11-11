package web

import (
	api "github.com/rocboss/paopao-ce/internal/mirc/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/servants/core"
)

var (
	_ api.WebCore        = (*webCoreSrv)(nil)
	_ api.WebCoreBinding = (*webCoreBinding)(nil)
	_ api.WebCoreRender  = (*webCoreRender)(nil)
)

type webCoreSrv struct {
	core.BaseServant
	api.UnimplementedWebCoreServant
}

type webCoreBinding struct {
	core.BaseBinding
	*api.UnimplementedWebCoreBinding
}

type webCoreRender struct {
	core.BaseRender
	*api.UnimplementedWebCoreRender
}

func newWebCoreSrv() api.WebCore {
	return &webCoreSrv{}
}

func newWebCoreBinding() api.WebCoreBinding {
	return &webCoreBinding{
		UnimplementedWebCoreBinding: &api.UnimplementedWebCoreBinding{
			BindAny: core.BindAny,
		},
	}
}

func newWebCoreRender() api.WebCoreRender {
	return &webCoreRender{
		UnimplementedWebCoreRender: &api.UnimplementedWebCoreRender{
			RenderAny: core.RenderAny,
		},
	}
}
