// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

const (
	TagTypeHot       TagType = "hot"
	TagTypeNew       TagType = "new"
	TagTypeFollow    TagType = "follow"
	TagTypeHotExtral TagType = "hot_extral"
)

type TagType string

type TimelineReq struct {
	BaseInfo   `form:"-"  binding:"-"`
	Query      string              `form:"query"`
	Visibility []core.PostVisibleT `form:"query"`
	Type       string              `form:"type"`
	Page       int                 `form:"-"  binding:"-"`
	PageSize   int                 `form:"-"  binding:"-"`
}

type TimelineResp base.PageResp

type GetUserTweetsReq struct {
	BaseInfo `form:"-" binding:"-"`
	Username string `form:"username" binding:"required"`
	Page     int    `form:"-" binding:"-"`
	PageSize int    `form:"-" binding:"-"`
}

type GetUserTweetsResp base.PageResp

type GetUserProfileReq struct {
	BaseInfo `form:"-" binding:"-"`
	Username string `form:"username" binding:"required"`
}

type GetUserProfileResp struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	IsAdmin  bool   `json:"is_admin"`
	IsFriend bool   `json:"is_friend"`
}

type TopicListReq struct {
	BaseInfo `form:"-"  binding:"-"`
	Type     TagType `json:"type" form:"type" binding:"required"`
	UserId   int     `json:"uid" form:"uid"`
	Num      int     `json:"num" form:"num" binding:"required"`
}

// TopicListResp 主题返回值
// TODO: 优化内容定义
type TopicListResp struct {
	Topics       []*core.TagFormated `json:"topics"`
	ExtralTopics []*core.TagFormated `json:"extral_topics,omitempty"`
}

func (r *GetUserTweetsReq) SetPageInfo(page int, pageSize int) {
	r.Page, r.PageSize = page, pageSize
}
