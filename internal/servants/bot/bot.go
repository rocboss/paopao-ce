// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package bot

import (
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/r/v1"
)

// RouteBot register Bot route
func RouteBot(e *gin.Engine) {
	api.RegisterUserServant(e, newUserSrv(), newUserBinding(), newUserRender())
}
