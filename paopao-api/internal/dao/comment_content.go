package dao

import "github.com/rocboss/paopao-api/internal/model"

func (d *Dao) GetCommentContentsByIDs(ids []int64) ([]*model.CommentContent, error) {
	commentContent := &model.CommentContent{}
	return commentContent.List(d.engine, &model.ConditionsT{
		"comment_id IN ?": ids,
	}, 0, 0)
}
func (d *Dao) GetCommentRepliesByID(ids []int64) ([]*model.CommentReplyFormated, error) {
	CommentReply := &model.CommentReply{}
	replies, err := CommentReply.List(d.engine, &model.ConditionsT{
		"comment_id IN ?": ids,
	}, 0, 0)

	if err != nil {
		return nil, err
	}

	userIds := []int64{}
	for _, reply := range replies {
		userIds = append(userIds, reply.UserID, reply.AtUserID)
	}

	users, err := d.GetUsersByIDs(userIds)
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

func (d *Dao) CreateCommentContent(content *model.CommentContent) (*model.CommentContent, error) {
	return content.Create(d.engine)
}
