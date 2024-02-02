// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
)

var (
	_alipayClient *alipay.Client
	_onceAlipay   sync.Once
)

func MustAlipayClient() *alipay.Client {
	_onceAlipay.Do(func() {
		s := AlipaySetting
		// 将 key 的验证调整到初始化阶段
		client, err := alipay.New(s.AppID, s.PrivateKey, s.InProduction)
		if err != nil {
			logrus.Fatalf("alipay.New err: %s", err)
		}
		// 加载应用公钥证书
		if err = client.LoadAppCertPublicKeyFromFile(s.AppPublicCertFile); err != nil {
			logrus.Fatalf("client.LoadAppCertPublicKeyFromFile err: %s\n", err)
		}
		// 加载支付宝根证书
		if err = client.LoadAliPayRootCertFromFile(s.RootCertFile); err != nil {
			logrus.Fatalf("client.LoadAliPayRootCertFromFile err: %s\n", err)
		}
		// 加载支付宝公钥证书
		if err = client.LoadAlipayCertPublicKeyFromFile(s.PublicCertFile); err != nil {
			logrus.Fatalf("client.LoadAlipayCertPublicKeyFromFile err: %s\n", err)
		}
		_alipayClient = client
	})
	return _alipayClient
}
