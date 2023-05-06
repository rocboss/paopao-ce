// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"sync"

	"github.com/alimy/cfg"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
)

var (
	_enablePhoneVerify    bool
	_disallowUserRegister bool
	_onceInitial          sync.Once
)

// RouteWeb register web route
func RouteWeb(e *gin.Engine) {
	lazyInitial()
	oss := dao.ObjectStorageService()
	ds := base.NewDaoServant()
	// aways register servants
	api.RegisterAdminServant(e, newAdminSrv(ds), newAdminBinding(), newAdminRender())
	api.RegisterCoreServant(e, newCoreSrv(ds, oss), newCoreBinding(), newCoreRender())
	api.RegisterLooseServant(e, newLooseSrv(ds), newLooseBinding(), newLooseRender())
	api.RegisterPrivServant(e, newPrivSrv(ds, oss), newPrivBinding(), newPrivRender())
	api.RegisterPubServant(e, newPubSrv(ds), newPubBinding(), newPubRender())
	// regster servants if needed by configure
	cfg.In(cfg.Actions{
		"Alipay": func() {
			client := mustAlipayClient()
			api.RegisterAlipayPubServant(e, newAlipayPubSrv(ds), newAlipayPubBinding(client), newAlipayPubRender())
			api.RegisterAlipayPrivServant(e, newAlipayPrivSrv(ds, client), newAlipayPrivBinding(), newAlipayPrivRender())
		},
		"Followship": func() {
			api.RegisterFollowshipServant(e, newFollowshipSrv(ds), newFollowshipBinding(), newFollowshipRender())
		},
		"Friendship": func() {
			api.RegisterFriendshipServant(e, newFriendshipSrv(ds), newFriendshipBinding(), newFriendshipRender())
		},
	})
}

func mustAlipayClient() *alipay.Client {
	s := conf.AlipaySetting
	// 将 key 的验证调整到初始化阶段
	client, err := alipay.New(s.AppID, s.PrivateKey, s.InProduction)
	if err != nil {
		logrus.Fatalf("alipay.New err: %s", err)
	}
	// 加载应用公钥证书
	if err = client.LoadAppPublicCertFromFile(s.AppPublicCertFile); err != nil {
		logrus.Fatalf("client.LoadAppPublicCertFromFile err: %s\n", err)
	}
	// 加载支付宝根证书
	if err = client.LoadAliPayRootCertFromFile(s.RootCertFile); err != nil {
		logrus.Fatalf("client.LoadAliPayRootCertFromFile err: %s\n", err)
	}
	// 加载支付宝公钥证书
	if err = client.LoadAliPayPublicCertFromFile(s.PublicCertFile); err != nil {
		logrus.Fatalf("client.LoadAliPayPublicCertFromFile err: %s\n", err)
	}
	return client
}

// lazyInitial do some package lazy initialize for performance
func lazyInitial() {
	_onceInitial.Do(func() {
		_enablePhoneVerify = cfg.If("Sms")
		_disallowUserRegister = cfg.If("Web:DisallowUserRegister")
	})
}
