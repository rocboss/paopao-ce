// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/ms"
)

// WalletService wallet service interface
type WalletService interface {
	GetUserWalletBills(userID int64, offset, limit int) ([]*ms.WalletStatement, error)
	GetUserWalletBillCount(userID int64) (int64, error)
	GetRechargeByID(id int64) (*ms.WalletRecharge, error)
	CreateRecharge(userId, amount int64) (*ms.WalletRecharge, error)
	HandleRechargeSuccess(recharge *ms.WalletRecharge, tradeNo string) error
	HandlePostAttachmentBought(post *ms.Post, user *ms.User) error
}
