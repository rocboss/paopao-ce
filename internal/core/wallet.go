// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

type (
	WalletStatement = dbr.WalletStatement
	WalletRecharge  = dbr.WalletRecharge
)

// WalletService wallet service interface
type WalletService interface {
	GetUserWalletBills(userID int64, offset, limit int) ([]*WalletStatement, error)
	GetUserWalletBillCount(userID int64) (int64, error)
	GetRechargeByID(id int64) (*WalletRecharge, error)
	CreateRecharge(userId, amount int64) (*WalletRecharge, error)
	HandleRechargeSuccess(recharge *WalletRecharge, tradeNo string) error
	HandlePostAttachmentBought(post *Post, user *User) error
}
