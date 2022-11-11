package servants

import (
	"github.com/gin-gonic/gin"

	api "github.com/rocboss/paopao-ce/internal/mirc/auto/api/v1"
)

// RegisterServants register all the servants to gin.Engine
func RegisterServants(e *gin.Engine) {
	api.RegisterWebCoreServant(e, newWebCoreSrv(), newWebCoreBinding(), newWebCoreRender())
}
