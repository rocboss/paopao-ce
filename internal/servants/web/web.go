// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
)

// RouteWeb register web route
func RouteWeb(e *gin.Engine) {
	api.RegisterWebCoreServant(e, newWebCoreSrv(), newWebCoreBinding(), newWebCoreRender())
	api.RegisterWebAdminServant(e, newWebAdminSrv(), newWebAdminBinding(), newWebAdminRender())
}
