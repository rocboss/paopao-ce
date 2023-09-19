// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
)

// TweetService 推文检索服务
type TweetService interface {
	GetPostByID(id int64) (*ms.Post, error)
	GetPosts(conditions ms.ConditionsT, offset, limit int) ([]*ms.Post, error)
	GetPostCount(conditions ms.ConditionsT) (int64, error)
	GetUserPostStar(postID, userID int64) (*ms.PostStar, error)
	GetUserPostStars(userID int64, limit int, offset int) ([]*ms.PostStar, error)
	GetUserPostStarCount(userID int64) (int64, error)
	GetUserPostCollection(postID, userID int64) (*ms.PostCollection, error)
	GetUserPostCollections(userID int64, offset, limit int) ([]*ms.PostCollection, error)
	GetUserPostCollectionCount(userID int64) (int64, error)
	GetPostAttatchmentBill(postID, userID int64) (*ms.PostAttachmentBill, error)
	GetPostContentsByIDs(ids []int64) ([]*ms.PostContent, error)
	GetPostContentByID(id int64) (*ms.PostContent, error)
	ListUserStarTweets(user *cs.VistUser, limit int, offset int) ([]*ms.PostStar, int64, error)
	ListUserMediaTweets(user *cs.VistUser, limit int, offset int) ([]*ms.Post, int64, error)
	ListUserCommentTweets(user *cs.VistUser, limit int, offset int) ([]*ms.Post, int64, error)
	ListUserTweets(userId int64, style uint8, justEssence bool, limit, offset int) ([]*ms.Post, int64, error)
	ListFollowingTweets(userId int64, limit, offset int) ([]*ms.Post, int64, error)
	ListIndexNewestTweets(limit, offset int) ([]*ms.Post, int64, error)
	ListIndexHotsTweets(limit, offset int) ([]*ms.Post, int64, error)
	ListSyncSearchTweets(limit, offset int) ([]*ms.Post, int64, error)
}

// TweetManageService 推文管理服务，包括创建/删除/更新推文
type TweetManageService interface {
	CreatePost(post *ms.Post) (*ms.Post, error)
	DeletePost(post *ms.Post) ([]string, error)
	LockPost(post *ms.Post) error
	StickPost(post *ms.Post) error
	HighlightPost(userId, postId int64) (int, error)
	VisiblePost(post *ms.Post, visibility cs.TweetVisibleType) error
	UpdatePost(post *ms.Post) error
	CreatePostStar(postID, userID int64) (*ms.PostStar, error)
	DeletePostStar(p *ms.PostStar) error
	CreatePostCollection(postID, userID int64) (*ms.PostCollection, error)
	DeletePostCollection(p *ms.PostCollection) error
	CreatePostContent(content *ms.PostContent) (*ms.PostContent, error)
	CreateAttachment(obj *ms.Attachment) (int64, error)
}

// TweetHelpService 推文辅助服务
type TweetHelpService interface {
	RevampPosts(posts []*ms.PostFormated) ([]*ms.PostFormated, error)
	MergePosts(posts []*ms.Post) ([]*ms.PostFormated, error)
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
