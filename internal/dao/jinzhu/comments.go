// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/pkg/types"
	"gorm.io/gorm"
)

var (
	_ core.CommentService       = (*commentSrv)(nil)
	_ core.CommentManageService = (*commentManageSrv)(nil)
)

type commentSrv struct {
	db *gorm.DB
}

type commentManageSrv struct {
	db *gorm.DB
}

func newCommentService(db *gorm.DB) core.CommentService {
	return &commentSrv{
		db: db,
	}
}

func newCommentManageService(db *gorm.DB) core.CommentManageService {
	return &commentManageSrv{
		db: db,
	}
}

func (s *commentSrv) GetCommentThumbsMap(userId int64, tweetId int64) (cs.CommentThumbsMap, cs.CommentThumbsMap, error) {
	if userId < 0 {
		return nil, nil, nil
	}
	commentThumbsList := cs.CommentThumbsList{}
	err := s.db.Model(&dbr.TweetCommentThumbs{}).Where("user_id=? AND tweet_id=?", userId, tweetId).Find(&commentThumbsList).Error
	if err != nil {
		return nil, nil, err
	}
	commentThumbs, replyThumbs := make(cs.CommentThumbsMap), make(cs.CommentThumbsMap)
	for _, thumbs := range commentThumbsList {
		if thumbs.CommentType == 0 {
			commentThumbs[thumbs.CommentID] = thumbs
		} else {
			replyThumbs[thumbs.ReplyID] = thumbs
		}
	}
	return commentThumbs, replyThumbs, nil
}

func (s *commentSrv) GetComments(conditions *ms.ConditionsT, offset, limit int) ([]*ms.Comment, error) {
	return (&dbr.Comment{}).List(s.db, conditions, offset, limit)
}

func (s *commentSrv) GetCommentByID(id int64) (*ms.Comment, error) {
	comment := &dbr.Comment{
		Model: &dbr.Model{
			ID: id,
		},
	}
	return comment.Get(s.db)
}

func (s *commentSrv) GetCommentReplyByID(id int64) (*ms.CommentReply, error) {
	reply := &dbr.CommentReply{
		Model: &dbr.Model{
			ID: id,
		},
	}
	return reply.Get(s.db)
}

func (s *commentSrv) GetCommentCount(conditions *ms.ConditionsT) (int64, error) {
	return (&dbr.Comment{}).Count(s.db, conditions)
}

func (s *commentSrv) GetCommentContentsByIDs(ids []int64) ([]*ms.CommentContent, error) {
	commentContent := &dbr.CommentContent{}
	return commentContent.List(s.db, &dbr.ConditionsT{
		"comment_id IN ?": ids,
	}, 0, 0)
}

func (s *commentSrv) GetCommentRepliesByID(ids []int64) ([]*ms.CommentReplyFormated, error) {
	CommentReply := &dbr.CommentReply{}
	replies, err := CommentReply.List(s.db, &dbr.ConditionsT{
		"comment_id IN ?": ids,
		"ORDER":           "id ASC",
	}, 0, 0)

	if err != nil {
		return nil, err
	}

	userIds := []int64{}
	for _, reply := range replies {
		userIds = append(userIds, reply.UserID, reply.AtUserID)
	}

	users, err := getUsersByIDs(s.db, userIds)
	if err != nil {
		return nil, err
	}
	repliesFormated := []*ms.CommentReplyFormated{}
	for _, reply := range replies {
		replyFormated := reply.Format()
		for _, user := range users {
			if reply.UserID == user.ID {
				replyFormated.User = user.Format()
			}
			if reply.AtUserID == user.ID {
				replyFormated.AtUser = user.Format()
			}
		}

		repliesFormated = append(repliesFormated, replyFormated)
	}

	return repliesFormated, nil
}

func (s *commentManageSrv) DeleteComment(comment *ms.Comment) error {
	db := s.db.Begin()
	defer db.Rollback()

	err := comment.Delete(s.db)
	if err != nil {
		return err
	}
	err = db.Model(&dbr.TweetCommentThumbs{}).Where("user_id=? AND tweet_id=? AND comment_id=?", comment.UserID, comment.PostID, comment.ID).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
	if err != nil {
		return err
	}
	db.Commit()
	return nil
}

func (s *commentManageSrv) CreateComment(comment *ms.Comment) (*ms.Comment, error) {
	return comment.Create(s.db)
}

func (s *commentManageSrv) CreateCommentReply(reply *ms.CommentReply) (*ms.CommentReply, error) {
	return reply.Create(s.db)
}

func (s *commentManageSrv) DeleteCommentReply(reply *ms.CommentReply) (err error) {
	db := s.db.Begin()
	defer db.Rollback()

	err = reply.Delete(s.db)
	if err != nil {
		return
	}
	err = db.Model(&dbr.TweetCommentThumbs{}).
		Where("user_id=? AND comment_id=? AND reply_id=?", reply.UserID, reply.CommentID, reply.ID).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
	if err != nil {
		return
	}
	db.Commit()
	return
}

func (s *commentManageSrv) CreateCommentContent(content *ms.CommentContent) (*ms.CommentContent, error) {
	return content.Create(s.db)
}

func (s *commentManageSrv) ThumbsUpComment(userId int64, tweetId, commentId int64) error {
	db := s.db.Begin()
	defer db.Rollback()

	var (
		thumbsUpCount   int32 = 0
		thumbsDownCount int32 = 0
	)
	commentThumbs := &dbr.TweetCommentThumbs{}
	// 检查thumbs状态
	err := s.db.Where("user_id=? AND tweet_id=? AND comment_id=? AND comment_type=0", userId, tweetId, commentId).Take(commentThumbs).Error
	if err == nil {
		switch {
		case commentThumbs.IsThumbsUp == types.Yes && commentThumbs.IsThumbsDown == types.No:
			thumbsUpCount, thumbsDownCount = -1, 0
		case commentThumbs.IsThumbsUp == types.No && commentThumbs.IsThumbsDown == types.No:
			thumbsUpCount, thumbsDownCount = 1, 0
		default:
			thumbsUpCount, thumbsDownCount = 1, -1
			commentThumbs.IsThumbsDown = types.No
		}
		commentThumbs.IsThumbsUp = 1 - commentThumbs.IsThumbsUp
		commentThumbs.ModifiedOn = time.Now().Unix()
	} else {
		commentThumbs = &dbr.TweetCommentThumbs{
			UserID:       userId,
			TweetID:      tweetId,
			CommentID:    commentId,
			IsThumbsUp:   types.Yes,
			IsThumbsDown: types.No,
			CommentType:  0,
			Model: &dbr.Model{
				CreatedOn: time.Now().Unix(),
			},
		}
		thumbsUpCount, thumbsDownCount = 1, 0
	}
	// 更新thumbs状态
	if err = s.db.Save(commentThumbs).Error; err != nil {
		return err
	}
	// 更新thumbsUpCount
	if err = s.updateCommentThumbsUpCount(&dbr.Comment{}, commentId, thumbsUpCount, thumbsDownCount); err != nil {
		return err
	}
	db.Commit()
	return nil
}

func (s *commentManageSrv) ThumbsDownComment(userId int64, tweetId, commentId int64) error {
	db := s.db.Begin()
	defer db.Rollback()

	var (
		thumbsUpCount   int32 = 0
		thumbsDownCount int32 = 0
	)
	commentThumbs := &dbr.TweetCommentThumbs{}
	// 检查thumbs状态
	err := s.db.Where("user_id=? AND tweet_id=? AND comment_id=? AND comment_type=0", userId, tweetId, commentId).Take(commentThumbs).Error
	if err == nil {
		switch {
		case commentThumbs.IsThumbsDown == types.Yes:
			thumbsUpCount, thumbsDownCount = 0, -1
		case commentThumbs.IsThumbsDown == types.No && commentThumbs.IsThumbsUp == types.No:
			thumbsUpCount, thumbsDownCount = 0, 1
		default:
			thumbsUpCount, thumbsDownCount = -1, 1
			commentThumbs.IsThumbsUp = types.No

		}
		commentThumbs.IsThumbsDown = 1 - commentThumbs.IsThumbsDown
		commentThumbs.ModifiedOn = time.Now().Unix()
	} else {
		commentThumbs = &dbr.TweetCommentThumbs{
			UserID:       userId,
			TweetID:      tweetId,
			CommentID:    commentId,
			IsThumbsUp:   types.No,
			IsThumbsDown: types.Yes,
			CommentType:  0,
			Model: &dbr.Model{
				CreatedOn: time.Now().Unix(),
			},
		}
		thumbsUpCount, thumbsDownCount = 0, 1
	}
	// 更新thumbs状态
	if err = s.db.Save(commentThumbs).Error; err != nil {
		return err
	}
	// 更新thumbsUpCount
	if err = s.updateCommentThumbsUpCount(&dbr.Comment{}, commentId, thumbsUpCount, thumbsDownCount); err != nil {
		return err
	}
	db.Commit()
	return nil
}

func (s *commentManageSrv) ThumbsUpReply(userId int64, tweetId, commentId, replyId int64) error {
	db := s.db.Begin()
	defer db.Rollback()

	var (
		thumbsUpCount   int32 = 0
		thumbsDownCount int32 = 0
	)
	commentThumbs := &dbr.TweetCommentThumbs{}
	// 检查thumbs状态
	err := s.db.Where("user_id=? AND tweet_id=? AND comment_id=? AND reply_id=? AND comment_type=1", userId, tweetId, commentId, replyId).Take(commentThumbs).Error
	if err == nil {
		switch {
		case commentThumbs.IsThumbsUp == types.Yes:
			thumbsUpCount, thumbsDownCount = -1, 0
		case commentThumbs.IsThumbsUp == types.No && commentThumbs.IsThumbsDown == types.No:
			thumbsUpCount, thumbsDownCount = 1, 0
		default:
			thumbsUpCount, thumbsDownCount = 1, -1
			commentThumbs.IsThumbsDown = types.No
		}
		commentThumbs.IsThumbsUp = 1 - commentThumbs.IsThumbsUp
		commentThumbs.ModifiedOn = time.Now().Unix()
	} else {
		commentThumbs = &dbr.TweetCommentThumbs{
			UserID:       userId,
			TweetID:      tweetId,
			CommentID:    commentId,
			ReplyID:      replyId,
			IsThumbsUp:   types.Yes,
			IsThumbsDown: types.No,
			CommentType:  1,
			Model: &dbr.Model{
				CreatedOn: time.Now().Unix(),
			},
		}
		thumbsUpCount, thumbsDownCount = 1, 0
	}
	// 更新thumbs状态
	if err = s.db.Save(commentThumbs).Error; err != nil {
		return err
	}
	// 更新thumbsUpCount
	if err = s.updateCommentThumbsUpCount(&dbr.CommentReply{}, replyId, thumbsUpCount, thumbsDownCount); err != nil {
		return err
	}
	db.Commit()
	return nil
}

func (s *commentManageSrv) ThumbsDownReply(userId int64, tweetId, commentId, replyId int64) error {
	db := s.db.Begin()
	defer db.Rollback()

	var (
		thumbsUpCount   int32 = 0
		thumbsDownCount int32 = 0
	)
	commentThumbs := &dbr.TweetCommentThumbs{}
	// 检查thumbs状态
	err := s.db.Where("user_id=? AND tweet_id=? AND comment_id=? AND reply_id=? AND comment_type=1", userId, tweetId, commentId, replyId).Take(commentThumbs).Error
	if err == nil {
		switch {
		case commentThumbs.IsThumbsDown == types.Yes:
			thumbsUpCount, thumbsDownCount = 0, -1
		case commentThumbs.IsThumbsUp == types.No && commentThumbs.IsThumbsDown == types.No:
			thumbsUpCount, thumbsDownCount = 0, 1
		default:
			thumbsUpCount, thumbsDownCount = -1, 1
			commentThumbs.IsThumbsUp = types.No
		}
		commentThumbs.IsThumbsDown = 1 - commentThumbs.IsThumbsDown
		commentThumbs.ModifiedOn = time.Now().Unix()
	} else {
		commentThumbs = &dbr.TweetCommentThumbs{
			UserID:       userId,
			TweetID:      tweetId,
			CommentID:    commentId,
			ReplyID:      replyId,
			IsThumbsUp:   types.No,
			IsThumbsDown: types.Yes,
			CommentType:  1,
			Model: &dbr.Model{
				CreatedOn: time.Now().Unix(),
			},
		}
		thumbsUpCount, thumbsDownCount = 0, 1
	}
	// 更新thumbs状态
	if err = s.db.Save(commentThumbs).Error; err != nil {
		return err
	}
	// 更新thumbsUpCount
	if err = s.updateCommentThumbsUpCount(&dbr.CommentReply{}, replyId, thumbsUpCount, thumbsDownCount); err != nil {
		return err
	}
	db.Commit()
	return nil
}

func (s *commentManageSrv) updateCommentThumbsUpCount(obj any, id int64, thumbsUpCount, thumbsDownCount int32) error {
	updateColumns := make(map[string]any, 2)
	if thumbsUpCount == 1 {
		updateColumns["thumbs_up_count"] = gorm.Expr("thumbs_up_count + 1")
	} else if thumbsUpCount == -1 {
		updateColumns["thumbs_up_count"] = gorm.Expr("thumbs_up_count - 1")
	}
	if thumbsDownCount == 1 {
		updateColumns["thumbs_down_count"] = gorm.Expr("thumbs_down_count + 1")
	} else if thumbsDownCount == -1 {
		updateColumns["thumbs_down_count"] = gorm.Expr("thumbs_down_count - 1")
	}
	if len(updateColumns) > 0 {
		updateColumns["modified_on"] = time.Now().Unix()
		return s.db.Model(obj).Where("id=?", id).UpdateColumns(updateColumns).Error
	}
	return nil
}
