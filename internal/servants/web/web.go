package web

import (
	api "github.com/rocboss/paopao-ce/internal/mirc/auto/api/v1"
)

func NewWebCoreSrv() api.WebCore {
	return &webCoreSrv{}
}

func NewWebCoreBinding() api.WebCoreBinding {
	return &webCoreBinding{}
}

func NewWebCoreRender() api.WebCoreRender {
	return &webCoreRender{}
}
