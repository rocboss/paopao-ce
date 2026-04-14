// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
)

type alipayPubSrv struct {
	api.UnimplementedAlipayPubServant
	*base.DaoServant
}

type alipayPrivSrv struct {
	api.UnimplementedAlipayPrivServant
	*base.DaoServant

	alipayClient *alipay.Client
}

func (s *alipayPubSrv) AlipayNotify(req *web.AlipayNotifyReq) error {
	if req.TradeStatus == alipay.TradeStatusSuccess {
		if err := s.Redis.SetRechargeStatus(req.Ctx, req.TradeNo); err == nil {
			recharge, err := s.Ds.GetRechargeByID(req.ID)
			if err != nil {
				logrus.Errorf("GetRechargeByID id:%d err: %s", req.ID, err)
				return web.ErrRechargeNotifyError
			}
			if recharge.TradeStatus != "TRADE_SUCCESS" {
				// 标记为已付款
				err := s.Ds.HandleRechargeSuccess(recharge, req.TradeNo)
				defer s.Redis.DelRechargeStatus(req.Ctx, req.TradeNo)
				if err != nil {
					logrus.Errorf("HandleRechargeSuccess id:%d err: %s", req.ID, err)
					return web.ErrRechargeNotifyError
				}
			}
		}
	}
	return nil
}

func (s *alipayPrivSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (s *alipayPrivSrv) UserWalletBills(req *web.UserWalletBillsReq) (*web.UserWalletBillsResp, error) {
	bills, err := s.Ds.GetUserWalletBills(req.UserId, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		logrus.Errorf("GetUserWalletBills err: %s", err)
		return nil, web.ErrUserWalletBillsFailed
	}
	totalRows, err := s.Ds.GetUserWalletBillCount(req.UserId)
	if err != nil {
		logrus.Errorf("GetUserWalletBillCount err: %s", err)
		return nil, web.ErrUserWalletBillsFailed
	}
	resp := base.PageRespFrom(bills, req.Page, req.PageSize, totalRows)
	return (*web.UserWalletBillsResp)(resp), nil
}

func (s *alipayPrivSrv) UserRechargeLink(req *web.UserRechargeLinkReq) (*web.UserRechargeLinkResp, error) {
	recharge, err := s.Ds.CreateRecharge(req.User.ID, req.Amount)
	if err != nil {
		logrus.Errorf("Ds.CreateRecharge err: %v", err)
		return nil, web.ErrRechargeReqFail
	}
	p := alipay.TradePreCreate{}
	p.OutTradeNo = fmt.Sprintf("%d", recharge.ID)
	p.Subject = "PaoPao用户钱包充值"
	p.TotalAmount = fmt.Sprintf("%.2f", float64(recharge.Amount)/100.0)
	p.NotifyURL = "https://" + req.Host + "/v1/alipay/notify"
	rsp, err := s.alipayClient.TradePreCreate(p)
	if err != nil {
		logrus.Errorf("client.TradePreCreate err: %v\n", err)
		return nil, web.ErrRechargeReqFail
	}
	if rsp.Code != alipay.CodeSuccess {
		return nil, web.ErrRechargeReqFail
	}
	return &web.UserRechargeLinkResp{
		Id:  recharge.ID,
		Pay: rsp.QRCode,
	}, nil
}

func (s *alipayPrivSrv) UserRechargeResult(req *web.UserRechargeResultReq) (*web.UserRechargeResultResp, error) {
	recharge, err := s.Ds.GetRechargeByID(req.Id)
	if err != nil {
		logrus.Errorf("Ds.GetRechargeByID err: %v", err)
		return nil, web.ErrGetRechargeFailed
	}
	if recharge.UserID != req.UserId {
		logrus.Errorf("Ds.GetRechargeByID userId not equel recharge.UserID: %d req.UserId %d", recharge.UserID, req.UserId)
		return nil, web.ErrGetRechargeFailed
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

func newAlipayPrivSrv(s *base.DaoServant, client *alipay.Client) api.AlipayPriv {
	return &alipayPrivSrv{
		DaoServant:   s,
		alipayClient: client,
	}
}
