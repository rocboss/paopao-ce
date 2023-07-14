// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.MessageService = (*messageSrv)(nil)
)

type messageSrv struct {
	*sqlxSrv
	q *cc.Message
}

func (s *messageSrv) CreateMessage(msg *ms.Message) (*ms.Message, error) {
	res, err := s.q.CreateMessage.Exec(msg)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	msg.Model = &dbr.Model{ID: id}
	return msg, nil
}

func (s *messageSrv) GetUnreadCount(userID int64) (res int64, err error) {
	err = s.q.GetUnreadCount.Get(&res)
	return
}

func (s *messageSrv) GetMessageByID(id int64) (*ms.Message, error) {
	res := &ms.Message{}
	err := s.q.GetMessageById.Get(res, id)
	return res, err
}

func (s *messageSrv) ReadMessage(message *ms.Message) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *messageSrv) GetMessages(r *ms.ConditionsT, offset, limit int) ([]*ms.MessageFormated, error) {
	res := []*ms.MessageFormated{}
	(*r)["limit"], (*r)["offset"] = limit, offset
	err := s.q.GetMessages.Select(&res, r)
	return res, err
}

func (s *messageSrv) GetMessageCount(r *ms.ConditionsT) (res int64, err error) {
	err = s.q.GetMessageCount.Get(&res, r)
	return
}

func newMessageService(db *sqlx.DB) core.MessageService {
	return &messageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildMessage),
	}
}
