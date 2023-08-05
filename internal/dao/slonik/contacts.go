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
	_ core.ContactManageService = (*contactManageSrv)(nil)
)

type contactManageSrv struct {
	*pgxSrv
}

func (s *contactManageSrv) RequestingFriend(userId int64, friendId int64, greetings string) (err error) {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *contactManageSrv) AddFriend(userId int64, friendId int64) (err error) {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *contactManageSrv) RejectFriend(userId int64, friendId int64) (err error) {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *contactManageSrv) DeleteFriend(userId int64, friendId int64) (err error) {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *contactManageSrv) GetContacts(userId int64, offset int, limit int) (*ms.ContactList, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *contactManageSrv) IsFriend(userId int64, friendId int64) bool {
	// TODO
	debug.NotImplemented()
	return false
}

func newContactManageService(db *pgx.Conn) core.ContactManageService {
	return &contactManageSrv{
		pgxSrv: newPgxSrv(db),
	}
}
