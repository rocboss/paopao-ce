// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package base

import (
	"net/http"
	"os"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

type BaseServant struct {
	// TODO
}

type BaseBinding struct {
	// TODO
}

type BaseRender struct {
	// TODO
}

func (BaseServant) userFrom(c *gin.Context) (*core.User, bool) {
	if u, exists := c.Get("USER"); exists {
		user, ok := u.(*core.User)
		return user, ok
	}
	return nil, false
}

func BindAny(c *gin.Context, obj any) mir.Error {
	var errs xerror.ValidErrors
	err := c.ShouldBind(obj)
	if err != nil {
		return mir.NewError(xerror.InvalidParams.Code(), xerror.InvalidParams.WithDetails(errs.Error()))
	}
	return nil
}

func RenderAny(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		hostname, _ := os.Hostname()
		if data == nil {
			data = gin.H{
				"code":      0,
				"msg":       "success",
				"tracehost": hostname,
			}
		} else {
			data = gin.H{
				"code":      0,
				"msg":       "success",
				"data":      data,
				"tracehost": hostname,
			}
		}
		c.JSON(http.StatusOK, data)
	} else {
		// TODO: error process logic
		resp := gin.H{"code": err.StatusCode(), "msg": err.Error()}
		// xerr := &xerror.Error{}
		// if errors.As(err, xerr) {
		// 	resp["code"], resp["msg"] = xerr.Code(), xerr.Msg()
		// 	details := xerr.Details()
		// 	if len(details) > 0 {
		// 		resp["details"] = details
		// 	}
		// }
		c.JSON(http.StatusInternalServerError, resp)
	}
}
