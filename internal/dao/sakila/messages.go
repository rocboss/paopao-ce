// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
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
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageSrv) GetUnreadCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *messageSrv) GetMessageByID(id int64) (*ms.Message, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageSrv) ReadMessage(message *ms.Message) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *messageSrv) GetMessages(conditions *ms.ConditionsT, offset, limit int) ([]*ms.MessageFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageSrv) GetMessageCount(conditions *ms.ConditionsT) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func newMessageService(db *sqlx.DB) core.MessageService {
	return &messageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildMessage),
	}
}
