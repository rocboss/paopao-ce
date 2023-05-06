// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"fmt"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/xerror"
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
	*base.DaoServant

	alipayClient *alipay.Client
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
		if err := s.Redis.SetRechargeStatus(req.Ctx, req.TradeNo); err == nil {
			recharge, err := s.Ds.GetRechargeByID(req.ID)
			if err != nil {
				logrus.Errorf("GetRechargeByID id:%d err: %s", req.ID, err)
				return _errRechargeNotifyError
			}
			if recharge.TradeStatus != "TRADE_SUCCESS" {
				// 标记为已付款
				err := s.Ds.HandleRechargeSuccess(recharge, req.TradeNo)
				defer s.Redis.DelRechargeStatus(req.Ctx, req.TradeNo)
				if err != nil {
					logrus.Errorf("HandleRechargeSuccess id:%d err: %s", req.ID, err)
					return _errRechargeNotifyError
				}
			}
		}
	}
	return nil
}

func (b *alipayPrivBinding) BindUserWalletBills(c *gin.Context) (*web.UserWalletBillsReq, mir.Error) {
	uid, ok := base.UserIdFrom(c)
	if !ok {
		return nil, xerror.UnauthorizedTokenError
	}
	page, pageSize := app.GetPageInfo(c)
	return &web.UserWalletBillsReq{
		UserId:   uid,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (b *alipayPrivBinding) BindUserRechargeLink(c *gin.Context) (*web.UserRechargeLinkReq, mir.Error) {
	v := &web.UserRechargeLinkReq{
		Host: c.Request.Host,
	}
	err := b.BindAny(c, v)
	return v, err
}

func (b *alipayPrivBinding) BindUserRechargeResult(c *gin.Context) (*web.UserRechargeResultReq, mir.Error) {
	uid, exist := base.UserIdFrom(c)
	if !exist {
		return nil, xerror.UnauthorizedTokenError
	}
	return &web.UserRechargeResultReq{
		UserId: uid,
		Id:     convert.StrTo(c.Query("id")).MustInt64(),
	}, nil
}

func (s *alipayPrivSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (s *alipayPrivSrv) UserWalletBills(req *web.UserWalletBillsReq) (*web.UserWalletBillsResp, mir.Error) {
	bills, err := s.Ds.GetUserWalletBills(req.UserId, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		logrus.Errorf("GetUserWalletBills err: %s", err)
		return nil, _errUserWalletBillsFailed
	}
	totalRows, err := s.Ds.GetUserWalletBillCount(req.UserId)
	if err != nil {
		logrus.Errorf("GetUserWalletBillCount err: %s", err)
		return nil, _errUserWalletBillsFailed
	}
	resp := base.PageRespFrom(bills, req.Page, req.PageSize, totalRows)
	return (*web.UserWalletBillsResp)(resp), nil
}

func (s *alipayPrivSrv) UserRechargeLink(req *web.UserRechargeLinkReq) (*web.UserRechargeLinkResp, mir.Error) {
	recharge, err := s.Ds.CreateRecharge(req.User.ID, req.Amount)
	if err != nil {
		logrus.Errorf("Ds.CreateRecharge err: %v", err)
		return nil, _errRechargeReqFail
	}
	p := alipay.TradePreCreate{}
	p.OutTradeNo = fmt.Sprintf("%d", recharge.ID)
	p.Subject = "PaoPao用户钱包充值"
	p.TotalAmount = fmt.Sprintf("%.2f", float64(recharge.Amount)/100.0)
	p.NotifyURL = "https://" + req.Host + "/v1/alipay/notify"
	rsp, err := s.alipayClient.TradePreCreate(p)
	if err != nil {
		logrus.Errorf("client.TradePreCreate err: %v\n", err)
		return nil, _errRechargeReqFail
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		return nil, _errRechargeReqFail
	}
	return &web.UserRechargeLinkResp{
		Id:  recharge.ID,
		Pay: rsp.Content.QRCode,
	}, nil
}

func (s *alipayPrivSrv) UserRechargeResult(req *web.UserRechargeResultReq) (*web.UserRechargeResultResp, mir.Error) {
	recharge, err := s.Ds.GetRechargeByID(req.Id)
	if err != nil {
		logrus.Errorf("Ds.GetRechargeByID err: %v", err)
		return nil, _errGetRechargeFailed
	}
	if recharge.UserID != req.UserId {
		logrus.Errorf("Ds.GetRechargeByID userId not equel recharge.UserID: %d req.UserId %d", recharge.UserID, req.UserId)
		return nil, _errGetRechargeFailed
	}
	return &web.UserRechargeResultResp{
		Id:     recharge.ID,
		Status: recharge.TradeStatus,
	}, nil
}

func newAlipayPubSrv(s *base.DaoServant) api.AlipayPub {
	return &alipayPubSrv{
		DaoServant: s,
	}
}

func newAlipayPubBinding(alipayClient *alipay.Client) api.AlipayPubBinding {
	return &alipayPubBinding{
		UnimplementedAlipayPubBinding: &api.UnimplementedAlipayPubBinding{
			BindAny: base.NewBindAnyFn(),
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

func newAlipayPrivSrv(s *base.DaoServant, client *alipay.Client) api.AlipayPriv {
	return &alipayPrivSrv{
		DaoServant:   s,
		alipayClient: client,
	}
}

func newAlipayPrivBinding() api.AlipayPrivBinding {
	return &alipayPrivBinding{
		UnimplementedAlipayPrivBinding: &api.UnimplementedAlipayPrivBinding{
			BindAny: base.NewBindAnyFn(),
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
