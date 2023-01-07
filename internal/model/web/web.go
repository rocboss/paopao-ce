// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

type BaseInfo struct {
	User *core.User
}

type SimpleInfo struct {
	Uid int64
}

type BasePageReq struct {
	UserId   int64
	Page     int
	PageSize int
}

func (b *BaseInfo) SetUser(user *core.User) {
	b.User = user
}

func (s *SimpleInfo) SetUserId(id int64) {
	s.Uid = id
}

func BasePageReqFrom(c *gin.Context) (*BasePageReq, mir.Error) {
	uid, ok := base.UserIdFrom(c)
	if !ok {
		return nil, xerror.UnauthorizedTokenError
	}
	page, pageSize := app.GetPageInfo(c)
	return &BasePageReq{
		UserId:   uid,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
