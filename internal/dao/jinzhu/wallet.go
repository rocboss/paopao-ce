// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.WalletService = (*walletServant)(nil)
)

type walletServant struct {
	db *gorm.DB
}

func newWalletService(db *gorm.DB) core.WalletService {
	return &walletServant{
		db: db,
	}
}

func (d *walletServant) GetRechargeByID(id int64) (*core.WalletRecharge, error) {
	recharge := &dbr.WalletRecharge{
		Model: &dbr.Model{
			ID: id,
		},
	}

	return recharge.Get(d.db)
}
func (d *walletServant) CreateRecharge(userId, amount int64) (*core.WalletRecharge, error) {
	recharge := &dbr.WalletRecharge{
		UserID: userId,
		Amount: amount,
	}

	return recharge.Create(d.db)
}

func (d *walletServant) GetUserWalletBills(userID int64, offset, limit int) ([]*core.WalletStatement, error) {
	statement := &dbr.WalletStatement{
		UserID: userID,
	}

	return statement.List(d.db, &dbr.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (d *walletServant) GetUserWalletBillCount(userID int64) (int64, error) {
	statement := &dbr.WalletStatement{
		UserID: userID,
	}
	return statement.Count(d.db, &dbr.ConditionsT{})
}

func (d *walletServant) HandleRechargeSuccess(recharge *core.WalletRecharge, tradeNo string) error {
	user, _ := (&dbr.User{
		Model: &dbr.Model{
			ID: recharge.UserID,
		},
	}).Get(d.db)

	return d.db.Transaction(func(tx *gorm.DB) error {
		// 扣除金额
		if err := tx.Model(user).Update("balance", gorm.Expr("balance + ?", recharge.Amount)).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		// 新增账单
		if err := tx.Create(&dbr.WalletStatement{
			UserID:          user.ID,
			ChangeAmount:    recharge.Amount,
			BalanceSnapshot: user.Balance + recharge.Amount,
			Reason:          "用户充值",
		}).Error; err != nil {
			return err
		}

		// 标记为已付款
		if err := tx.Model(recharge).Updates(map[string]any{
			"trade_no":     tradeNo,
			"trade_status": "TRADE_SUCCESS",
		}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (d *walletServant) HandlePostAttachmentBought(post *core.Post, user *core.User) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		// 扣除金额
		if err := tx.Model(user).Update("balance", gorm.Expr("balance - ?", post.AttachmentPrice)).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		// 新增账单
		if err := tx.Create(&dbr.WalletStatement{
			PostID:          post.ID,
			UserID:          user.ID,
			ChangeAmount:    -post.AttachmentPrice,
			BalanceSnapshot: user.Balance - post.AttachmentPrice,
			Reason:          "购买附件支出",
		}).Error; err != nil {
			return err
		}

		// 新增附件购买记录
		if err := tx.Create(&dbr.PostAttachmentBill{
			PostID:     post.ID,
			UserID:     user.ID,
			PaidAmount: post.AttachmentPrice,
		}).Error; err != nil {
			return err
		}

		// 对附件主新增账单
		income := int64(float64(post.AttachmentPrice) * conf.AppSetting.AttachmentIncomeRate)
		if income > 0 {
			master := &dbr.User{
				Model: &dbr.Model{
					ID: post.UserID,
				},
			}
			master, _ = master.Get(d.db)

			if err := tx.Model(master).Update("balance", gorm.Expr("balance + ?", income)).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}

			// 新增账单
			if err := tx.Create(&dbr.WalletStatement{
				PostID:          post.ID,
				UserID:          master.ID,
				ChangeAmount:    income,
				BalanceSnapshot: master.Balance + income,
				Reason:          "出售附件收入",
			}).Error; err != nil {
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})
}
