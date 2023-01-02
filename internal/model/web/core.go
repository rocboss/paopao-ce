// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

type ChangeAvatarReq struct {
	BaseInfo `json:"-" binding:"-"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required"`
}

type SyncSearchIndexReq struct {
	BaseInfo `json:"-" binding:"-"`
}

type UserInfoReq struct {
	BaseInfo `json:"-" binding:"-"`
	Username string `json:"username" form:"username" binding:"required"`
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
	SimpleInfo `json:"-" binding:"-"`
}

type GetUnreadMsgCountResp struct {
	Count int64 `json:"count"`
}

type GetMessagesReq BasePageReq

type GetMessagesResp base.PageResp

type ReadMessageReq struct {
	SimpleInfo `json:"-" binding:"-"`
	ID         int64 `json:"id" binding:"required"`
}
type SendWhisperReq struct {
	SimpleInfo `json:"-" binding:"-"`
	UserID     int64  `json:"user_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type GetCollectionsReq BasePageReq
type GetCollectionsResp base.PageResp

type GetStarsReq BasePageReq
type GetStarsResp base.PageResp

type UserPhoneBindReq struct {
	BaseInfo `json:"-" binding:"-"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Captcha  string `json:"captcha" form:"captcha" binding:"required"`
}

type ChangePasswordReq struct {
	BaseInfo    `json:"-" binding:"-"`
	Password    string `json:"password" form:"password" binding:"required"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
}

type ChangeNicknameReq struct {
	BaseInfo `json:"-" binding:"-"`
	Nickname string `json:"nickname" form:"nickname" binding:"required"`
}

type SuggestUsersReq struct {
	Keyword string
}

type SuggestUsersResp struct {
	Suggests []string `json:"suggest"`
}

type SuggestTagsReq struct {
	Keyword string
}

type SuggestTagsResp struct {
	Suggests []string `json:"suggest"`
}

type TweetStarStatusReq struct {
	SimpleInfo `json:"-" binding:"-"`
	TweetId    int64 `form:"id"`
}

type TweetStarStatusResp struct {
	Status bool `json:"status"`
}

type TweetCollectionStatusReq struct {
	SimpleInfo `json:"-" binding:"-"`
	TweetId    int64 `form:"id"`
}

type TweetCollectionStatusResp struct {
	Status bool `json:"status"`
}
