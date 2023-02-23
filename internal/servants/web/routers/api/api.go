// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/alimy/cfg"
	"github.com/redis/go-redis/v9"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
)

var (
	redisClient   *redis.Client
	alipayClient  *alipay.Client
	objectStorage core.ObjectStorageService
)

func Initialize() {
	redisClient = conf.MustRedis()
	objectStorage = dao.ObjectStorageService()

	if cfg.If("Alipay") {
		alipayClient = mustAlipayClient()
	}
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
