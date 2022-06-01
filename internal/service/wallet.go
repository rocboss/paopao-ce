package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/errcode"
)

type RechargeReq struct {
	Amount int64 `json:"amount" form:"amount" binding:"required"`
}

func GetRechargeByID(id int64) (*model.WalletRecharge, error) {
	return myDao.GetRechargeByID(id)
}

func CreateRecharge(userID, amount int64) (*model.WalletRecharge, error) {
	return myDao.CreateRecharge(userID, amount)
}

func FinishRecharge(ctx *gin.Context, id int64, tradeNo string) error {
	if ok, _ := global.Redis.SetNX(ctx, "PaoPaoRecharge:"+tradeNo, 1, time.Second*5).Result(); ok {
		recharge, err := myDao.GetRechargeByID(id)
		if err != nil {
			return err
		}

		if recharge.TradeStatus != "TRADE_SUCCESS" {

			// 标记为已付款
			err := myDao.HandleRechargeSuccess(recharge, tradeNo)
			defer global.Redis.Del(ctx, "PaoPaoRecharge:"+tradeNo)

			if err != nil {
				return err
			}
		}

	}

	return nil
}

func BuyPostAttachment(post *model.Post, user *model.User) error {
	if user.Balance < post.AttachmentPrice {
		return errcode.InsuffientDownloadMoney
	}

	// 执行购买
	return myDao.HandlePostAttachmentBought(post, user)
}
