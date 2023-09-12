// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

type TweetCommentThumbsReq struct {
	SimpleInfo `json:"-" binding:"-"`
	TweetId    int64 `json:"tweet_id" binding:"required"`
	CommentId  int64 `json:"comment_id" binding:"required"`
}

type TweetReplyThumbsReq struct {
	SimpleInfo `json:"-" binding:"-"`
	TweetId    int64 `json:"tweet_id" binding:"required"`
	CommentId  int64 `json:"comment_id" binding:"required"`
	ReplyId    int64 `json:"reply_id" binding:"required"`
}

type PostContentItem struct {
	Content string          `json:"content"  binding:"required"`
	Type    ms.PostContentT `json:"type"  binding:"required"`
	Sort    int64           `json:"sort"  binding:"required"`
}

type CreateTweetReq struct {
	BaseInfo        `json:"-" binding:"-"`
	Contents        []*PostContentItem `json:"contents" binding:"required"`
	Tags            []string           `json:"tags" binding:"required"`
	Users           []string           `json:"users" binding:"required"`
	AttachmentPrice int64              `json:"attachment_price"`
	Visibility      core.PostVisibleT  `json:"visibility"`
	ClientIP        string             `json:"-" binding:"-"`
}

type CreateTweetResp ms.PostFormated

type DeleteTweetReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type StarTweetReq struct {
	SimpleInfo `json:"-" binding:"-"`
	ID         int64 `json:"id" binding:"required"`
}

type StarTweetResp struct {
	Status bool `json:"status"`
}

type CollectionTweetReq struct {
	SimpleInfo `json:"-" binding:"-"`
	ID         int64 `json:"id" binding:"required"`
}

type CollectionTweetResp struct {
	Status bool `json:"status"`
}

type LockTweetReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type LockTweetResp struct {
	LockStatus int `json:"lock_status"`
}

type StickTweetReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type HighlightTweetReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type StickTweetResp struct {
	StickStatus int `json:"top_status"`
}

type HighlightTweetResp struct {
	HighlightStatus int `json:"highlight_status"`
}

type VisibleTweetReq struct {
	BaseInfo   `json:"-" binding:"-"`
	ID         int64             `json:"id"`
	Visibility core.PostVisibleT `json:"visibility"`
}

type VisibleTweetResp struct {
	Visibility core.PostVisibleT `json:"visibility"`
}

type CreateCommentReq struct {
	SimpleInfo `json:"-" binding:"-"`
	PostID     int64              `json:"post_id" binding:"required"`
	Contents   []*PostContentItem `json:"contents" binding:"required"`
	Users      []string           `json:"users" binding:"required"`
	ClientIP   string             `json:"-" binding:"-"`
}

type CreateCommentResp ms.Comment

type CreateCommentReplyReq struct {
	SimpleInfo `json:"-" binding:"-"`
	CommentID  int64  `json:"comment_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
	AtUserID   int64  `json:"at_user_id"`
	ClientIP   string `json:"-" binding:"-"`
}

type CreateCommentReplyResp ms.CommentReply

type DeleteCommentReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}
type DeleteCommentReplyReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type UploadAttachmentReq struct {
	SimpleInfo  `json:"-" binding:"-"`
	UploadType  string
	ContentType string
	File        multipart.File
	FileSize    int64
	FileExt     string
}

type UploadAttachmentResp struct {
	UserID    int64             `json:"user_id"`
	FileSize  int64             `json:"file_size"`
	ImgWidth  int               `json:"img_width"`
	ImgHeight int               `json:"img_height"`
	Type      ms.AttachmentType `json:"type"`
	Content   string            `json:"content"`
}

type DownloadAttachmentPrecheckReq struct {
	BaseInfo  `form:"-" binding:"-"`
	ContentID int64 `form:"id"`
}

type DownloadAttachmentPrecheckResp struct {
	Paid bool `json:"paid"`
}

type DownloadAttachmentReq struct {
	BaseInfo  `form:"-" binding:"-"`
	ContentID int64 `form:"id"`
}

type DownloadAttachmentResp struct {
	SignedURL string `json:"signed_url"`
}

type StickTopicReq struct {
	SimpleInfo `json:"-" binding:"-"`
	TopicId    int64 `json:"topic_id" binding:"required"`
}

type StickTopicResp struct {
	StickStatus int8 `json:"top_status"`
}

type FollowTopicReq struct {
	SimpleInfo `json:"-" binding:"-"`
	TopicId    int64 `json:"topic_id" binding:"required"`
}

type UnfollowTopicReq struct {
	SimpleInfo `json:"-" binding:"-"`
	TopicId    int64 `json:"topic_id" binding:"required"`
}

// Check 检查PostContentItem属性
func (p *PostContentItem) Check(acs core.AttachmentCheckService) error {
	// 检查附件是否是本站资源
	if p.Type == ms.ContentTypeImage || p.Type == ms.ContentTypeVideo || p.Type == ms.ContentTypeAttachment {
		if err := acs.CheckAttachment(p.Content); err != nil {
			return err
		}
	}
	// 检查链接是否合法
	if p.Type == ms.ContentTypeLink {
		if strings.Index(p.Content, "http://") != 0 && strings.Index(p.Content, "https://") != 0 {
			return fmt.Errorf("链接不合法")
		}
	}
	return nil
}

func (r *UploadAttachmentReq) Bind(c *gin.Context) (xerr mir.Error) {
	userId, exist := base.UserIdFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}

	uploadType := c.Request.FormValue("type")
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		return ErrFileUploadFailed
	}
	defer func() {
		if xerr != nil {
			file.Close()
		}
	}()

	if err := fileCheck(uploadType, fileHeader.Size); err != nil {
		return err
	}
	contentType := fileHeader.Header.Get("Content-Type")
	fileExt, xerr := getFileExt(contentType)
	if xerr != nil {
		return xerr
	}
	r.SimpleInfo = SimpleInfo{
		Uid: userId,
	}
	r.UploadType, r.ContentType = uploadType, contentType
	r.File, r.FileSize, r.FileExt = file, fileHeader.Size, fileExt
	return nil
}

func (r *DownloadAttachmentPrecheckReq) Bind(c *gin.Context) mir.Error {
	user, exist := base.UserFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}
	r.BaseInfo = BaseInfo{
		User: user,
	}
	r.ContentID = convert.StrTo(c.Query("id")).MustInt64()
	return nil
}

func (r *DownloadAttachmentReq) Bind(c *gin.Context) mir.Error {
	user, exist := base.UserFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}
	r.BaseInfo = BaseInfo{
		User: user,
	}
	r.ContentID = convert.StrTo(c.Query("id")).MustInt64()
	return nil
}

func (r *CreateTweetReq) Bind(c *gin.Context) mir.Error {
	r.ClientIP = c.ClientIP()
	return bindAny(c, r)
}

func (r *CreateCommentReplyReq) Bind(c *gin.Context) mir.Error {
	r.ClientIP = c.ClientIP()
	return bindAny(c, r)
}

func (r *CreateCommentReq) Bind(c *gin.Context) mir.Error {
	r.ClientIP = c.ClientIP()
	return bindAny(c, r)
}

func (r *CreateTweetResp) Render(c *gin.Context) {
	c.JSON(http.StatusOK, &joint.JsonResp{
		Code: 0,
		Msg:  "success",
		Data: r,
	})
	// 设置审核元信息，用于接下来的审核逻辑
	c.Set(AuditHookCtxKey, &AuditMetaInfo{
		Style: AuditStyleUserTweet,
		Id:    r.ID,
	})
}
