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
	_ core.ContactManageService = (*contactManageSrv)(nil)
)

type contactManageSrv struct {
	*sqlxSrv
	q *cc.ContactManager
}

func (s *contactManageSrv) RequestingFriend(userId int64, friendId int64, greetings string) (err error) {
	_, err = s.q.RejectFriend.Query(userId, friendId, greetings)
	return
}

func (s *contactManageSrv) AddFriend(userId int64, friendId int64) (err error) {
	_, err = s.q.AddFriend.Exec(userId, friendId)
	return
}

func (s *contactManageSrv) RejectFriend(userId int64, friendId int64) (err error) {
	_, err = s.q.RejectFriend.Exec(userId, friendId)
	return
}

func (s *contactManageSrv) DeleteFriend(userId int64, friendId int64) (err error) {
	_, err = s.q.DelFriend.Exec(userId, friendId)
	return
}

func (s *contactManageSrv) GetContacts(userId int64, offset int, limit int) (*ms.ContactList, error) {
	// TODO
	return nil, nil
}

func (s *contactManageSrv) IsFriend(userId int64, friendId int64) bool {
	// TODO
	debug.NotImplemented()
	return false
}

func newContactManageService(db *sqlx.DB) core.ContactManageService {
	return &contactManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildContactManager),
	}
}
