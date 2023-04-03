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
	_ core.WalletService = (*walletSrv)(nil)
)

type walletSrv struct {
	*pgxSrv
}

func (s *walletSrv) GetRechargeByID(id int64) (*core.WalletRecharge, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}
func (s *walletSrv) CreateRecharge(userId, amount int64) (*core.WalletRecharge, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *walletSrv) GetUserWalletBills(userID int64, offset, limit int) ([]*core.WalletStatement, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *walletSrv) GetUserWalletBillCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *walletSrv) HandleRechargeSuccess(recharge *core.WalletRecharge, tradeNo string) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *walletSrv) HandlePostAttachmentBought(post *core.Post, user *core.User) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func newWalletService(db *pgx.Conn) core.WalletService {
	return &walletSrv{
		pgxSrv: newPgxSrv(db),
	}
}
