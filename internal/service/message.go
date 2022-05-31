package service

import (
	"fmt"
	"time"

	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/errcode"
)

// 当日单用户私信总数限制（TODO 配置化、积分兑换等）
const MAX_WHISPER_NUM_DAILY = 20

type ReadMessageReq struct {
	ID int64 `json:"id" binding:"required"`
}
type WhisperReq struct {
	UserID  int64  `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 创建私信
func (svc *Service) CreateWhisper(msg *model.Message) (*model.Message, error) {
	whisperKey := fmt.Sprintf("WhisperTimes:%d", msg.SenderUserID)

	// 今日频次限制
	if res, _ := global.Redis.Get(svc.ctx, whisperKey).Result(); convert.StrTo(res).MustInt() >= MAX_WHISPER_NUM_DAILY {
		return nil, errcode.TooManyWhisperNum
	}

	// 创建私信
	msg, err := svc.dao.CreateMessage(msg)
	if err != nil {
		return nil, err
	}

	// 写入当日（自然日）计数缓存
	global.Redis.Incr(svc.ctx, whisperKey).Result()

	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	global.Redis.Expire(svc.ctx, whisperKey, endTime.Sub(currentTime))

	return msg, err
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
