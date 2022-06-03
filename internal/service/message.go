package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/errcode"
)

// MAX_WHISPER_NUM_DAILY 当日单用户私信总数限制（TODO 配置化、积分兑换等）
const MAX_WHISPER_NUM_DAILY = 20

type ReadMessageReq struct {
	ID int64 `json:"id" binding:"required"`
}
type WhisperReq struct {
	UserID  int64  `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreateWhisper 创建私信
func CreateWhisper(c *gin.Context, msg *model.Message) (*model.Message, error) {
	whisperKey := fmt.Sprintf("WhisperTimes:%d", msg.SenderUserID)

	// 今日频次限制
	if res, _ := global.Redis.Get(c, whisperKey).Result(); convert.StrTo(res).MustInt() >= MAX_WHISPER_NUM_DAILY {
		return nil, errcode.TooManyWhisperNum
	}

	// 创建私信
	msg, err := myDao.CreateMessage(msg)
	if err != nil {
		return nil, err
	}

	// 写入当日（自然日）计数缓存
	global.Redis.Incr(c, whisperKey).Result()

	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	global.Redis.Expire(c, whisperKey, endTime.Sub(currentTime))

	return msg, err
}

func GetUnreadCount(userID int64) (int64, error) {
	return myDao.GetUnreadCount(userID)
}
func ReadMessage(id, userID int64) error {
	// 获取message
	message, err := myDao.GetMessageByID(id)
	if err != nil {
		return err
	}

	if message.ReceiverUserID != userID {
		return errcode.NoPermission
	}

	// 已读消息
	return myDao.ReadMessage(message)
}

func GetMessages(userID int64, offset, limit int) ([]*model.MessageFormated, int64, error) {
	conditions := &model.ConditionsT{
		"receiver_user_id": userID,
		"ORDER":            "id DESC",
	}
	messages, err := myDao.GetMessages(conditions, offset, limit)

	for _, mf := range messages {

		if mf.SenderUserID > 0 {
			user, err := myDao.GetUserByID(mf.SenderUserID)
			if err == nil {
				mf.SenderUser = user.Format()
			}

		}

		if mf.PostID > 0 {
			post, err := GetPost(mf.PostID)
			if err == nil {
				mf.Post = post
				if mf.CommentID > 0 {
					comment, err := GetPostComment(mf.CommentID)

					if err == nil {
						mf.Comment = comment

						if mf.ReplyID > 0 {
							reply, err := GetPostCommentReply(mf.ReplyID)
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
	totalRows, _ := myDao.GetMessageCount(conditions)

	return messages, totalRows, nil
}
