package servants

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/servants/docs"
	"github.com/rocboss/paopao-ce/internal/servants/localoss"
	"github.com/rocboss/paopao-ce/internal/servants/statick"
	"github.com/rocboss/paopao-ce/internal/servants/web"
	"github.com/rocboss/paopao-ce/pkg/cfg"
)

// RegisterWebServants register all the servants to gin.Engine
func RegisterWebServants(e *gin.Engine) {
	// 按需注册 docs、静态资源
	{
		docs.RegisterDocs(e)
		statick.RegisterStatick(e)
	}

	web.RouteWeb(e)
	{
		cfg.Be("LocalOSS", func() {
			localoss.RouteLocalOSS(e)
		})
	}
}
