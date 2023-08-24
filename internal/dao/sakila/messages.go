// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/alimy/cfg"
	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pgc"
	"github.com/sirupsen/logrus"
)

var (
	_ core.MessageService = (*messageSrv)(nil)
)

type messageSrv struct {
	*sqlxSrv
	q *cc.Message
}

func (s *messageSrv) CreateMessage(r *ms.Message) (*ms.Message, error) {
	r.Model = &ms.Model{
		CreatedOn: time.Now().Unix(),
	}
	res, err := s.q.CreateMessage.Exec(r)
	if err != nil {
		return nil, err
	}
	r.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *messageSrv) GetUnreadCount(userID int64) (res int64, err error) {
	err = s.q.GetUnreadCount.Get(&res, userID)
	return
}

func (s *messageSrv) GetMessageByID(id int64) (res *ms.Message, err error) {
	err = stmtGet(s.q.GetMessageById, res, id)
	return
}

func (s *messageSrv) ReadMessage(r *ms.Message) (err error) {
	_, err = s.q.ReadMessage.Exec(time.Now().Unix(), r.ID)
	return
}

func (s *messageSrv) GetMessages(r *ms.ConditionsT, offset, limit int) ([]*ms.MessageFormated, error) {
	var messages []*ms.Message
	(*r)["limit"], (*r)["offset"] = limit, offset
	if err := s.q.GetMessages.Select(&messages, *r); err != nil {
		return nil, err
	}
	mfs := make([]*ms.MessageFormated, 0, len(messages))
	for _, message := range messages {
		mf := message.Format()
		mfs = append(mfs, mf)
	}
	return mfs, nil
}

func (s *messageSrv) GetMessageCount(r *ms.ConditionsT) (res int64, err error) {
	if err = s.q.GetMessageCount.Get(&res, *r); err != nil {
		logrus.Errorf("get message count error: %s", err)
	}
	return
}

func newMessageService(db *sqlx.DB) (s core.MessageService) {
	ms := &messageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildMessage),
	}
	s = ms
	if cfg.Any("PostgreSQL", "PgSQL", "Postgres") {
		s = &pgcMessageSrv{
			messageSrv: ms,
			p:          pgcBuild(db, pgc.BuildMessage),
		}
	}
	return
}
