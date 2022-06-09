package dao

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/model"
	"gorm.io/gorm"
)

func (d *dataServant) GetRechargeByID(id int64) (*model.WalletRecharge, error) {
	recharge := &model.WalletRecharge{
		Model: &model.Model{
			ID: id,
		},
	}

	return recharge.Get(d.engine)
}
func (d *dataServant) CreateRecharge(userId, amount int64) (*model.WalletRecharge, error) {
	recharge := &model.WalletRecharge{
		UserID: userId,
		Amount: amount,
	}

	return recharge.Create(d.engine)
}

func (d *dataServant) HandleRechargeSuccess(recharge *model.WalletRecharge, tradeNo string) error {
	user, _ := (&model.User{
		Model: &model.Model{
			ID: recharge.UserID,
		},
	}).Get(d.engine)

	return d.engine.Transaction(func(tx *gorm.DB) error {
		// 扣除金额
		if err := tx.Model(user).Update("balance", gorm.Expr("balance + ?", recharge.Amount)).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		// 新增账单
		if err := tx.Create(&model.WalletStatement{
			UserID:          user.ID,
			ChangeAmount:    recharge.Amount,
			BalanceSnapshot: user.Balance + recharge.Amount,
			Reason:          "用户充值",
		}).Error; err != nil {
			return err
		}

		// 标记为已付款
		if err := tx.Model(recharge).Updates(map[string]interface{}{
			"trade_no":     tradeNo,
			"trade_status": "TRADE_SUCCESS",
		}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (d *dataServant) HandlePostAttachmentBought(post *model.Post, user *model.User) error {
	return d.engine.Transaction(func(tx *gorm.DB) error {
		// 扣除金额
		if err := tx.Model(user).Update("balance", gorm.Expr("balance - ?", post.AttachmentPrice)).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		// 新增账单
		if err := tx.Create(&model.WalletStatement{
			PostID:          post.ID,
			UserID:          user.ID,
			ChangeAmount:    -post.AttachmentPrice,
			BalanceSnapshot: user.Balance - post.AttachmentPrice,
			Reason:          "购买附件支出",
		}).Error; err != nil {
			return err
		}

		// 新增附件购买记录
		if err := tx.Create(&model.PostAttachmentBill{
			PostID:     post.ID,
			UserID:     user.ID,
			PaidAmount: post.AttachmentPrice,
		}).Error; err != nil {
			return err
		}

		// 对附件主新增账单
		income := int64(float64(post.AttachmentPrice) * conf.AppSetting.AttachmentIncomeRate)
		if income > 0 {
			master := &model.User{
				Model: &model.Model{
					ID: post.UserID,
				},
			}
			master, _ = master.Get(d.engine)

			if err := tx.Model(master).Update("balance", gorm.Expr("balance + ?", income)).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}

			// 新增账单
			if err := tx.Create(&model.WalletStatement{
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
