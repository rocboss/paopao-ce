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
	"github.com/rocboss/paopao-ce/pkg/types"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

type BaseServant struct {
	// TODO
}

type DaoServant struct {
	Redis *redis.Client
	Ds    core.DataService
}

type BaseBinding types.Empty

type BaseRender types.Empty

type JsonResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:",omitempty"`
}

type UserSetter interface {
	SetUser(*core.User)
}

type UserIdSetter interface {
	SetUserId(int64)
}

func UserFrom(c *gin.Context) (*core.User, bool) {
	if u, exists := c.Get("USER"); exists {
		user, ok := u.(*core.User)
		return user, ok
	}
	return nil, false
}

func UserIdFrom(c *gin.Context) (int64, bool) {
	if uid, exists := c.Get("UID"); exists {
		v, ok := uid.(int64)
		return v, ok
	}
	return -1, false
}

func UserNameFrom(c *gin.Context) (string, bool) {
	if username, exists := c.Get("USERNAME"); exists {
		v, ok := username.(string)
		return v, ok
	}
	return "", false
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
	// setup UserId if needed
	if setter, ok := obj.(UserIdSetter); ok {
		uid, _ := UserIdFrom(c)
		setter.SetUserId(uid)
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

func (s *DaoServant) GetTweetBy(id int64) (*core.PostFormated, error) {
	post, err := s.Ds.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	postContents, err := s.Ds.GetPostContentsByIDs([]int64{post.ID})
	if err != nil {
		return nil, err
	}
	users, err := s.Ds.GetUsersByIDs([]int64{post.UserID})
	if err != nil {
		return nil, err
	}
	// 数据整合
	postFormated := post.Format()
	for _, user := range users {
		postFormated.User = user.Format()
	}
	for _, content := range postContents {
		if content.PostID == post.ID {
			postFormated.Contents = append(postFormated.Contents, content.Format())
		}
	}
	return postFormated, nil
}
