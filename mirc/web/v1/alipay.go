package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntries(new(WebAlipayPub), new(WebAlipayPriv))
}

// // routeAlipay register Alipay feature releated route if needed
// func routeAlipay(public gin.IRoutes, authApi gin.IRoutes) {
// 	// 支付宝回调
// 	public.POST("/alipay/notify", api.AlipayNotify)

// 	// 用户充值
// 	authApi.POST("/user/recharge", api.GetUserRechargeLink)

// 	// 获取钱包余额
// 	authApi.GET("/user/recharge", api.GetUserRechargeResult)

// 	// 获取用户账单
// 	authApi.GET("/user/wallet/bills", api.GetUserWalletBills)
// }

type WebAlipayPub struct {
	Group Group `mir:"v1"`

	// AlipayNotify 支付宝回调
	AlipayNotify func(Post, web.AlipayNotifyReq) `mir:"/alipay/notify"`
}

type WebAlipayPriv struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`

	// 	UserRechargeLink 用户充值
	UserRechargeLink func(Post) `mir:"/user/recharge"`
	// UserRechargeResult 获取钱包余额
	UserRechargeResult func(Get) `mir:"/user/recharge"`
	// UserWalletBills 获取用户账单
	UserWalletBills func(Get, web.UserWalletBillsReq) web.UserWalletBillsResp `mir:"/user/wallet/bills"`
}
