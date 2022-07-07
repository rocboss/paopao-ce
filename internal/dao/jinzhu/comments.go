package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
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

func (s *commentServant) GetComments(conditions *model.ConditionsT, offset, limit int) ([]*model.Comment, error) {
	return (&model.Comment{}).List(s.db, conditions, offset, limit)
}

func (s *commentServant) GetCommentByID(id int64) (*model.Comment, error) {
	comment := &model.Comment{
		Model: &model.Model{
			ID: id,
		},
	}
	return comment.Get(s.db)
}

func (s *commentServant) GetCommentReplyByID(id int64) (*model.CommentReply, error) {
	reply := &model.CommentReply{
		Model: &model.Model{
			ID: id,
		},
	}
	return reply.Get(s.db)
}

func (s *commentServant) GetCommentCount(conditions *model.ConditionsT) (int64, error) {
	return (&model.Comment{}).Count(s.db, conditions)
}

func (s *commentServant) GetCommentContentsByIDs(ids []int64) ([]*model.CommentContent, error) {
	commentContent := &model.CommentContent{}
	return commentContent.List(s.db, &model.ConditionsT{
		"comment_id IN ?": ids,
	}, 0, 0)
}

func (s *commentServant) GetCommentRepliesByID(ids []int64) ([]*model.CommentReplyFormated, error) {
	CommentReply := &model.CommentReply{}
	replies, err := CommentReply.List(s.db, &model.ConditionsT{
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
	repliesFormated := []*model.CommentReplyFormated{}
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

func (s *commentManageServant) DeleteComment(comment *model.Comment) error {
	return comment.Delete(s.db)
}

func (s *commentManageServant) CreateComment(comment *model.Comment) (*model.Comment, error) {
	return comment.Create(s.db)
}

func (s *commentManageServant) CreateCommentReply(reply *model.CommentReply) (*model.CommentReply, error) {
	return reply.Create(s.db)
}

func (s *commentManageServant) DeleteCommentReply(reply *model.CommentReply) error {
	return reply.Delete(s.db)
}

func (s *commentManageServant) CreateCommentContent(content *model.CommentContent) (*model.CommentContent, error) {
	return content.Create(s.db)
}
