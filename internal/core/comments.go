// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

type (
	Comment              = dbr.Comment
	CommentFormated      = dbr.CommentFormated
	CommentReply         = dbr.CommentReply
	CommentContent       = dbr.CommentContent
	CommentReplyFormated = dbr.CommentReplyFormated
)

// CommentService 评论检索服务
type CommentService interface {
	GetComments(conditions *ConditionsT, offset, limit int) ([]*Comment, error)
	GetCommentByID(id int64) (*Comment, error)
	GetCommentCount(conditions *ConditionsT) (int64, error)
	GetCommentReplyByID(id int64) (*CommentReply, error)
	GetCommentContentsByIDs(ids []int64) ([]*CommentContent, error)
	GetCommentRepliesByID(ids []int64) ([]*CommentReplyFormated, error)
	GetCommentThumbsMap(userId int64, tweetId int64) (cs.CommentThumbsMap, cs.CommentThumbsMap, error)
}

// CommentManageService 评论管理服务
type CommentManageService interface {
	DeleteComment(comment *Comment) error
	CreateComment(comment *Comment) (*Comment, error)
	CreateCommentReply(reply *CommentReply) (*CommentReply, error)
	DeleteCommentReply(reply *CommentReply) error
	CreateCommentContent(content *CommentContent) (*CommentContent, error)
	ThumbsUpComment(userId int64, tweetId, commentId int64) error
	ThumbsDownComment(userId int64, tweetId, commentId int64) error
	ThumbsUpReply(userId int64, tweetId, commentId, replyId int64) error
	ThumbsDownReply(userId int64, tweetId, commentId, replyId int64) error
}
