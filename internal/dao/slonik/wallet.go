// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	dbr "github.com/rocboss/paopao-ce/internal/dao/slonik/ce/postgres"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.WalletService = (*walletServant)(nil)
)

type walletServant struct {
	db *pgx.Conn
	q  dbr.Querier
}

func (s *walletServant) GetRechargeByID(id int64) (*core.WalletRecharge, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}
func (s *walletServant) CreateRecharge(userId, amount int64) (*core.WalletRecharge, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *walletServant) GetUserWalletBills(userID int64, offset, limit int) ([]*core.WalletStatement, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *walletServant) GetUserWalletBillCount(userID int64) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *walletServant) HandleRechargeSuccess(recharge *core.WalletRecharge, tradeNo string) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *walletServant) HandlePostAttachmentBought(post *core.Post, user *core.User) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func newWalletService(db *pgx.Conn) core.WalletService {
	return &walletServant{
		db: db,
		q:  dbr.New(db),
	}
}
