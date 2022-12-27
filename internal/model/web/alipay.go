// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"context"

	"github.com/smartwalle/alipay/v3"
)

type UserWalletBillsReq struct {
	*BaseInfo `json:"-" binding:"-"`
}

type UserWalletBillsResp struct {
	// TODO
}

type AlipayNotifyReq struct {
	Ctx         context.Context
	ID          int64
	TradeNo     string
	TradeStatus alipay.TradeStatus
}
