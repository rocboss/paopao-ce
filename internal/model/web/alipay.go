// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"context"

	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/smartwalle/alipay/v3"
)

type UserWalletBillsReq struct {
	UserId   int64
	Page     int
	PageSize int
}

type UserWalletBillsResp base.PageResp

type UserRechargeLinkReq struct {
	BaseInfo `json:"-" form:"-" binding:"-"`
	Host     string `json:"-" form:"-" binding:"-"`
	Amount   int64  `json:"amount" form:"amount" binding:"required"`
}

type UserRechargeLinkResp struct {
	Id  int64  `json:"id"`
	Pay string `json:"pay"`
}

type UserRechargeResultReq struct {
	UserId int64
	Id     int64
}

type UserRechargeResultResp struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

type AlipayNotifyReq struct {
	Ctx         context.Context
	ID          int64
	TradeNo     string
	TradeStatus alipay.TradeStatus
}
