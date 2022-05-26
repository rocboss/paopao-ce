package dao

import "github.com/rocboss/paopao-api/internal/model"

func (d *Dao) GetComments(conditions *model.ConditionsT, offset, limit int) ([]*model.Comment, error) {
	return (&model.Comment{}).List(d.engine, conditions, offset, limit)
}

func (d *Dao) GetCommentByID(id int64) (*model.Comment, error) {
	comment := &model.Comment{
		Model: &model.Model{
			ID: id,
		},
	}
	return comment.Get(d.engine)
}

func (d *Dao) DeleteComment(comment *model.Comment) error {
	return comment.Delete(d.engine)
}

func (d *Dao) GetCommentCount(conditions *model.ConditionsT) (int64, error) {
	return (&model.Comment{}).Count(d.engine, conditions)
}

func (d *Dao) CreateComment(comment *model.Comment) (*model.Comment, error) {
	return comment.Create(d.engine)
}

func (d *Dao) CreateCommentReply(reply *model.CommentReply) (*model.CommentReply, error) {
	return reply.Create(d.engine)
}

func (d *Dao) GetCommentReplyByID(id int64) (*model.CommentReply, error) {
	reply := &model.CommentReply{
		Model: &model.Model{
			ID: id,
		},
	}
	return reply.Get(d.engine)
}

func (d *Dao) DeleteCommentReply(reply *model.CommentReply) error {
	return reply.Delete(d.engine)
}
