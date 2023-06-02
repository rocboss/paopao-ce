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
	_ core.WalletService = (*walletSrv)(nil)
)

type walletSrv struct {
	*sqlxSrv
	q *cc.Wallet
}

func (s *walletSrv) GetRechargeByID(id int64) (*ms.WalletRecharge, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}
func (s *walletSrv) CreateRecharge(userId, amount int64) (*ms.WalletRecharge, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *walletSrv) GetUserWalletBills(userID int64, offset, limit int) ([]*ms.WalletStatement, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *walletSrv) GetUserWalletBillCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
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
