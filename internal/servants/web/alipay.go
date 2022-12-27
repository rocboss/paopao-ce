// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"time"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
)

var (
	_ api.AlipayPub         = (*alipayPubSrv)(nil)
	_ api.AlipayPubBinding  = (*alipayPubBinding)(nil)
	_ api.AlipayPubRender   = (*alipayPubRender)(nil)
	_ api.AlipayPriv        = (*alipayPrivSrv)(nil)
	_ api.AlipayPrivBinding = (*alipayPrivBinding)(nil)
	_ api.AlipayPrivRender  = (*alipayPrivRender)(nil)
)

type alipayPubSrv struct {
	api.UnimplementedAlipayPubServant
	*base.DaoServant
}

type alipayPubBinding struct {
	*api.UnimplementedAlipayPubBinding

	alipayClient *alipay.Client
}

type alipayPubRender struct {
	*api.UnimplementedAlipayPubRender
}

type alipayPrivSrv struct {
	api.UnimplementedAlipayPrivServant
}

type alipayPrivBinding struct {
	*api.UnimplementedAlipayPrivBinding
}

type alipayPrivRender struct {
	*api.UnimplementedAlipayPrivRender
}

func (b *alipayPubBinding) BindAlipayNotify(c *gin.Context) (*web.AlipayNotifyReq, mir.Error) {
	if err := c.Request.ParseForm(); err != nil {
		logrus.Errorf("parse form err: %s", err)
		return nil, _errRechargeNotifyError
	}
	noti, err := b.alipayClient.GetTradeNotification(c.Request)
	if err != nil {
		logrus.Errorf("alipayClient.GetTradeNotification err: %s form: %v", err, c.Request.Form)
		return nil, _errRechargeNotifyError
	}
	return &web.AlipayNotifyReq{
		Ctx:         c.Request.Context(),
		ID:          convert.StrTo(noti.OutTradeNo).MustInt64(),
		TradeNo:     noti.TradeNo,
		TradeStatus: noti.TradeStatus,
	}, nil
}

func (s *alipayPubSrv) AlipayNotify(req *web.AlipayNotifyReq) mir.Error {
	if req.TradeStatus == alipay.TradeStatusSuccess {
		if ok, _ := s.Redis.SetNX(req.Ctx, "PaoPaoRecharge:"+req.TradeNo, 1, time.Second*5).Result(); ok {
			recharge, err := s.Ds.GetRechargeByID(req.ID)
			if err != nil {
				logrus.Errorf("GetRechargeByID id:%d err: %s", req.ID, err)
				return _errRechargeNotifyError
			}
			if recharge.TradeStatus != "TRADE_SUCCESS" {
				// 标记为已付款
				err := s.Ds.HandleRechargeSuccess(recharge, req.TradeNo)
				defer s.Redis.Del(req.Ctx, "PaoPaoRecharge:"+req.TradeNo)
				if err != nil {
					logrus.Errorf("HandleRechargeSuccess id:%d err: %s", req.ID, err)
					return _errRechargeNotifyError
				}
			}
		}
	}
	return nil
}

func (s *alipayPrivSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func newAlipayPubSrv(s *base.DaoServant) api.AlipayPub {
	return &alipayPubSrv{
		DaoServant: s,
	}
}

func newAlipayPubBinding(alipayClient *alipay.Client) api.AlipayPubBinding {
	return &alipayPubBinding{
		UnimplementedAlipayPubBinding: &api.UnimplementedAlipayPubBinding{
			BindAny: base.BindAny,
		},
		alipayClient: alipayClient,
	}
}

func newAlipayPubRender() api.AlipayPubRender {
	return &alipayPubRender{
		UnimplementedAlipayPubRender: &api.UnimplementedAlipayPubRender{
			RenderAny: base.RenderAny,
		},
	}
}

func newAlipayPrivSrv() api.AlipayPriv {
	return &alipayPrivSrv{}
}

func newAlipayPrivBinding() api.AlipayPrivBinding {
	return &alipayPrivBinding{
		UnimplementedAlipayPrivBinding: &api.UnimplementedAlipayPrivBinding{
			BindAny: base.BindAny,
		},
	}
}

func newAlipayPrivRender() api.AlipayPrivRender {
	return &alipayPrivRender{
		UnimplementedAlipayPrivRender: &api.UnimplementedAlipayPrivRender{
			RenderAny: base.RenderAny,
		},
	}
}
