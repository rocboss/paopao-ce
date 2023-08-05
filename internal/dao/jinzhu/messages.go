// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
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
	return (&dbr.Message{}).Count(s.db, &dbr.ConditionsT{
		"receiver_user_id": userID,
		"is_read":          dbr.MsgStatusUnread,
	})
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

func (s *messageSrv) GetMessages(conditions *ms.ConditionsT, offset, limit int) ([]*ms.MessageFormated, error) {
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

func (s *messageSrv) GetMessageCount(conditions *ms.ConditionsT) (int64, error) {
	return (&dbr.Message{}).Count(s.db, conditions)
}
