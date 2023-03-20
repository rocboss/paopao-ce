// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package chain

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/app"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, exist := c.Get("USER"); exist {
			if userModel, ok := user.(*core.User); ok {
				if userModel.Status == core.UserStatusNormal && userModel.IsAdmin {
					c.Next()
					return
				}
			}
		}

		response := app.NewResponse(c)
		response.ToErrorResponse(_errNoAdminPermission)
		c.Abort()
	}
}
