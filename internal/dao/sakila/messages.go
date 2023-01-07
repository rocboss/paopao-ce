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
	_ core.MessageService = (*messageServant)(nil)
)

type messageServant struct {
	*sqlxServant
	stmtAddMsg  *sqlx.Stmt
	stmtGetMsg  *sqlx.Stmt
	stmtReadMsg *sqlx.Stmt
}

func (s *messageServant) CreateMessage(msg *core.Message) (*core.Message, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageServant) GetUnreadCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *messageServant) GetMessageByID(id int64) (*core.Message, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageServant) ReadMessage(message *core.Message) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *messageServant) GetMessages(conditions *core.ConditionsT, offset, limit int) ([]*core.MessageFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageServant) GetMessageCount(conditions *core.ConditionsT) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func newMessageService(db *sqlx.DB) core.MessageService {
	return &messageServant{
		sqlxServant: newSqlxServant(db),
		stmtAddMsg:  c(`SELECT * FROM @person WHERE first_name=?`),
		stmtGetMsg:  c(`SELECT * FROM @person WHERE first_name=?`),
		stmtReadMsg: c(`SELECT * FROM @person WHERE first_name=?`),
	}
}
