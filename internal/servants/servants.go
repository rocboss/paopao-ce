package servants

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/servants/localoss"
	"github.com/rocboss/paopao-ce/internal/servants/web"
)

// RegisterWebServants register all the servants to gin.Engine
func RegisterWebServants(e *gin.Engine) {
	// 按需注册 docs、静态资源
	{
		registerDocs(e)
		registerStatick(e)
	}

	web.RouteWeb(e)
	localoss.RouteLocalOSS(e)
}
