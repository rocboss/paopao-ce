// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package chain

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

func JWT() gin.HandlerFunc {
	ums := userManageService()
	return func(c *gin.Context) {
		var (
			token string
			ecode = xerror.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization")

			// 验证前端传过来的token格式，不为空，开头为Bearer
			if token == "" || !strings.HasPrefix(token, "Bearer ") {
				response := app.NewResponse(c)
				response.ToErrorResponse(xerror.UnauthorizedTokenError)
				c.Abort()
				return
			}

			// 验证通过，提取有效部分（除去Bearer)
			token = token[7:]
		}
		if token == "" {
			ecode = xerror.InvalidParams
		} else {
			claims, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = xerror.UnauthorizedTokenTimeout
				default:
					ecode = xerror.UnauthorizedTokenError
				}
			} else {
				c.Set("UID", claims.UID)
				c.Set("USERNAME", claims.Username)

				// 加载用户信息
				user, err := ums.GetUserByID(claims.UID)
				if err == nil {
					c.Set("USER", user)
				} else {
					ecode = xerror.UnauthorizedAuthNotExist
				}

				// 强制下线机制
				if (conf.JWTSetting.Issuer + ":" + user.Salt) != claims.Issuer {
					ecode = xerror.UnauthorizedTokenTimeout
				}
			}
		}

		if ecode != xerror.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}

func JwtLoose() gin.HandlerFunc {
	ums := userManageService()
	return func(c *gin.Context) {
		token, exist := c.GetQuery("token")
		if !exist {
			token = c.GetHeader("Authorization")
			// 验证前端传过来的token格式，不为空，开头为Bearer
			if strings.HasPrefix(token, "Bearer ") {
				// 验证通过，提取有效部分（除去Bearer)
				token = token[7:]
			} else {
				c.Next()
			}
		}
		if len(token) > 0 {
			if claims, err := app.ParseToken(token); err == nil {
				c.Set("UID", claims.UID)
				c.Set("USERNAME", claims.Username)
				// 加载用户信息
				user, err := ums.GetUserByID(claims.UID)
				if err == nil && (conf.JWTSetting.Issuer+":"+user.Salt) == claims.Issuer {
					c.Set("USER", user)
				}
			}
		}
		c.Next()
	}
}
