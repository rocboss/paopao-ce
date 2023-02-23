// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

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

type (
	PostStar           = dbr.PostStar
	PostCollection     = dbr.PostCollection
	PostAttachmentBill = dbr.PostAttachmentBill
	PostContent        = dbr.PostContent
	Attachment         = dbr.Attachment
	AttachmentType     = dbr.AttachmentType
	PostContentT       = dbr.PostContentT

	IndexTweetList struct {
		Tweets []*PostFormated
		Total  int64
	}
)

// TweetService 推文检索服务
type TweetService interface {
	GetPostByID(id int64) (*Post, error)
	GetPosts(conditions *ConditionsT, offset, limit int) ([]*Post, error)
	GetPostCount(conditions *ConditionsT) (int64, error)
	GetUserPostStar(postID, userID int64) (*PostStar, error)
	GetUserPostStars(userID int64, offset, limit int) ([]*PostStar, error)
	GetUserPostStarCount(userID int64) (int64, error)
	GetUserPostCollection(postID, userID int64) (*PostCollection, error)
	GetUserPostCollections(userID int64, offset, limit int) ([]*PostCollection, error)
	GetUserPostCollectionCount(userID int64) (int64, error)
	GetPostAttatchmentBill(postID, userID int64) (*PostAttachmentBill, error)
	GetPostContentsByIDs(ids []int64) ([]*PostContent, error)
	GetPostContentByID(id int64) (*PostContent, error)
}

// TweetManageService 推文管理服务，包括创建/删除/更新推文
type TweetManageService interface {
	CreateAttachment(attachment *Attachment) (*Attachment, error)
	CreatePost(post *Post) (*Post, error)
	DeletePost(post *Post) ([]string, error)
	LockPost(post *Post) error
	StickPost(post *Post) error
	VisiblePost(post *Post, visibility PostVisibleT) error
	UpdatePost(post *Post) error
	CreatePostStar(postID, userID int64) (*PostStar, error)
	DeletePostStar(p *PostStar) error
	CreatePostCollection(postID, userID int64) (*PostCollection, error)
	DeletePostCollection(p *PostCollection) error
	CreatePostContent(content *PostContent) (*PostContent, error)
}

// TweetHelpService 推文辅助服务
type TweetHelpService interface {
	RevampPosts(posts []*PostFormated) ([]*PostFormated, error)
	MergePosts(posts []*Post) ([]*PostFormated, error)
}

// IndexPostsService 广场首页推文列表服务
type IndexPostsService interface {
	IndexPosts(user *User, offset int, limit int) (*IndexTweetList, error)
}
