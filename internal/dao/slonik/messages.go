// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.MessageService = (*messageSrv)(nil)
)

type messageSrv struct {
	*pgxSrv
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

func (s *messageSrv) GetMessages(userId int64, offset, limit int) ([]*ms.MessageFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *messageSrv) GetMessageCount(userId int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func newMessageService(db *pgx.Conn) core.MessageService {
	return &messageSrv{
		pgxSrv: newPgxSrv(db),
	}
}
