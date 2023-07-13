// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"strings"
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/pkg/debug"
	"gorm.io/gorm"
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
	db *gorm.DB
}

type tweetManageSrv struct {
	cacheIndex core.CacheIndexService
	db         *gorm.DB
}

type tweetHelpSrv struct {
	db *gorm.DB
}

type tweetSrvA struct {
	db *gorm.DB
}

type tweetManageSrvA struct {
	db *gorm.DB
}

type tweetHelpSrvA struct {
	db *gorm.DB
}

func newTweetService(db *gorm.DB) core.TweetService {
	return &tweetSrv{
		db: db,
	}
}

func newTweetManageService(db *gorm.DB, cacheIndex core.CacheIndexService) core.TweetManageService {
	return &tweetManageSrv{
		cacheIndex: cacheIndex,
		db:         db,
	}
}

func newTweetHelpService(db *gorm.DB) core.TweetHelpService {
	return &tweetHelpSrv{
		db: db,
	}
}

func newTweetServantA(db *gorm.DB) core.TweetServantA {
	return &tweetSrvA{
		db: db,
	}
}

func newTweetManageServantA(db *gorm.DB) core.TweetManageServantA {
	return &tweetManageSrvA{
		db: db,
	}
}

func newTweetHelpServantA(db *gorm.DB) core.TweetHelpServantA {
	return &tweetHelpSrvA{
		db: db,
	}
}

// MergePosts post数据整合
func (s *tweetHelpSrv) MergePosts(posts []*ms.Post) ([]*ms.PostFormated, error) {
	postIds := make([]int64, 0, len(posts))
	userIds := make([]int64, 0, len(posts))
	for _, post := range posts {
		postIds = append(postIds, post.ID)
		userIds = append(userIds, post.UserID)
	}

	postContents, err := s.getPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}

	users, err := s.getUsersByIDs(userIds)
	if err != nil {
		return nil, err
	}

	userMap := make(map[int64]*dbr.UserFormated, len(users))
	for _, user := range users {
		userMap[user.ID] = user.Format()
	}

	contentMap := make(map[int64][]*dbr.PostContentFormated, len(postContents))
	for _, content := range postContents {
		contentMap[content.PostID] = append(contentMap[content.PostID], content.Format())
	}

	// 数据整合
	postsFormated := make([]*dbr.PostFormated, 0, len(posts))
	for _, post := range posts {
		postFormated := post.Format()
		postFormated.User = userMap[post.UserID]
		postFormated.Contents = contentMap[post.ID]
		postsFormated = append(postsFormated, postFormated)
	}
	return postsFormated, nil
}

// RevampPosts post数据整形修复
func (s *tweetHelpSrv) RevampPosts(posts []*ms.PostFormated) ([]*ms.PostFormated, error) {
	postIds := make([]int64, 0, len(posts))
	userIds := make([]int64, 0, len(posts))
	for _, post := range posts {
		postIds = append(postIds, post.ID)
		userIds = append(userIds, post.UserID)
	}

	postContents, err := s.getPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}

	users, err := s.getUsersByIDs(userIds)
	if err != nil {
		return nil, err
	}

	userMap := make(map[int64]*dbr.UserFormated, len(users))
	for _, user := range users {
		userMap[user.ID] = user.Format()
	}

	contentMap := make(map[int64][]*dbr.PostContentFormated, len(postContents))
	for _, content := range postContents {
		contentMap[content.PostID] = append(contentMap[content.PostID], content.Format())
	}

	// 数据整合
	for _, post := range posts {
		post.User = userMap[post.UserID]
		post.Contents = contentMap[post.ID]
	}
	return posts, nil
}

func (s *tweetHelpSrv) getPostContentsByIDs(ids []int64) ([]*dbr.PostContent, error) {
	return (&dbr.PostContent{}).List(s.db, &dbr.ConditionsT{
		"post_id IN ?": ids,
		"ORDER":        "sort ASC",
	}, 0, 0)
}

func (s *tweetHelpSrv) getUsersByIDs(ids []int64) ([]*dbr.User, error) {
	user := &dbr.User{}

	return user.List(s.db, &dbr.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}

func (s *tweetManageSrv) CreatePostCollection(postID, userID int64) (*ms.PostCollection, error) {
	collection := &dbr.PostCollection{
		PostID: postID,
		UserID: userID,
	}

	return collection.Create(s.db)
}

func (s *tweetManageSrv) DeletePostCollection(p *ms.PostCollection) error {
	return p.Delete(s.db)
}

func (s *tweetManageSrv) CreatePostContent(content *ms.PostContent) (*ms.PostContent, error) {
	return content.Create(s.db)
}

func (s *tweetManageSrv) CreateAttachment(obj *ms.Attachment) (int64, error) {
	attachment, err := obj.Create(s.db)
	return attachment.ID, err
}

func (s *tweetManageSrv) CreatePost(post *ms.Post) (*ms.Post, error) {
	post.LatestRepliedOn = time.Now().Unix()
	p, err := post.Create(s.db)
	if err != nil {
		return nil, err
	}
	s.cacheIndex.SendAction(core.IdxActCreatePost, post)
	return p, nil
}

func (s *tweetManageSrv) DeletePost(post *ms.Post) ([]string, error) {
	var mediaContents []string

	postId := post.ID
	postContent := &dbr.PostContent{}
	err := s.db.Transaction(
		func(tx *gorm.DB) error {
			if contents, err := postContent.MediaContentsByPostId(tx, postId); err == nil {
				mediaContents = contents
			} else {
				return err
			}

			// 删推文
			if err := post.Delete(tx); err != nil {
				return err
			}

			// 删内容
			if err := postContent.DeleteByPostId(tx, postId); err != nil {
				return err
			}

			// 删评论
			if contents, err := s.deleteCommentByPostId(tx, postId); err == nil {
				mediaContents = append(mediaContents, contents...)
			} else {
				return err
			}

			if tags := strings.Split(post.Tags, ","); len(tags) > 0 {
				// 删tag，宽松处理错误，有错误不会回滚
				deleteTags(tx, tags)
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	s.cacheIndex.SendAction(core.IdxActDeletePost, post)
	return mediaContents, nil
}

func (s *tweetManageSrv) deleteCommentByPostId(db *gorm.DB, postId int64) ([]string, error) {
	comment := &dbr.Comment{}
	commentContent := &dbr.CommentContent{}

	// 获取推文的所有评论id
	commentIds, err := comment.CommentIdsByPostId(db, postId)
	if err != nil {
		return nil, err
	}

	// 获取评论的媒体内容
	mediaContents, err := commentContent.MediaContentsByCommentId(db, commentIds)
	if err != nil {
		return nil, err
	}

	// 删评论
	if err = comment.DeleteByPostId(db, postId); err != nil {
		return nil, err
	}

	// 删评论内容
	if err = commentContent.DeleteByCommentIds(db, commentIds); err != nil {
		return nil, err
	}

	// 删评论的评论
	if err = (&dbr.CommentReply{}).DeleteByCommentIds(db, commentIds); err != nil {
		return nil, err
	}

	return mediaContents, nil
}

func (s *tweetManageSrv) LockPost(post *ms.Post) error {
	post.IsLock = 1 - post.IsLock
	return post.Update(s.db)
}

func (s *tweetManageSrv) StickPost(post *ms.Post) error {
	post.IsTop = 1 - post.IsTop
	if err := post.Update(s.db); err != nil {
		return err
	}
	s.cacheIndex.SendAction(core.IdxActStickPost, post)
	return nil
}

func (s *tweetManageSrv) VisiblePost(post *ms.Post, visibility core.PostVisibleT) error {
	oldVisibility := post.Visibility
	post.Visibility = visibility
	// TODO: 这个判断是否可以不要呢
	if oldVisibility == visibility {
		return nil
	}
	// 私密推文 特殊处理
	if visibility == dbr.PostVisitPrivate {
		// 强制取消置顶
		// TODO: 置顶推文用户是否有权设置成私密？ 后续完善
		post.IsTop = 0
	}
	db := s.db.Begin()
	err := post.Update(db)
	if err != nil {
		db.Rollback()
		return err
	}
	// tag处理
	tags := strings.Split(post.Tags, ",")
	// TODO: 暂时宽松不处理错误，这里或许可以有优化，后续完善
	if oldVisibility == dbr.PostVisitPrivate {
		// 从私密转为非私密才需要重新创建tag
		createTags(db, post.UserID, tags)
	} else if visibility == dbr.PostVisitPrivate {
		// 从非私密转为私密才需要删除tag
		deleteTags(db, tags)
	}
	db.Commit()
	s.cacheIndex.SendAction(core.IdxActVisiblePost, post)
	return nil
}

func (s *tweetManageSrv) UpdatePost(post *ms.Post) error {
	if err := post.Update(s.db); err != nil {
		return err
	}
	s.cacheIndex.SendAction(core.IdxActUpdatePost, post)
	return nil
}

func (s *tweetManageSrv) CreatePostStar(postID, userID int64) (*ms.PostStar, error) {
	star := &dbr.PostStar{
		PostID: postID,
		UserID: userID,
	}
	return star.Create(s.db)
}

func (s *tweetManageSrv) DeletePostStar(p *ms.PostStar) error {
	return p.Delete(s.db)
}

func (s *tweetSrv) GetPostByID(id int64) (*ms.Post, error) {
	post := &dbr.Post{
		Model: &dbr.Model{
			ID: id,
		},
	}
	return post.Get(s.db)
}

func (s *tweetSrv) GetPosts(conditions *ms.ConditionsT, offset, limit int) ([]*ms.Post, error) {
	return (&dbr.Post{}).List(s.db, conditions, offset, limit)
}

func (s *tweetSrv) GetPostCount(conditions *ms.ConditionsT) (int64, error) {
	return (&dbr.Post{}).Count(s.db, conditions)
}

func (s *tweetSrv) GetUserPostStar(postID, userID int64) (*ms.PostStar, error) {
	star := &dbr.PostStar{
		PostID: postID,
		UserID: userID,
	}
	return star.Get(s.db)
}

func (s *tweetSrv) GetUserPostStars(userID int64, offset, limit int) ([]*ms.PostStar, error) {
	star := &dbr.PostStar{
		UserID: userID,
	}

	return star.List(s.db, &dbr.ConditionsT{
		"ORDER": s.db.NamingStrategy.TableName("PostStar") + ".id DESC",
	}, offset, limit)
}

func (s *tweetSrv) GetUserPostStarCount(userID int64) (int64, error) {
	star := &dbr.PostStar{
		UserID: userID,
	}
	return star.Count(s.db, &dbr.ConditionsT{})
}

func (s *tweetSrv) GetUserPostCollection(postID, userID int64) (*ms.PostCollection, error) {
	star := &dbr.PostCollection{
		PostID: postID,
		UserID: userID,
	}
	return star.Get(s.db)
}

func (s *tweetSrv) GetUserPostCollections(userID int64, offset, limit int) ([]*ms.PostCollection, error) {
	collection := &dbr.PostCollection{
		UserID: userID,
	}

	return collection.List(s.db, &dbr.ConditionsT{
		"ORDER": s.db.NamingStrategy.TableName("PostCollection") + ".id DESC",
	}, offset, limit)
}

func (s *tweetSrv) GetUserPostCollectionCount(userID int64) (int64, error) {
	collection := &dbr.PostCollection{
		UserID: userID,
	}
	return collection.Count(s.db, &dbr.ConditionsT{})
}

func (s *tweetSrv) GetUserWalletBills(userID int64, offset, limit int) ([]*ms.WalletStatement, error) {
	statement := &dbr.WalletStatement{
		UserID: userID,
	}

	return statement.List(s.db, &dbr.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (s *tweetSrv) GetUserWalletBillCount(userID int64) (int64, error) {
	statement := &dbr.WalletStatement{
		UserID: userID,
	}
	return statement.Count(s.db, &dbr.ConditionsT{})
}

func (s *tweetSrv) GetPostAttatchmentBill(postID, userID int64) (*ms.PostAttachmentBill, error) {
	bill := &dbr.PostAttachmentBill{
		PostID: postID,
		UserID: userID,
	}

	return bill.Get(s.db)
}

func (s *tweetSrv) GetPostContentsByIDs(ids []int64) ([]*ms.PostContent, error) {
	return (&dbr.PostContent{}).List(s.db, &dbr.ConditionsT{
		"post_id IN ?": ids,
		"ORDER":        "sort ASC",
	}, 0, 0)
}

func (s *tweetSrv) GetPostContentByID(id int64) (*ms.PostContent, error) {
	return (&dbr.PostContent{
		Model: &dbr.Model{
			ID: id,
		},
	}).Get(s.db)
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

func (s *tweetSrvA) UserReactions(userId int64, offset int, limit int) (cs.ReactionList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrvA) FavoriteByTweetId(userId int64, tweetId int64) (*cs.FavoriteItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetSrvA) UserFavorites(userId int64, offset int, limit int) (cs.FavoriteList, error) {
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
