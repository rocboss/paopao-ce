package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

// WalletService wallet service interface
type WalletService interface {
	GetRechargeByID(id int64) (*model.WalletRecharge, error)
	CreateRecharge(userId, amount int64) (*model.WalletRecharge, error)
	HandleRechargeSuccess(recharge *model.WalletRecharge, tradeNo string) error
	HandlePostAttachmentBought(post *model.Post, user *model.User) error
}
