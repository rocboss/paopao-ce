// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ms

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

const (
	AttachmentTypeImage = dbr.AttachmentTypeImage
	AttachmentTypeVideo = dbr.AttachmentTypeVideo
	AttachmentTypeOther = dbr.AttachmentTypeOther

	// 类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源
	ContentTypeTitle            = dbr.ContentTypeTitle
	ContentTypeText             = dbr.ContentTypeText
	ContentTypeImage            = dbr.ContentTypeImage
	ContentTypeVideo            = dbr.ContentTypeVideo
	ContentTypeAudio            = dbr.ContentTypeAudio
	ContentTypeLink             = dbr.ContentTypeLink
	ContentTypeAttachment       = dbr.ContentTypeAttachment
	ContentTypeChargeAttachment = dbr.ContentTypeChargeAttachment
)

const (
	PostVisitPublic PostVisibleT = iota
	PostVisitPrivate
	PostVisitFriend
	PostVisitInvalid
)

type (
	PostStar           = dbr.PostStar
	PostCollection     = dbr.PostCollection
	PostAttachmentBill = dbr.PostAttachmentBill
	PostContent        = dbr.PostContent
	Attachment         = dbr.Attachment
	AttachmentType     = dbr.AttachmentType
	PostContentT       = dbr.PostContentT
	PostVisibleT       = dbr.PostVisibleT
)
