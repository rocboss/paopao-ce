// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/rocboss/paopao-ce/internal/core"
)

type PostContentItem struct {
	Content string            `json:"content"  binding:"required"`
	Type    core.PostContentT `json:"type"  binding:"required"`
	Sort    int64             `json:"sort"  binding:"required"`
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

type CreateTweetResp core.PostFormated

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

type StickTweetResp struct {
	StickStatus int `json:"top_status"`
}

type VisiblePostReq struct {
	BaseInfo   `json:"-" binding:"-"`
	ID         int64             `json:"id" binding:"required"`
	Visibility core.PostVisibleT `json:"visibility" binding:"required"`
}

type VisiblePostResp struct {
	Visibility core.PostVisibleT `json:"visibility"`
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
	UserID    int64               `json:"user_id"`
	FileSize  int64               `json:"file_size"`
	ImgWidth  int                 `json:"img_width"`
	ImgHeight int                 `json:"img_height"`
	Type      core.AttachmentType `json:"type"`
	Content   string              `json:"content"`
}

type DownloadAttachmentPrecheckReq struct {
	BaseInfo  `json:"-" binding:"-"`
	ContentID int64 `form:"id"`
}

type DownloadAttachmentPrecheckResp struct {
	Paid bool `json:"paid"`
}

type DownloadAttachmentReq struct {
	BaseInfo  `json:"-" binding:"-"`
	ContentID int64 `form:"id"`
}

type DownloadAttachmentResp struct {
	SignedURL string `json:"signed_url"`
}

// Check 检查PostContentItem属性
func (p *PostContentItem) Check(acs core.AttachmentCheckService) error {
	// 检查附件是否是本站资源
	if p.Type == core.ContentTypeImage || p.Type == core.ContentTypeVideo || p.Type == core.ContentTypeAttachment {
		if err := acs.CheckAttachment(p.Content); err != nil {
			return err
		}
	}
	// 检查链接是否合法
	if p.Type == core.ContentTypeLink {
		if strings.Index(p.Content, "http://") != 0 && strings.Index(p.Content, "https://") != 0 {
			return fmt.Errorf("链接不合法")
		}
	}
	return nil
}
