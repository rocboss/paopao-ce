// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.MessageService = (*messageSrv)(nil)
)

type messageSrv struct {
	*sqlxSrv
	stmtAddMsg  *sqlx.Stmt
	stmtGetMsg  *sqlx.Stmt
	stmtReadMsg *sqlx.Stmt
}

func (s *messageSrv) CreateMessage(msg *core.Message) (*core.Message, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageSrv) GetUnreadCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *messageSrv) GetMessageByID(id int64) (*core.Message, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageSrv) ReadMessage(message *core.Message) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *messageSrv) GetMessages(conditions *core.ConditionsT, offset, limit int) ([]*core.MessageFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageSrv) GetMessageCount(conditions *core.ConditionsT) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func newMessageService(db *sqlx.DB) core.MessageService {
	return &messageSrv{
		sqlxSrv:     newSqlxSrv(db),
		stmtAddMsg:  c(`SELECT * FROM @person WHERE first_name=?`),
		stmtGetMsg:  c(`SELECT * FROM @person WHERE first_name=?`),
		stmtReadMsg: c(`SELECT * FROM @person WHERE first_name=?`),
	}
}
