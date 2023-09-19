// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/pkg/debug"
	"gorm.io/gorm"
)

var (
	_ core.TweetService       = (*tweetSrv)(nil)
	_ core.TweetManageService = (*tweetManageSrv)(nil)
	_ core.TweetHelpService   = (*tweetHelpSrv)(nil)
)

type tweetSrv struct {
	*pgxSrv
}

type tweetManageSrv struct {
	*pgxSrv
}

type tweetHelpSrv struct {
	*pgxSrv
}

// MergePosts post数据整合
func (s *tweetHelpSrv) MergePosts(posts []*ms.Post) ([]*ms.PostFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

// RevampPosts post数据整形修复
func (s *tweetHelpSrv) RevampPosts(posts []*ms.PostFormated) ([]*ms.PostFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetHelpSrv) RevampTweets(tweets cs.TweetList) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetHelpSrv) MergeTweets(tweets cs.TweetInfo) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetHelpSrv) getPostContentsByIDs(ids []int64) ([]*ms.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetHelpSrv) getUsersByIDs(ids []int64) ([]*ms.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) CreatePostCollection(postID, userID int64) (*ms.PostCollection, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) DeletePostCollection(p *ms.PostCollection) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) CreatePostContent(content *ms.PostContent) (*ms.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) CreatePost(post *ms.Post) (*ms.Post, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) DeletePost(post *ms.Post) ([]string, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) deleteCommentByPostId(db *gorm.DB, postId int64) ([]string, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) LockPost(post *ms.Post) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) StickPost(post *ms.Post) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) HighlightPost(userId, postId int64) (res int, err error) {
	// TODO
	debug.NotImplemented()
	return
}

func (s *tweetManageSrv) VisiblePost(post *ms.Post, visibility cs.TweetVisibleType) error {
	// TODO
	return cs.ErrNotImplemented
}

func (s *tweetManageSrv) UpdatePost(post *ms.Post) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) CreatePostStar(postID, userID int64) (*ms.PostStar, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) DeletePostStar(p *ms.PostStar) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) CreateAttachment(obj *ms.Attachment) (int64, error) {
	// TODO
	return 0, debug.ErrNotImplemented
}

func (s *tweetManageSrv) CreateTweet(userId int64, req *cs.NewTweetReq) (*cs.TweetItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetManageSrv) DeleteTweet(userId int64, tweetId int64) ([]string, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetManageSrv) LockTweet(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrv) StickTweet(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrv) VisibleTweet(userId int64, visibility cs.TweetVisibleType) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrv) CreateReaction(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrv) DeleteReaction(userId int64, reactionId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrv) CreateFavorite(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrv) DeleteFavorite(userId int64, favoriteId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetSrv) GetPostByID(id int64) (*ms.Post, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetPosts(conditions ms.ConditionsT, limit int, offset int) ([]*ms.Post, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetPostCount(conditions ms.ConditionsT) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetSrv) GetUserPostStar(postID, userID int64) (*ms.PostStar, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserPostStars(userID int64, limit int, offset int) ([]*ms.PostStar, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserPostStarCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetSrv) GetUserPostCollection(postID, userID int64) (*ms.PostCollection, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserPostCollections(userID int64, limit int, offset int) ([]*ms.PostCollection, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserPostCollectionCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetSrv) GetUserWalletBills(userID int64, limit int, offset int) ([]*ms.WalletStatement, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserWalletBillCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetSrv) GetPostAttatchmentBill(postID, userID int64) (*ms.PostAttachmentBill, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetPostContentsByIDs(ids []int64) ([]*ms.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetPostContentByID(id int64) (*ms.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) TweetInfoById(id int64) (*cs.TweetInfo, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrv) TweetItemById(id int64) (*cs.TweetItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrv) UserTweets(visitorId, userId int64) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrv) ReactionByTweetId(userId int64, tweetId int64) (*cs.ReactionItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrv) UserReactions(userId int64, limit int, offset int) (cs.ReactionList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrv) FavoriteByTweetId(userId int64, tweetId int64) (*cs.FavoriteItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrv) UserFavorites(userId int64, limit int, offset int) (cs.FavoriteList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrv) AttachmentByTweetId(userId int64, tweetId int64) (*cs.AttachmentBill, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrv) ListUserStarTweets(user *cs.VistUser, limit int, offset int) ([]*ms.PostStar, int64, error) {
	// TODO
	return nil, 0, debug.ErrNotImplemented
}

func (s *tweetSrv) ListUserMediaTweets(user *cs.VistUser, limit int, offset int) ([]*ms.Post, int64, error) {
	// TODO
	return nil, 0, debug.ErrNotImplemented
}

func (s *tweetSrv) ListUserCommentTweets(user *cs.VistUser, limit int, offset int) ([]*ms.Post, int64, error) {
	// TODO
	return nil, 0, debug.ErrNotImplemented
}

func (s *tweetSrv) ListUserTweets(userId int64, style uint8, justEssence bool, limit, offset int) ([]*ms.Post, int64, error) {
	// TODO
	return nil, 0, cs.ErrNotImplemented
}

func (s *tweetSrv) ListFollowingTweets(userId int64, limit, offset int) ([]*ms.Post, int64, error) {
	// TODO
	return nil, 0, cs.ErrNotImplemented
}

func (s *tweetSrv) ListIndexNewestTweets(limit, offset int) ([]*ms.Post, int64, error) {
	// TODO
	return nil, 0, cs.ErrNotImplemented
}

func (s *tweetSrv) ListIndexHotsTweets(limit, offset int) ([]*ms.Post, int64, error) {
	// TODO
	return nil, 0, cs.ErrNotImplemented
}

func (s *tweetSrv) ListSyncSearchTweets(limit, offset int) ([]*ms.Post, int64, error) {
	// TODO
	return nil, 0, cs.ErrNotImplemented
}

func newTweetService(db *pgx.Conn) core.TweetService {
	return &tweetSrv{
		pgxSrv: newPgxSrv(db),
	}
}

func newTweetManageService(db *pgx.Conn, cacheIndex core.CacheIndexService) core.TweetManageService {
	return &tweetManageSrv{
		pgxSrv: newPgxSrv(db),
	}
}

func newTweetHelpService(db *pgx.Conn) core.TweetHelpService {
	return &tweetHelpSrv{
		pgxSrv: newPgxSrv(db),
	}
}
