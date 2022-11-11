package servants

import (
	api "github.com/rocboss/paopao-ce/internal/mirc/auto/api/v1"
)

type webCoreSrv struct {
	baseServant
	api.UnimplementedWebCoreServant
}

type webCoreBinding struct {
	baseBinding
	api.UnimplementedWebCoreBinding
}

type webCoreRender struct {
	baseRender
	api.UnimplementedWebCoreRender
}

func newWebCoreSrv() api.WebCore {
	return &webCoreSrv{}
}

func newWebCoreBinding() api.WebCoreBinding {
	return &webCoreBinding{}
}

func newWebCoreRender() api.WebCoreRender {
	return &webCoreRender{}
}
