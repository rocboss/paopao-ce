package service

import (
	"time"

	"github.com/rocboss/paopao-api/global"
	"github.com/rocboss/paopao-api/internal/model"
	"github.com/rocboss/paopao-api/pkg/errcode"
)

type RechargeReq struct {
	Amount int64 `json:"amount" form:"amount" binding:"required"`
}

func (svc *Service) GetRechargeByID(id int64) (*model.WalletRecharge, error) {
	return svc.dao.GetRechargeByID(id)
}

func (svc *Service) CreateRecharge(userID, amount int64) (*model.WalletRecharge, error) {
	return svc.dao.CreateRecharge(userID, amount)
}

func (svc *Service) FinishRecharge(id int64, tradeNo string) error {
	if ok, _ := global.Redis.SetNX(svc.ctx, "PaoPaoRecharge:"+tradeNo, 1, time.Second*5).Result(); ok {
		recharge, err := svc.dao.GetRechargeByID(id)
		if err != nil {
			return err
		}

		if recharge.TradeStatus != "TRADE_SUCCESS" {

			// 标记为已付款
			err := svc.dao.HandleRechargeSuccess(recharge, tradeNo)
			defer global.Redis.Del(svc.ctx, "PaoPaoRecharge:"+tradeNo)

			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (svc *Service) BuyPostAttachment(post *model.Post, user *model.User) error {
	if user.Balance < post.AttachmentPrice {
		return errcode.InsuffientDownloadMoney
	}

	// 执行购买
	return svc.dao.HandlePostAttachmentBought(post, user)
}
