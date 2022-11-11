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
	api.UnimplementedWebCoreBinding
}

type webCoreRender struct {
	core.BaseRender
	api.UnimplementedWebCoreRender
}
