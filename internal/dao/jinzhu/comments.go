// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.CommentService       = (*commentServant)(nil)
	_ core.CommentManageService = (*commentManageServant)(nil)
)

type commentServant struct {
	db *gorm.DB
}

type commentManageServant struct {
	db *gorm.DB
}

func newCommentService(db *gorm.DB) core.CommentService {
	return &commentServant{
		db: db,
	}
}

func newCommentManageService(db *gorm.DB) core.CommentManageService {
	return &commentManageServant{
		db: db,
	}
}

func (s *commentServant) GetComments(conditions *core.ConditionsT, offset, limit int) ([]*core.Comment, error) {
	return (&dbr.Comment{}).List(s.db, conditions, offset, limit)
}

func (s *commentServant) GetCommentByID(id int64) (*core.Comment, error) {
	comment := &dbr.Comment{
		Model: &dbr.Model{
			ID: id,
		},
	}
	return comment.Get(s.db)
}

func (s *commentServant) GetCommentReplyByID(id int64) (*core.CommentReply, error) {
	reply := &dbr.CommentReply{
		Model: &dbr.Model{
			ID: id,
		},
	}
	return reply.Get(s.db)
}

func (s *commentServant) GetCommentCount(conditions *core.ConditionsT) (int64, error) {
	return (&dbr.Comment{}).Count(s.db, conditions)
}

func (s *commentServant) GetCommentContentsByIDs(ids []int64) ([]*core.CommentContent, error) {
	commentContent := &dbr.CommentContent{}
	return commentContent.List(s.db, &dbr.ConditionsT{
		"comment_id IN ?": ids,
	}, 0, 0)
}

func (s *commentServant) GetCommentRepliesByID(ids []int64) ([]*core.CommentReplyFormated, error) {
	CommentReply := &dbr.CommentReply{}
	replies, err := CommentReply.List(s.db, &dbr.ConditionsT{
		"comment_id IN ?": ids,
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
	repliesFormated := []*core.CommentReplyFormated{}
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

func (s *commentManageServant) DeleteComment(comment *core.Comment) error {
	return comment.Delete(s.db)
}

func (s *commentManageServant) CreateComment(comment *core.Comment) (*core.Comment, error) {
	return comment.Create(s.db)
}

func (s *commentManageServant) CreateCommentReply(reply *core.CommentReply) (*core.CommentReply, error) {
	return reply.Create(s.db)
}

func (s *commentManageServant) DeleteCommentReply(reply *core.CommentReply) error {
	return reply.Delete(s.db)
}

func (s *commentManageServant) CreateCommentContent(content *core.CommentContent) (*core.CommentContent, error) {
	return content.Create(s.db)
}
