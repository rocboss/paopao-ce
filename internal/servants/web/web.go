package web

import (
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/internal/mirc/auto/api/v1"
)

// RouteWeb register web route
func RouteWeb(e *gin.Engine) {
	api.RegisterWebCoreServant(e, newWebCoreSrv(), newWebCoreBinding(), newWebCoreRender())
}
