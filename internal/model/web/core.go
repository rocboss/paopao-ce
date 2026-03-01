// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

type MessageStyle = cs.MessageStyle

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
	Id          int64  `json:"id"`
	Nickname    string `json:"nickname"`
	Username    string `json:"username"`
	Status      int    `json:"status"`
	Avatar      string `json:"avatar"`
	Balance     int64  `json:"balance"`
	Phone       string `json:"phone"`
	IsAdmin     bool   `json:"is_admin"`
	CreatedOn   int64  `json:"created_on"`
	Follows     int64  `json:"follows"`
	Followings  int64  `json:"followings"`
	TweetsCount int    `json:"tweets_count"`
	Experience  int    `json:"experience"`
	Level       int    `json:"level"`
}

type GetMessagesReq struct {
	SimpleInfo `json:"-" binding:"-"`
	joint.BasePageInfo
	Style MessageStyle `form:"style" binding:"required"`
}

type GetMessagesResp struct {
	joint.CachePageResp
}

type ReadMessageReq struct {
	SimpleInfo `json:"-" binding:"-"`
	ID         int64 `json:"id" binding:"required"`
}

type ReadAllMessageReq struct {
	SimpleInfo `json:"-" binding:"-"`
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

func (r *UserInfoReq) Bind(c *gin.Context) error {
	username, exist := base.UserNameFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}
	r.Username = username
	return nil
}

func (r *GetCollectionsReq) Bind(c *gin.Context) error {
	return (*BasePageReq)(r).Bind(c)
}

func (r *GetStarsReq) Bind(c *gin.Context) error {
	return (*BasePageReq)(r).Bind(c)
}

func (r *SuggestTagsReq) Bind(c *gin.Context) error {
	r.Keyword = c.Query("k")
	return nil
}

func (r *SuggestUsersReq) Bind(c *gin.Context) error {
	r.Keyword = c.Query("k")
	return nil
}

func (r *TweetCollectionStatusReq) Bind(c *gin.Context) error {
	userId, exist := base.UserIdFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}
	r.SimpleInfo = SimpleInfo{
		Uid: userId,
	}
	r.TweetId = convert.StrTo(c.Query("id")).MustInt64()
	return nil
}

func (r *TweetStarStatusReq) Bind(c *gin.Context) error {
	UserId, exist := base.UserIdFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}
	r.SimpleInfo = SimpleInfo{
		Uid: UserId,
	}
	r.TweetId = convert.StrTo(c.Query("id")).MustInt64()
	return nil
}

type StreamMessagesReq struct {
	SimpleInfo `json:"-" binding:"-"`
}

// 表情包相关请求和响应

// UploadEmojiReq 上传表情包请求
type UploadEmojiReq struct {
	BaseInfo    `json:"-" binding:"-"`
	Name        string      `form:"name" binding:"required"`
	File        interface{} `form:"file" binding:"required"`
	FileExt     string      `form:"file_ext" binding:"required"`
	FileSize    int64       `form:"file_size" binding:"required"`
	ContentType string      `form:"content_type" binding:"required"`
	Width       int         `form:"width" binding:"required"`
	Height      int         `form:"height" binding:"required"`
}

// UploadEmojiResp 上传表情包响应
type UploadEmojiResp struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// GetEmojiListReq 获取表情包列表请求
type GetEmojiListReq struct {
	SimpleInfo `json:"-" binding:"-"`
	BasePageInfo
}

// GetEmojiListResp 获取表情包列表响应
type GetEmojiListResp struct {
	base.PageResp
}

// GetUserEmojiListReq 获取用户上传的表情包列表请求
type GetUserEmojiListReq struct {
	SimpleInfo `json:"-" binding:"-"`
	BasePageInfo
	UserID int64 `form:"user_id" binding:"required"`
}

// GetUserEmojiListResp 获取用户上传的表情包列表响应
type GetUserEmojiListResp struct {
	base.PageResp
}

// GetUserCollectedEmojiListReq 获取用户收藏的表情包列表请求
type GetUserCollectedEmojiListReq struct {
	SimpleInfo `json:"-" binding:"-"`
	BasePageInfo
}

// GetUserCollectedEmojiListResp 获取用户收藏的表情包列表响应
type GetUserCollectedEmojiListResp struct {
	base.PageResp
}

// CollectEmojiReq 收藏表情包请求
type CollectEmojiReq struct {
	SimpleInfo `json:"-" binding:"-"`
	EmojiID    int64 `json:"emoji_id" binding:"required"`
}

// CollectEmojiResp 收藏表情包响应
type CollectEmojiResp struct {
	Status bool `json:"status"`
}

// UncollectEmojiReq 取消收藏表情包请求
type UncollectEmojiReq struct {
	SimpleInfo `json:"-" binding:"-"`
	EmojiID    int64 `json:"emoji_id" binding:"required"`
}

// UncollectEmojiResp 取消收藏表情包响应
type UncollectEmojiResp struct {
	Status bool `json:"status"`
}

// DeleteEmojiReq 删除表情包请求
type DeleteEmojiReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}
