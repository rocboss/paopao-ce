// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/alimy/tryst/cfg"
	"github.com/bitbus/sqlx"
	"github.com/bitbus/sqlx/db"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pgc"
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

func (s *messageSrv) GetMessageByID(id int64) (*ms.Message, error) {
	return db.Get[ms.Message](s.q.GetMessageById, id)
}

func (s *messageSrv) ReadMessage(r *ms.Message) (err error) {
	_, err = s.q.ReadMessage.Exec(time.Now().Unix(), r.ID)
	return
}

func (s *messageSrv) GetMessages(userId int64, style cs.MessageStyle, limit, offset int) (res []*ms.MessageFormated, total int64, err error) {
	var messages []*ms.Message
	// 1动态，2评论，3回复，4私信，5好友申请，99系统通知
	switch style {
	case cs.StyleMsgSystem:
		if err = s.q.CountSystemMessages.Get(&total, userId); err == nil && total > 0 {
			err = s.q.GetSystemMessages.Select(&messages, userId, limit, offset)
		}
	case cs.StyleMsgWhisper:
		if err = s.q.CountWhisperMessages.Get(&total, userId, userId); err == nil && total > 0 {
			err = s.q.GetWhisperMessages.Select(&messages, userId, userId, limit, offset)
		}
	case cs.StyleMsgRequesting:
		if err = s.q.CountRequestingMessages.Get(&total, userId); err == nil && total > 0 {
			err = s.q.GetRequestingMessages.Select(&messages, userId, limit, offset)
		}
	case cs.StyleMsgUnread:
		if err = s.q.CountUnreadMessages.Get(&total, userId); err == nil && total > 0 {
			err = s.q.GetUnreadMessages.Select(&messages, userId, limit, offset)
		}
	case cs.StyleMsgAll:
		fallthrough
	default:
		if err = s.q.CountAllMessages.Get(&total, userId, userId); err == nil && total > 0 {
			err = s.q.GetAllMessages.Select(&messages, userId, userId, limit, offset)
		}
	}
	if err == nil {
		for _, message := range messages {
			res = append(res, message.Format())
		}
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
