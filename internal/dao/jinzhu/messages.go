// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.MessageService = (*messageSrv)(nil)
)

type messageSrv struct {
	db *gorm.DB
}

func newMessageService(db *gorm.DB) core.MessageService {
	return &messageSrv{
		db: db,
	}
}

func (s *messageSrv) CreateMessage(msg *ms.Message) (*ms.Message, error) {
	return msg.Create(s.db)
}

func (s *messageSrv) GetUnreadCount(userID int64) (int64, error) {
	return (&dbr.Message{}).CountUnread(s.db, userID)
}

func (s *messageSrv) GetMessageByID(id int64) (*ms.Message, error) {
	return (&dbr.Message{
		Model: &dbr.Model{
			ID: id,
		},
	}).Get(s.db)
}

func (s *messageSrv) ReadMessage(message *ms.Message) error {
	message.IsRead = 1
	return message.Update(s.db)
}

func (s *messageSrv) GetMessages(userId int64, style cs.MessageStyle, limit int, offset int) (res []*ms.MessageFormated, total int64, err error) {
	var messages []*dbr.Message
	db := s.db.Table(_message_)
	// 1动态，2评论，3回复，4私信，5好友申请，99系统通知'
	switch style {
	case cs.StyleMsgSystem:
		db = db.Where("receiver_user_id=? AND type IN (1, 2, 3, 99)", userId)
	case cs.StyleMsgWhisper:
		db = db.Where("(receiver_user_id=? OR sender_user_id=?) AND type=4", userId, userId)
	case cs.StyleMsgRequesting:
		db = db.Where("receiver_user_id=? AND type=5", userId)
	case cs.StyleMsgUnread:
		db = db.Where("receiver_user_id=? AND is_read=0", userId)
	case cs.StyleMsgAll:
		fallthrough
	default:
		db = db.Where("receiver_user_id=? OR (sender_user_id=? AND type=4)", userId, userId)
	}
	if err = db.Count(&total).Error; err != nil || total == 0 {
		return
	}
	if offset >= 0 && limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	if err = db.Order("id DESC").Find(&messages).Error; err != nil {
		return
	}
	for _, message := range messages {
		res = append(res, message.Format())
	}
	return
}
