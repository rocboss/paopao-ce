package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

// CommentService 评论检索服务
type CommentService interface {
	GetComments(conditions *model.ConditionsT, offset, limit int) ([]*model.Comment, error)
	GetCommentByID(id int64) (*model.Comment, error)
	GetCommentCount(conditions *model.ConditionsT) (int64, error)
	GetCommentReplyByID(id int64) (*model.CommentReply, error)
	GetCommentContentsByIDs(ids []int64) ([]*model.CommentContent, error)
	GetCommentRepliesByID(ids []int64) ([]*model.CommentReplyFormated, error)
}

// CommentManageService 评论管理服务
type CommentManageService interface {
	DeleteComment(comment *model.Comment) error
	CreateComment(comment *model.Comment) (*model.Comment, error)
	CreateCommentReply(reply *model.CommentReply) (*model.CommentReply, error)
	DeleteCommentReply(reply *model.CommentReply) error
	CreateCommentContent(content *model.CommentContent) (*model.CommentContent, error)
}
