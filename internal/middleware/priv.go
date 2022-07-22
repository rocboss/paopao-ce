package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/errcode"
)

func Priv() gin.HandlerFunc {
	enablePhoneBind := conf.CfgIf("PhoneBind")
	return func(c *gin.Context) {
		if user, exist := c.Get("USER"); exist {
			if userModel, ok := user.(*model.User); ok {
				if userModel.Status == model.UserStatusNormal {
					if enablePhoneBind && userModel.Phone == "" {
						response := app.NewResponse(c)
						response.ToErrorResponse(errcode.AccountNoPhoneBind)
						c.Abort()
						return
					}
					c.Next()
					return
				}
			}
		}
		response := app.NewResponse(c)
		response.ToErrorResponse(errcode.UserHasBeenBanned)
		c.Abort()
	}
}
