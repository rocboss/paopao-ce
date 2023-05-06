// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"strings"
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.TweetService       = (*tweetServant)(nil)
	_ core.TweetManageService = (*tweetManageServant)(nil)
	_ core.TweetHelpService   = (*tweetHelpServant)(nil)
)

type tweetServant struct {
	db *gorm.DB
}

type tweetManageServant struct {
	cacheIndex core.CacheIndexService
	db         *gorm.DB
}

type tweetHelpServant struct {
	db *gorm.DB
}

func newTweetService(db *gorm.DB) core.TweetService {
	return &tweetServant{
		db: db,
	}
}

func newTweetManageService(db *gorm.DB, cacheIndex core.CacheIndexService) core.TweetManageService {
	return &tweetManageServant{
		cacheIndex: cacheIndex,
		db:         db,
	}
}

func newTweetHelpService(db *gorm.DB) core.TweetHelpService {
	return &tweetHelpServant{
		db: db,
	}
}

// MergePosts post数据整合
func (s *tweetHelpServant) MergePosts(posts []*core.Post) ([]*core.PostFormated, error) {
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
func (s *tweetHelpServant) RevampPosts(posts []*core.PostFormated) ([]*core.PostFormated, error) {
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

func (s *tweetHelpServant) getPostContentsByIDs(ids []int64) ([]*dbr.PostContent, error) {
	return (&dbr.PostContent{}).List(s.db, &dbr.ConditionsT{
		"post_id IN ?": ids,
		"ORDER":        "sort ASC",
	}, 0, 0)
}

func (s *tweetHelpServant) getUsersByIDs(ids []int64) ([]*dbr.User, error) {
	user := &dbr.User{}

	return user.List(s.db, &dbr.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}

func (s *tweetManageServant) CreatePostCollection(postID, userID int64) (*core.PostCollection, error) {
	collection := &dbr.PostCollection{
		PostID: postID,
		UserID: userID,
	}

	return collection.Create(s.db)
}

func (s *tweetManageServant) DeletePostCollection(p *core.PostCollection) error {
	return p.Delete(s.db)
}

func (s *tweetManageServant) CreatePostContent(content *core.PostContent) (*core.PostContent, error) {
	return content.Create(s.db)
}

func (s *tweetManageServant) CreateAttachment(attachment *core.Attachment) (*core.Attachment, error) {
	return attachment.Create(s.db)
}

func (s *tweetManageServant) CreatePost(post *core.Post) (*core.Post, error) {
	post.LatestRepliedOn = time.Now().Unix()
	p, err := post.Create(s.db)
	if err != nil {
		return nil, err
	}
	s.cacheIndex.SendAction(core.IdxActCreatePost, post)
	return p, nil
}

func (s *tweetManageServant) DeletePost(post *core.Post) ([]string, error) {
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

func (s *tweetManageServant) deleteCommentByPostId(db *gorm.DB, postId int64) ([]string, error) {
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

func (s *tweetManageServant) LockPost(post *core.Post) error {
	post.IsLock = 1 - post.IsLock
	return post.Update(s.db)
}

func (s *tweetManageServant) StickPost(post *core.Post) error {
	post.IsTop = 1 - post.IsTop
	if err := post.Update(s.db); err != nil {
		return err
	}
	s.cacheIndex.SendAction(core.IdxActStickPost, post)
	return nil
}

func (s *tweetManageServant) VisiblePost(post *core.Post, visibility core.PostVisibleT) error {
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
	for _, t := range tags {
		tag := &dbr.Tag{
			UserID: post.UserID,
			Tag:    t,
		}
		// TODO: 暂时宽松不处理错误，这里或许可以有优化，后续完善
		if oldVisibility == dbr.PostVisitPrivate {
			// 从私密转为非私密才需要重新创建tag
			createTag(db, tag)
		} else if visibility == dbr.PostVisitPrivate {
			// 从非私密转为私密才需要删除tag
			deleteTag(db, tag)
		}
	}
	db.Commit()
	s.cacheIndex.SendAction(core.IdxActVisiblePost, post)
	return nil
}

func (s *tweetManageServant) UpdatePost(post *core.Post) error {
	if err := post.Update(s.db); err != nil {
		return err
	}
	s.cacheIndex.SendAction(core.IdxActUpdatePost, post)
	return nil
}

func (s *tweetManageServant) CreatePostStar(postID, userID int64) (*core.PostStar, error) {
	star := &dbr.PostStar{
		PostID: postID,
		UserID: userID,
	}
	return star.Create(s.db)
}

func (s *tweetManageServant) DeletePostStar(p *core.PostStar) error {
	return p.Delete(s.db)
}

func (s *tweetServant) GetPostByID(id int64) (*core.Post, error) {
	post := &dbr.Post{
		Model: &dbr.Model{
			ID: id,
		},
	}
	return post.Get(s.db)
}

func (s *tweetServant) GetPosts(conditions *core.ConditionsT, offset, limit int) ([]*core.Post, error) {
	return (&dbr.Post{}).List(s.db, conditions, offset, limit)
}

func (s *tweetServant) GetPostCount(conditions *core.ConditionsT) (int64, error) {
	return (&dbr.Post{}).Count(s.db, conditions)
}

func (s *tweetServant) GetUserPostStar(postID, userID int64) (*core.PostStar, error) {
	star := &dbr.PostStar{
		PostID: postID,
		UserID: userID,
	}
	return star.Get(s.db)
}

func (s *tweetServant) GetUserPostStars(userID int64, offset, limit int) ([]*core.PostStar, error) {
	star := &dbr.PostStar{
		UserID: userID,
	}

	return star.List(s.db, &dbr.ConditionsT{
		"ORDER": s.db.NamingStrategy.TableName("PostStar") + ".id DESC",
	}, offset, limit)
}

func (s *tweetServant) GetUserPostStarCount(userID int64) (int64, error) {
	star := &dbr.PostStar{
		UserID: userID,
	}
	return star.Count(s.db, &dbr.ConditionsT{})
}

func (s *tweetServant) GetUserPostCollection(postID, userID int64) (*core.PostCollection, error) {
	star := &dbr.PostCollection{
		PostID: postID,
		UserID: userID,
	}
	return star.Get(s.db)
}

func (s *tweetServant) GetUserPostCollections(userID int64, offset, limit int) ([]*core.PostCollection, error) {
	collection := &dbr.PostCollection{
		UserID: userID,
	}

	return collection.List(s.db, &dbr.ConditionsT{
		"ORDER": s.db.NamingStrategy.TableName("PostCollection") + ".id DESC",
	}, offset, limit)
}

func (s *tweetServant) GetUserPostCollectionCount(userID int64) (int64, error) {
	collection := &dbr.PostCollection{
		UserID: userID,
	}
	return collection.Count(s.db, &dbr.ConditionsT{})
}

func (s *tweetServant) GetUserWalletBills(userID int64, offset, limit int) ([]*core.WalletStatement, error) {
	statement := &dbr.WalletStatement{
		UserID: userID,
	}

	return statement.List(s.db, &dbr.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (s *tweetServant) GetUserWalletBillCount(userID int64) (int64, error) {
	statement := &dbr.WalletStatement{
		UserID: userID,
	}
	return statement.Count(s.db, &dbr.ConditionsT{})
}

func (s *tweetServant) GetPostAttatchmentBill(postID, userID int64) (*core.PostAttachmentBill, error) {
	bill := &dbr.PostAttachmentBill{
		PostID: postID,
		UserID: userID,
	}

	return bill.Get(s.db)
}

func (s *tweetServant) GetPostContentsByIDs(ids []int64) ([]*core.PostContent, error) {
	return (&dbr.PostContent{}).List(s.db, &dbr.ConditionsT{
		"post_id IN ?": ids,
		"ORDER":        "sort ASC",
	}, 0, 0)
}

func (s *tweetServant) GetPostContentByID(id int64) (*core.PostContent, error) {
	return (&dbr.PostContent{
		Model: &dbr.Model{
			ID: id,
		},
	}).Get(s.db)
}
