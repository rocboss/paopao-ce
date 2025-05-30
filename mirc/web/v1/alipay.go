package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// AlipayPub 支付宝相关不用授权的服务
type AlipayPub struct {
	Schema `mir:"v1"`

	// AlipayNotify 支付宝回调
	AlipayNotify func(Post, web.AlipayNotifyReq) `mir:"alipay/notify"`
}

// AlipayPub 支付宝相关授权的服务
type AlipayPriv struct {
	Schema `mir:"v1,chain"`

	// 	UserRechargeLink 用户充值
	UserRechargeLink func(Post, web.UserRechargeLinkReq) web.UserRechargeLinkResp `mir:"user/recharge"`

	// UserRechargeResult 获取充值结果
	UserRechargeResult func(Get, web.UserRechargeResultReq) web.UserRechargeResultResp `mir:"user/recharge"`

	// UserWalletBills 获取用户账单
	UserWalletBills func(Get, web.UserWalletBillsReq) web.UserWalletBillsResp `mir:"user/wallet/bills"`
}
