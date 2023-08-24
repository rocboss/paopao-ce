// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pgc"
)

var (
	_ core.ContactManageService = (*pgcContactManageSrv)(nil)
)

type pgcContactManageSrv struct {
	*contactManageSrv
	p *pgc.ContactManager
}

func (s *pgcContactManageSrv) pgcFetchOrNewContact(tx *sqlx.Tx, userId int64, friendId int64, status int8) (res *cs.Contact, err error) {
	if err = stmtGet(tx.Stmtx(s.q.GetContact), res, userId, friendId); err != nil {
		err = stmtGet(tx.Stmtx(s.p.CreateContact), res, userId, friendId, status, time.Now().Unix())
	}
	return
}
