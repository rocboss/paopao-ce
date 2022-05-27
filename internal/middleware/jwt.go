package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization")

			// 验证前端传过来的token格式，不为空，开头为Bearer
			if token == "" || !strings.HasPrefix(token, "Bearer ") {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.UnauthorizedTokenError)
				c.Abort()
				return
			}

			// 验证通过，提取有效部分（除去Bearer)
			token = token[7:]
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			claims, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else {
				c.Set("UID", claims.UID)
				c.Set("USERNAME", claims.Username)

				// 加载用户信息
				user := &model.User{
					Model: &model.Model{
						ID: claims.UID,
					},
				}
				user, _ = user.Get(global.DBEngine)
				c.Set("USER", user)

				// 强制下线机制
				if (global.JWTSetting.Issuer + ":" + user.Salt) != claims.Issuer {
					ecode = errcode.UnauthorizedTokenTimeout
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
