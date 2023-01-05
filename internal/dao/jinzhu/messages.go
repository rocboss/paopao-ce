// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.MessageService = (*messageServant)(nil)
)

type messageServant struct {
	db *gorm.DB
}

func newMessageService(db *gorm.DB) core.MessageService {
	return &messageServant{
		db: db,
	}
}

func (s *messageServant) CreateMessage(msg *core.Message) (*core.Message, error) {
	return msg.Create(s.db)
}

func (s *messageServant) GetUnreadCount(userID int64) (int64, error) {
	return (&dbr.Message{}).Count(s.db, &dbr.ConditionsT{
		"receiver_user_id": userID,
		"is_read":          dbr.MsgStatusUnread,
	})
}

func (s *messageServant) GetMessageByID(id int64) (*core.Message, error) {
	return (&dbr.Message{
		Model: &dbr.Model{
			ID: id,
		},
	}).Get(s.db)
}

func (s *messageServant) ReadMessage(message *core.Message) error {
	message.IsRead = 1
	return message.Update(s.db)
}

func (s *messageServant) GetMessages(conditions *core.ConditionsT, offset, limit int) ([]*core.MessageFormated, error) {
	messages, err := (&dbr.Message{}).List(s.db, conditions, offset, limit)
	if err != nil {
		return nil, err
	}

	mfs := []*dbr.MessageFormated{}
	for _, message := range messages {
		mf := message.Format()
		mfs = append(mfs, mf)
	}

	return mfs, nil
}

func (s *messageServant) GetMessageCount(conditions *core.ConditionsT) (int64, error) {
	return (&dbr.Message{}).Count(s.db, conditions)
}
