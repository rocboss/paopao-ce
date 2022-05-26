package service

import (
	"github.com/rocboss/paopao-api/internal/model"
	"github.com/rocboss/paopao-api/pkg/errcode"
)

type ReadMessageReq struct {
	ID int64 `json:"id" binding:"required"`
}

func (svc *Service) CreateMessage(msg *model.Message) (*model.Message, error) {
	return svc.dao.CreateMessage(msg)
}

func (svc *Service) GetUnreadCount(userID int64) (int64, error) {
	return svc.dao.GetUnreadCount(userID)
}
func (svc *Service) ReadMessage(id, userID int64) error {
	// 获取message
	message, err := svc.dao.GetMessageByID(id)
	if err != nil {
		return err
	}

	if message.ReceiverUserID != userID {
		return errcode.NoPermission
	}

	// 已读消息
	return svc.dao.ReadMessage(message)
}

func (svc *Service) GetMessages(userID int64, offset, limit int) ([]*model.MessageFormated, int64, error) {
	conditions := &model.ConditionsT{
		"receiver_user_id": userID,
		"ORDER":            "id DESC",
	}
	messages, err := svc.dao.GetMessages(conditions, offset, limit)

	for _, mf := range messages {

		if mf.SenderUserID > 0 {
			user, err := svc.dao.GetUserByID(mf.SenderUserID)
			if err == nil {
				mf.SenderUser = user.Format()
			}

		}

		if mf.PostID > 0 {
			post, err := svc.GetPost(mf.PostID)
			if err == nil {
				mf.Post = post
				if mf.CommentID > 0 {
					comment, err := svc.GetPostComment(mf.CommentID)

					if err == nil {
						mf.Comment = comment

						if mf.ReplyID > 0 {
							reply, err := svc.GetPostCommentReply(mf.ReplyID)
							if err == nil {
								mf.Reply = reply
							}
						}
					}
				}
			}
		}
	}

	if err != nil {
		return nil, 0, err
	}

	// 获取总量
	totalRows, _ := svc.dao.GetMessageCount(conditions)

	return messages, totalRows, nil
}
