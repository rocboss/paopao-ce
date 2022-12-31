// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import "github.com/rocboss/paopao-ce/internal/servants/base"

type RequestingFriendReq struct {
	BaseInfo  `json:"-" binding:"-"`
	UserId    int64  `json:"user_id" binding:"required"`
	Greetings string `json:"greetings" binding:"required"`
}

type AddFriendReq struct {
	BaseInfo `json:"-" binding:"-"`
	UserId   int64 `json:"user_id" binding:"required"`
}

type RejectFriendReq struct {
	BaseInfo `json:"-" binding:"-"`
	UserId   int64 `json:"user_id" binding:"required"`
}

type DeleteFriendReq struct {
	BaseInfo `json:"-" binding:"-"`
	UserId   int64 `json:"user_id"`
}

type GetContactsReq struct {
	BaseInfo `form:"-" binding:"-"`
	Page     int `form:"-" binding:"-"`
	PageSize int `form:"-" binding:"-"`
}

type GetContactsResp base.PageResp

func (r *GetContactsReq) SetPageInfo(page int, pageSize int) {
	r.Page, r.PageSize = page, pageSize
}
