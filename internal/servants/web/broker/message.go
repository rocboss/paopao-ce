// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package broker

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core"
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
func CreateWhisper(c *gin.Context, msg *core.Message) (*core.Message, error) {
	// 今日频次限制
	if count, err := _redis.GetCountWhisper(c, msg.SenderUserID); err != nil || count >= MAX_WHISPER_NUM_DAILY {
		return nil, errcode.TooManyWhisperNum
	}

	// 创建私信
	msg, err := ds.CreateMessage(msg)
	if err != nil {
		return nil, err
	}

	// 写入当日（自然日）计数缓存
	_redis.IncrCountWhisper(c, msg.SenderUserID)
	return msg, err
}

func GetUnreadCount(userID int64) (int64, error) {
	return ds.GetUnreadCount(userID)
}

func ReadMessage(id, userID int64) error {
	// 获取message
	message, err := ds.GetMessageByID(id)
	if err != nil {
		return err
	}

	if message.ReceiverUserID != userID {
		return errcode.NoPermission
	}

	// 已读消息
	return ds.ReadMessage(message)
}

func GetMessages(userID int64, offset, limit int) ([]*core.MessageFormated, int64, error) {
	conditions := &core.ConditionsT{
		"receiver_user_id": userID,
		"ORDER":            "id DESC",
	}
	messages, err := ds.GetMessages(conditions, offset, limit)

	for _, mf := range messages {

		if mf.SenderUserID > 0 {
			user, err := ds.GetUserByID(mf.SenderUserID)
			if err == nil {
				mf.SenderUser = user.Format()
			}

		}

		// 好友申请消息不需要获取其他信息
		if mf.Type == core.MsgTypeRequestingFriend {
			continue
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
	totalRows, _ := ds.GetMessageCount(conditions)

	return messages, totalRows, nil
}
