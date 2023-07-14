// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.WalletService = (*walletSrv)(nil)
)

type walletSrv struct {
	*sqlxSrv
	q *cc.Wallet
}

func (s *walletSrv) GetRechargeByID(id int64) (*ms.WalletRecharge, error) {
	res := &ms.WalletRecharge{}
	err := s.q.GetRechargeById.Get(res, id)
	return res, err
}

func (s *walletSrv) CreateRecharge(userId, amount int64) (*ms.WalletRecharge, error) {
	res, err := s.q.CreateRecharge.Exec(userId, amount)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &ms.WalletRecharge{
		Model: &dbr.Model{
			ID: id,
		},
		UserID: userId,
		Amount: amount,
	}, nil
}

func (s *walletSrv) GetUserWalletBills(userID int64, offset, limit int) ([]*ms.WalletStatement, error) {
	res := []*ms.WalletStatement{}
	err := s.q.GetUserWalletBills.Select(&res, userID, limit, offset)
	return res, err
}

func (s *walletSrv) GetUserWalletBillCount(userID int64) (res int64, err error) {
	err = s.q.GetUserWalletBillCount.Get(&res, userID)
	return
}

func (s *walletSrv) HandleRechargeSuccess(recharge *ms.WalletRecharge, tradeNo string) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *walletSrv) HandlePostAttachmentBought(post *ms.Post, user *ms.User) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func newWalletService(db *sqlx.DB) core.WalletService {
	return &walletSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildWallet),
	}
}
