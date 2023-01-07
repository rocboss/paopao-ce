// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/pkg/debug"
	"gorm.io/gorm"
)

var (
	_ core.TweetService       = (*tweetServant)(nil)
	_ core.TweetManageService = (*tweetManageServant)(nil)
	_ core.TweetHelpService   = (*tweetHelpServant)(nil)
)

type tweetServant struct {
	*sqlxServant
	stmtGetTweet  *sqlx.Stmt
	stmtListTweet *sqlx.Stmt
	stmtListStar  *sqlx.Stmt
}

type tweetManageServant struct {
	*sqlxServant
	cacheIndex     core.CacheIndexService
	stmtAddTweet   *sqlx.Stmt
	stmtDelTweet   *sqlx.Stmt
	stmtStickTweet *sqlx.Stmt
}

type tweetHelpServant struct {
	*sqlxServant
	stmtAddTag  *sqlx.Stmt
	stmtDelTag  *sqlx.Stmt
	stmtListTag *sqlx.Stmt
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

func (s *tweetHelpServant) getPostContentsByIDs(ids []int64) ([]*dbr.PostContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetHelpServant) getUsersByIDs(ids []int64) ([]*dbr.User, error) {
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

func (s *tweetManageServant) CreateAttachment(attachment *core.Attachment) (*core.Attachment, error) {
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

func newTweetService(db *sqlx.DB) core.TweetService {
	return &tweetServant{
		sqlxServant:   newSqlxServant(db),
		stmtGetTweet:  c(`SELECT * FROM @person WHERE first_name=?`),
		stmtListTweet: c(`SELECT * FROM @person WHERE first_name=?`),
		stmtListStar:  c(`SELECT * FROM @person WHERE first_name=?`),
	}
}

func newTweetManageService(db *sqlx.DB, cacheIndex core.CacheIndexService) core.TweetManageService {
	return &tweetManageServant{
		sqlxServant:    newSqlxServant(db),
		cacheIndex:     cacheIndex,
		stmtAddTweet:   c(`SELECT * FROM @person WHERE first_name=?`),
		stmtDelTweet:   c(`SELECT * FROM @person WHERE first_name=?`),
		stmtStickTweet: c(`SELECT * FROM @person WHERE first_name=?`),
	}
}

func newTweetHelpService(db *sqlx.DB) core.TweetHelpService {
	return &tweetHelpServant{
		sqlxServant: newSqlxServant(db),
		stmtAddTag:  c(`SELECT * FROM @person WHERE first_name=?`),
		stmtDelTag:  c(`SELECT * FROM @person WHERE first_name=?`),
		stmtListTag: c(`SELECT * FROM @person WHERE first_name=?`),
	}
}
