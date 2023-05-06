// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package chain

import (
	"github.com/alimy/cfg"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/app"
)

func Priv() gin.HandlerFunc {
	if cfg.If("PhoneBind") {
		return func(c *gin.Context) {
			if u, exist := c.Get("USER"); exist {
				if user, ok := u.(*core.User); ok {
					if user.Status == core.UserStatusNormal {
						if user.Phone == "" {
							response := app.NewResponse(c)
							response.ToErrorResponse(_errAccountNoPhoneBind)
							c.Abort()
							return
						}
						c.Next()
						return
					}
				}
			}
			response := app.NewResponse(c)
			response.ToErrorResponse(_errUserHasBeenBanned)
			c.Abort()
		}
	} else {
		return func(c *gin.Context) {
			if u, exist := c.Get("USER"); exist {
				if user, ok := u.(*core.User); ok && user.Status == core.UserStatusNormal {
					c.Next()
					return
				}
			}
			response := app.NewResponse(c)
			response.ToErrorResponse(_errUserHasBeenBanned)
			c.Abort()
		}
	}
}
