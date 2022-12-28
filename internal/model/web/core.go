// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"context"

	"github.com/rocboss/paopao-ce/internal/servants/base"
)

type ChangeAvatarReq struct {
	*BaseInfo `json:"-" binding:"-"`
}

type SyncSearchIndexReq struct {
	*BaseInfo `json:"-" binding:"-"`

	Ctx context.Context `json:"-" binding:"-"`
}

type UserInfoReq struct {
	*BaseInfo `json:"-" binding:"-"`
	Username  string `json:"username" form:"username" binding:"required"`
}

type UserInfoResp struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	Balance  int64  `json:"balance"`
	Phone    string `json:"phone"`
	IsAdmin  bool   `json:"is_admin"`
}

type GetUnreadMsgCountReq struct {
	*SimpleInfo `json:"-" binding:"-"`
}

type GetUnreadMsgCountResp struct {
	Count int64 `json:"count"`
}

type GetMessagesReq struct {
	UserId   int64
	Page     int
	PageSize int
}

type GetMessagesResp = base.PageResp

type ReadMessageReq struct {
	*SimpleInfo `json:"-" binding:"-"`

	ID int64 `json:"id" binding:"required"`
}
