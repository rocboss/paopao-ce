// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.TweetService       = (*tweetSrv)(nil)
	_ core.TweetManageService = (*tweetManageSrv)(nil)
	_ core.TweetHelpService   = (*tweetHelpSrv)(nil)

	_ core.TweetServantA       = (*tweetSrvA)(nil)
	_ core.TweetManageServantA = (*tweetManageSrvA)(nil)
	_ core.TweetHelpServantA   = (*tweetHelpSrvA)(nil)
)

type tweetSrv struct {
	*sqlxSrv
	q *cc.Tweet
}

type tweetManageSrv struct {
	*sqlxSrv
	cis core.CacheIndexService
	q   *cc.TweetManage
}

type tweetHelpSrv struct {
	*sqlxSrv
	q *cc.TweetHelp
}

type tweetSrvA struct {
	*sqlxSrv
	q *cc.TweetA
}

type tweetManageSrvA struct {
	*sqlxSrv
	cis core.CacheIndexService
	q   *cc.TweetManageA
}

type tweetHelpSrvA struct {
	*sqlxSrv
	q *cc.TweetHelpA
}

// MergePosts post数据整合
func (s *tweetHelpSrv) MergePosts(posts []*core.Post) ([]*core.PostFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

// RevampPosts post数据整形修复
func (s *tweetHelpSrv) RevampPosts(posts []*core.PostFormated) ([]*core.PostFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetHelpSrv) getPostContentsByIDs(ids []int64) ([]*dbr.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetHelpSrv) getUsersByIDs(ids []int64) ([]*dbr.User, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) CreatePostCollection(postID, userID int64) (*core.PostCollection, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) DeletePostCollection(p *core.PostCollection) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) CreatePostContent(content *core.PostContent) (*core.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) CreateAttachment(obj *cs.Attachment) (int64, error) {
	// TODO
	return 0, debug.ErrNotImplemented
}

func (s *tweetManageSrv) CreatePost(post *core.Post) (*core.Post, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) DeletePost(post *core.Post) ([]string, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) LockPost(post *core.Post) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) StickPost(post *core.Post) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) VisiblePost(post *core.Post, visibility core.PostVisibleT) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) UpdatePost(post *core.Post) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetManageSrv) CreatePostStar(postID, userID int64) (*core.PostStar, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) DeletePostStar(p *core.PostStar) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *tweetSrv) GetPostByID(id int64) (*core.Post, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetPosts(conditions *core.ConditionsT, offset, limit int) ([]*core.Post, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetPostCount(conditions *core.ConditionsT) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetSrv) GetUserPostStar(postID, userID int64) (*core.PostStar, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserPostStars(userID int64, offset, limit int) ([]*core.PostStar, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserPostStarCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetSrv) GetUserPostCollection(postID, userID int64) (*core.PostCollection, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserPostCollections(userID int64, offset, limit int) ([]*core.PostCollection, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserPostCollectionCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetSrv) GetUserWalletBills(userID int64, offset, limit int) ([]*core.WalletStatement, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetUserWalletBillCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetSrv) GetPostAttatchmentBill(postID, userID int64) (*core.PostAttachmentBill, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetPostContentsByIDs(ids []int64) ([]*core.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetPostContentByID(id int64) (*core.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrvA) TweetInfoById(id int64) (*cs.TweetInfo, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrvA) TweetItemById(id int64) (*cs.TweetItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrvA) UserTweets(visitorId, userId int64) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrvA) ReactionByTweetId(userId int64, tweetId int64) (*cs.ReactionItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrvA) UserReactions(userId int64, limit int, offset int) (cs.ReactionList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrvA) FavoriteByTweetId(userId int64, tweetId int64) (*cs.FavoriteItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrvA) UserFavorites(userId int64, limit int, offset int) (cs.FavoriteList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrvA) AttachmentByTweetId(userId int64, tweetId int64) (*cs.AttachmentBill, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetManageSrvA) CreateAttachment(obj *cs.Attachment) (int64, error) {
	// TODO
	return 0, debug.ErrNotImplemented
}

func (s *tweetManageSrvA) CreateTweet(userId int64, req *cs.NewTweetReq) (*cs.TweetItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetManageSrvA) DeleteTweet(userId int64, tweetId int64) ([]string, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetManageSrvA) LockTweet(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) StickTweet(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) VisibleTweet(userId int64, visibility cs.TweetVisibleType) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) CreateReaction(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) DeleteReaction(userId int64, reactionId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) CreateFavorite(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) DeleteFavorite(userId int64, favoriteId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetHelpSrvA) RevampTweets(tweets cs.TweetList) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetHelpSrvA) MergeTweets(tweets cs.TweetInfo) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func newTweetService(db *sqlx.DB) core.TweetService {
	return &tweetSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTweet),
	}
}

func newTweetManageService(db *sqlx.DB, cacheIndex core.CacheIndexService) core.TweetManageService {
	return &tweetManageSrv{
		sqlxSrv: newSqlxSrv(db),
		cis:     cacheIndex,
		q:       mustBuild(db, cc.BuildTweetManage),
	}
}

func newTweetHelpService(db *sqlx.DB) core.TweetHelpService {
	return &tweetHelpSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTweetHelp),
	}
}

func newTweetServantA(db *sqlx.DB) core.TweetServantA {
	return &tweetSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTweetA),
	}
}

func newTweetManageServantA(db *sqlx.DB, cacheIndex core.CacheIndexService) core.TweetManageServantA {
	return &tweetManageSrvA{
		sqlxSrv: newSqlxSrv(db),
		cis:     cacheIndex,
		q:       mustBuild(db, cc.BuildTweetManageA),
	}
}

func newTweetHelpServantA(db *sqlx.DB) core.TweetHelpServantA {
	return &tweetHelpSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTweetHelpA),
	}
}
