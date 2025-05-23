// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
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

func (r *AlipayNotifyReq) Bind(c *gin.Context) error {
	if err := c.Request.ParseForm(); err != nil {
		logrus.Errorf("parse form err: %s", err)
		return ErrRechargeNotifyError
	}
	noti, err := conf.MustAlipayClient().GetTradeNotification(c.Request)
	if err != nil {
		logrus.Errorf("alipayClient.GetTradeNotification err: %s form: %v", err, c.Request.Form)
		return ErrRechargeNotifyError
	}
	r.Ctx = c.Request.Context()
	r.ID = convert.StrTo(noti.OutTradeNo).MustInt64()
	r.TradeNo, r.TradeStatus = noti.TradeNo, noti.TradeStatus

	return nil
}

func (r *UserWalletBillsReq) Bind(c *gin.Context) error {
	uid, ok := base.UserIdFrom(c)
	if !ok {
		return xerror.UnauthorizedTokenError
	}
	r.UserId = uid
	r.Page, r.PageSize = app.GetPageInfo(c)
	return nil
}

func (r *UserRechargeLinkReq) Bind(c *gin.Context) error {
	r.Host = c.Request.Host
	return bindAny(c, r)
}

func (r *UserRechargeResultReq) Bind(c *gin.Context) error {
	uid, exist := base.UserIdFrom(c)
	if !exist {
		return xerror.UnauthorizedTokenError
	}
	r.UserId, r.Id = uid, convert.StrTo(c.Query("id")).MustInt64()
	return nil
}
