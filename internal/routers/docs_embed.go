//go:build docs
// +build docs

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/docs/openapi"
	"github.com/rocboss/paopao-ce/internal/conf"
)

// registerDocs register docs asset route
func registerDocs(e *gin.Engine) {
	if conf.CfgIf("Docs:OpenAPI") {
		e.StaticFS("/docs/openapi", openapi.NewFileSystem())
	}
}
