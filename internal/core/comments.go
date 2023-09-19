// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
)

// CommentService 评论检索服务
type CommentService interface {
	GetComments(tweetId int64, style cs.StyleCommentType, limit int, offset int) ([]*ms.Comment, int64, error)
	GetCommentByID(id int64) (*ms.Comment, error)
	GetCommentReplyByID(id int64) (*ms.CommentReply, error)
	GetCommentContentsByIDs(ids []int64) ([]*ms.CommentContent, error)
	GetCommentRepliesByID(ids []int64) ([]*ms.CommentReplyFormated, error)
	GetCommentThumbsMap(userId int64, tweetId int64) (cs.CommentThumbsMap, cs.CommentThumbsMap, error)
}

// CommentManageService 评论管理服务
type CommentManageService interface {
	DeleteComment(comment *ms.Comment) error
	CreateComment(comment *ms.Comment) (*ms.Comment, error)
	CreateCommentReply(reply *ms.CommentReply) (*ms.CommentReply, error)
	DeleteCommentReply(reply *ms.CommentReply) error
	CreateCommentContent(content *ms.CommentContent) (*ms.CommentContent, error)
	ThumbsUpComment(userId int64, tweetId, commentId int64) error
	ThumbsDownComment(userId int64, tweetId, commentId int64) error
	ThumbsUpReply(userId int64, tweetId, commentId, replyId int64) error
	ThumbsDownReply(userId int64, tweetId, commentId, replyId int64) error
}
