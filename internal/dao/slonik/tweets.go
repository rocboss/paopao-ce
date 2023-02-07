// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/pkg/debug"
	"gorm.io/gorm"
)

var (
	_ core.TweetService       = (*tweetServant)(nil)
	_ core.TweetManageService = (*tweetManageServant)(nil)
	_ core.TweetHelpService   = (*tweetHelpServant)(nil)
)

type tweetServant struct {
	*pgxServant
}

type tweetManageServant struct {
	*pgxServant
}

type tweetHelpServant struct {
	*pgxServant
}

// MergePosts post数据整合
func (s *tweetHelpServant) MergePosts(posts []*core.Post) ([]*core.PostFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

// RevampPosts post数据整形修复
func (s *tweetHelpServant) RevampPosts(posts []*core.PostFormated) ([]*core.PostFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetHelpServant) RevampTweets(tweets cs.TweetList) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetHelpServant) MergeTweets(tweets cs.TweetInfo) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetHelpServant) getPostContentsByIDs(ids []int64) ([]*core.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetHelpServant) getUsersByIDs(ids []int64) ([]*core.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageServant) CreatePostCollection(postID, userID int64) (*core.PostCollection, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageServant) DeletePostCollection(p *core.PostCollection) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageServant) CreatePostContent(content *core.PostContent) (*core.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageServant) CreatePost(post *core.Post) (*core.Post, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageServant) DeletePost(post *core.Post) ([]string, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageServant) deleteCommentByPostId(db *gorm.DB, postId int64) ([]string, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageServant) LockPost(post *core.Post) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageServant) StickPost(post *core.Post) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageServant) VisiblePost(post *core.Post, visibility core.PostVisibleT) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageServant) UpdatePost(post *core.Post) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageServant) CreatePostStar(postID, userID int64) (*core.PostStar, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageServant) DeletePostStar(p *core.PostStar) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageServant) CreateAttachment(obj *cs.Attachment) (int64, error) {
	// TODO
	return 0, debug.ErrNotImplemented
}

func (s *tweetManageServant) CreateTweet(userId int64, req *cs.NewTweetReq) (*cs.TweetItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetManageServant) DeleteTweet(userId int64, tweetId int64) ([]string, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetManageServant) LockTweet(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageServant) StickTweet(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageServant) VisibleTweet(userId int64, visibility cs.TweetVisibleType) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageServant) CreateReaction(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageServant) DeleteReaction(userId int64, reactionId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageServant) CreateFavorite(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageServant) DeleteFavorite(userId int64, favoriteId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetServant) GetPostByID(id int64) (*core.Post, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) GetPosts(conditions *core.ConditionsT, offset, limit int) ([]*core.Post, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) GetPostCount(conditions *core.ConditionsT) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetServant) GetUserPostStar(postID, userID int64) (*core.PostStar, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) GetUserPostStars(userID int64, offset, limit int) ([]*core.PostStar, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) GetUserPostStarCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetServant) GetUserPostCollection(postID, userID int64) (*core.PostCollection, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) GetUserPostCollections(userID int64, offset, limit int) ([]*core.PostCollection, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) GetUserPostCollectionCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetServant) GetUserWalletBills(userID int64, offset, limit int) ([]*core.WalletStatement, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) GetUserWalletBillCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetServant) GetPostAttatchmentBill(postID, userID int64) (*core.PostAttachmentBill, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) GetPostContentsByIDs(ids []int64) ([]*core.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) GetPostContentByID(id int64) (*core.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetServant) TweetInfoById(id int64) (*cs.TweetInfo, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetServant) TweetItemById(id int64) (*cs.TweetItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetServant) UserTweets(visitorId, userId int64) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetServant) ReactionByTweetId(userId int64, tweetId int64) (*cs.ReactionItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetServant) UserReactions(userId int64, offset int, limit int) (cs.ReactionList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetServant) FavoriteByTweetId(userId int64, tweetId int64) (*cs.FavoriteItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetServant) UserFavorites(userId int64, offset int, limit int) (cs.FavoriteList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetServant) AttachmentByTweetId(userId int64, tweetId int64) (*cs.AttachmentBill, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func newTweetService(db *pgx.Conn) core.TweetService {
	return &tweetServant{
		pgxServant: newPgxServant(db),
	}
}

func newTweetManageService(db *pgx.Conn, cacheIndex core.CacheIndexService) core.TweetManageService {
	return &tweetManageServant{
		pgxServant: newPgxServant(db),
	}
}

func newTweetHelpService(db *pgx.Conn) core.TweetHelpService {
	return &tweetHelpServant{
		pgxServant: newPgxServant(db),
	}
}
