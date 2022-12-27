// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package base

import (
	"net/http"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

type BaseServant struct {
	// TODO
}

type DaoServant struct {
	Redis *redis.Client
	Ds    core.DataService
}

type BaseBinding struct {
	// TODO
}

type BaseRender struct {
	// TODO
}

type JsonResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:",omitempty"`
}

type UserSetter interface {
	SetUser(*core.User)
}

func UserFrom(c *gin.Context) (*core.User, bool) {
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
		return mir.NewError(xerror.InvalidParams.StatusCode(), xerror.InvalidParams.WithDetails(errs.Error()))
	}
	// setup *core.User if needed
	if setter, ok := obj.(UserSetter); ok {
		user, _ := UserFrom(c)
		setter.SetUser(user)
	}
	return nil
}

func RenderAny(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, &JsonResp{
			Code: 0,
			Msg:  "success",
			Data: data,
		})
	} else {
		c.JSON(xerror.HttpStatusCode(err), &JsonResp{
			Code: err.StatusCode(),
			Msg:  err.Error(),
		})
	}
}
