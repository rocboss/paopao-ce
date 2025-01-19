// Copyright 2025 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/model/joint"
)

const (
	TagTypeHot       = cs.TagTypeHot
	TagTypeNew       = cs.TagTypeNew
	TagTypeFollow    = cs.TagTypeFollow
	TagTypePin       = cs.TagTypePin
	TagTypeHotExtral = cs.TagTypeHotExtral
)

const (
	UserPostsStylePost      = "post"
	UserPostsStyleComment   = "comment"
	UserPostsStyleHighlight = "highlight"
	UserPostsStyleMedia     = "media"
	UserPostsStyleStar      = "star"

	StyleTweetsNewest    = "newest"
	StyleTweetsHots      = "hots"
	StyleTweetsFollowing = "following"
)

type TagType = cs.TagType

type CommentStyleType string

type TweetCommentsReq struct {
	SimpleInfo
	joint.PageInfo
	TweetId  int64            `form:"id" query:"id" binding:"required"`
	Style    CommentStyleType `form:"style" query:"style"`
	Page     int              `form:"page" query:"page" binding:"-"`
	PageSize int              `form:"page_size" query:"page_size" binding:"-"`
}

type TweetCommentsResp struct {
	joint.CachePageResp
}

type TimelineReq struct {
	BaseInfo
	joint.PageInfo
	Query      string              `form:"query" query:"query"`
	Visibility []core.PostVisibleT `form:"visibility" query:"visibility"`
	Type       string              `form:"type" query:"type"`
	Style      string              `form:"style" query:"query"`
	Page       int                 `form:"page" query:"page"`
	PageSize   int                 `form:"page_size" query:"page_size"`
}

type TimelineResp struct {
	joint.CachePageResp
}

type GetUserTweetsReq struct {
	BaseInfo `form:"-" binding:"-"`
	joint.PageInfo
	Username string `form:"username" query:"username" binding:"required"`
	Style    string `form:"style" query:"style"`
	Page     int    `form:"page" query:"page"`
	PageSize int    `form:"page_size" query:"page_size"`
}

type GetUserTweetsResp struct {
	joint.CachePageResp
}

type GetUserProfileReq struct {
	BaseInfo
	Username string `form:"username" query:"username" binding:"required"`
}

type GetUserProfileResp struct {
	ID          int64  `json:"id"`
	Nickname    string `json:"nickname"`
	Username    string `json:"username"`
	Status      int    `json:"status"`
	Avatar      string `json:"avatar"`
	IsAdmin     bool   `json:"is_admin"`
	IsFriend    bool   `json:"is_friend"`
	IsFollowing bool   `json:"is_following"`
	CreatedOn   int64  `json:"created_on"`
	Follows     int64  `json:"follows"`
	Followings  int64  `json:"followings"`
	TweetsCount int    `json:"tweets_count"`
}

type TopicListReq struct {
	SimpleInfo
	Type      TagType `json:"type" form:"type" query:"type" binding:"required"`
	Num       int     `json:"num" form:"num" query:"num" binding:"required"`
	ExtralNum int     `json:"extral_num" form:"extral_num" query:"extral_num"`
}

// TopicListResp 主题返回值
// TODO: 优化内容定义
type TopicListResp struct {
	Topics       cs.TagList `json:"topics"`
	ExtralTopics cs.TagList `json:"extral_topics,omitempty"`
}

type TweetDetailReq struct {
	BaseInfo
	TweetId int64 `form:"id" query:"id"`
}

type TweetDetailResp ms.PostFormated

func (r *GetUserTweetsReq) Ajust(page int, pageSize int) {
	r.BuildPageInfo(r.Page, r.PageSize)
}

func (r *TweetCommentsReq) Ajust(page int, pageSize int) {
	r.BuildPageInfo(r.Page, r.PageSize)
}

func (r *TimelineReq) Ajust() {
	r.BuildPageInfo(r.Page, r.PageSize)
}

func (s CommentStyleType) ToInnerValue() (res cs.StyleCommentType) {
	switch s {
	case "hots":
		res = cs.StyleCommentHots
	case "newest":
		res = cs.StyleCommentNewest
	case "default":
		fallthrough
	default:
		res = cs.StyleCommentDefault
	}
	return
}

func (s CommentStyleType) String() (res string) {
	switch s {
	case "default":
		res = conf.InfixCommentDefault
	case "hots":
		res = conf.InfixCommentHots
	case "newest":
		res = conf.InfixCommentNewest
	default:
		res = "_"
	}
	return
}
