//go:build docs
// +build docs

package docs

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/docs/openapi"
	"github.com/rocboss/paopao-ce/pkg/cfg"
)

// RegisterDocs register docs asset route
func RegisterDocs(e *gin.Engine) {
	cfg.Be("Docs:OpenAPI", func() {
		e.StaticFS("/docs/openapi", openapi.NewFileSystem())
	})
}
