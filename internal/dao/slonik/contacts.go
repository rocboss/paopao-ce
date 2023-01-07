// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.ContactManageService = (*contactManageServant)(nil)
)

type contactManageServant struct {
	*pgxServant
}

func (s *contactManageServant) RequestingFriend(userId int64, friendId int64, greetings string) (err error) {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *contactManageServant) AddFriend(userId int64, friendId int64) (err error) {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *contactManageServant) RejectFriend(userId int64, friendId int64) (err error) {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *contactManageServant) DeleteFriend(userId int64, friendId int64) (err error) {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *contactManageServant) GetContacts(userId int64, offset int, limit int) (*core.ContactList, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *contactManageServant) IsFriend(userId int64, friendId int64) bool {
	// TODO
	debug.NotImplemented()
	return false
}

func newContactManageService(db *pgx.Conn) core.ContactManageService {
	return &contactManageServant{
		pgxServant: newPgxServant(db),
	}
}
