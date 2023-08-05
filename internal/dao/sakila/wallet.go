// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
)

var (
	_ core.WalletService = (*walletSrv)(nil)
)

type walletSrv struct {
	*sqlxSrv
	q *cc.Wallet
}

func (s *walletSrv) GetRechargeByID(id int64) (res *ms.WalletRecharge, err error) {
	err = s.q.GetRechargeById.Get(res, id)
	return
}

func (s *walletSrv) CreateRecharge(userId, amount int64) (*ms.WalletRecharge, error) {
	now := time.Now().Unix()
	res, err := s.q.CreateRecharge.Exec(userId, amount, now)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &ms.WalletRecharge{
		Model: &dbr.Model{
			ID:        id,
			CreatedOn: now,
		},
		UserID: userId,
		Amount: amount,
	}, nil
}

func (s *walletSrv) GetUserWalletBills(userID int64, offset, limit int) (res []*ms.WalletStatement, err error) {
	err = s.q.GetUserWalletBills.Select(&res, userID, limit, offset)
	return
}

func (s *walletSrv) GetUserWalletBillCount(userID int64) (res int64, err error) {
	err = s.q.GetUserWalletBillCount.Get(&res, userID)
	return
}

func (s *walletSrv) HandleRechargeSuccess(r *ms.WalletRecharge, tradeNo string) error {
	return s.with(func(tx *sqlx.Tx) error {
		var oldBalance int64
		// 获取当前金额
		err := tx.Stmtx(s.q.GetUserBalance).Get(&oldBalance, r.UserID)
		if err != nil {
			return err
		}
		now := time.Now().Unix()
		// 更新金额
		if _, err := tx.Stmtx(s.q.AddUserBalance).Exec(r.Amount, now, r.UserID); err != nil {
			return err
		}
		// 新增账单
		args := []any{r.UserID, r.Amount, oldBalance + r.Amount, "用户充值", now}
		if _, err = tx.Stmtx(s.q.CreateWalletStatement).Exec(args...); err != nil {
			return err
		}
		// 标记为已付款
		if _, err = tx.Stmtx(s.q.MarkSuccessRecharge).Exec(tradeNo, now, r.ID); err != nil {
			return err
		}
		return nil
	})
}

func (s *walletSrv) HandlePostAttachmentBought(post *ms.Post, user *ms.User) error {
	return s.with(func(tx *sqlx.Tx) error {
		now := time.Now().Unix()
		// 扣除金额
		_, err := tx.Stmtx(s.q.MinusUserBalance).Exec(post.AttachmentPrice, now, user.ID)
		if err != nil {
			return err
		}
		// 新增账单
		args := []any{post.ID, user.ID, -post.AttachmentPrice, user.Balance - post.AttachmentPrice, "购买附件支出", now}
		if _, err = tx.Stmtx(s.q.NewPostBill).Exec(args...); err != nil {
			return err
		}
		// 新增附件购买记录
		args = []any{post.ID, user.ID, post.AttachmentPrice, now}
		if _, err = tx.Stmtx(s.q.NewPostAttachmentBill).Exec(args...); err != nil {
			return err
		}
		// 对附件主新增账单
		income := int64(float64(post.AttachmentPrice) * conf.AppSetting.AttachmentIncomeRate)
		if income > 0 {
			var oldBalance int64
			if err = tx.Stmtx(s.q.GetUserBalance).Get(&oldBalance, post.UserID); err != nil {
				return err
			}
			if _, err = tx.Stmtx(s.q.AddUserBalance).Exec(income, now, post.UserID); err != nil {
				return err
			}
			// 新增账单
			args = []any{post.ID, post.UserID, income, oldBalance + income, "出售附件收入", now}
			if _, err = tx.Stmtx(s.q.NewPostBill).Exec(args...); err != nil {
				return err
			}
		}
		return nil
	})
}

func newWalletService(db *sqlx.DB) core.WalletService {
	return &walletSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildWallet),
	}
}
