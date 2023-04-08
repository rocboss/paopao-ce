// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
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
	CreateAttachment(obj *cs.Attachment) (int64, error)
}

// TweetHelpService 推文辅助服务
type TweetHelpService interface {
	RevampPosts(posts []*PostFormated) ([]*PostFormated, error)
	MergePosts(posts []*Post) ([]*PostFormated, error)
}

// TweetServantA 推文检索服务(版本A)
type TweetServantA interface {
	TweetInfoById(id int64) (*cs.TweetInfo, error)
	TweetItemById(id int64) (*cs.TweetItem, error)
	UserTweets(visitorId, userId int64) (cs.TweetList, error)
	ReactionByTweetId(userId int64, tweetId int64) (*cs.ReactionItem, error)
	UserReactions(userId int64, limit int, offset int) (cs.ReactionList, error)
	FavoriteByTweetId(userId int64, tweetId int64) (*cs.FavoriteItem, error)
	UserFavorites(userId int64, limit int, offset int) (cs.FavoriteList, error)
	AttachmentByTweetId(userId int64, tweetId int64) (*cs.AttachmentBill, error)
}

// TweetManageServantA 推文管理服务，包括创建/删除/更新推文(版本A)
type TweetManageServantA interface {
	CreateAttachment(obj *cs.Attachment) (int64, error)
	CreateTweet(userId int64, req *cs.NewTweetReq) (*cs.TweetItem, error)
	DeleteTweet(userId int64, tweetId int64) ([]string, error)
	LockTweet(userId int64, tweetId int64) error
	StickTweet(userId int64, tweetId int64) error
	VisibleTweet(userId int64, visibility cs.TweetVisibleType) error
	CreateReaction(userId int64, tweetId int64) error
	DeleteReaction(userId int64, reactionId int64) error
	CreateFavorite(userId int64, tweetId int64) error
	DeleteFavorite(userId int64, favoriteId int64) error
}

// TweetHelpServantA 推文辅助服务(版本A)
type TweetHelpServantA interface {
	RevampTweets(tweets cs.TweetList) (cs.TweetList, error)
	MergeTweets(tweets cs.TweetInfo) (cs.TweetList, error)
}
