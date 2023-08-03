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
	api.RegisterAdminServant(e, newAdminSrv(ds))
	api.RegisterCoreServant(e, newCoreSrv(ds, oss))
	api.RegisterLooseServant(e, newLooseSrv(ds))
	api.RegisterPrivServant(e, newPrivSrv(ds, oss))
	api.RegisterPubServant(e, newPubSrv(ds))
	api.RegisterKeyQueryServant(e, NewShareKeyServant(ds))
	// regster servants if needed by configure
	cfg.In(cfg.Actions{
		"Alipay": func() {
			client := conf.MustAlipayClient()
			api.RegisterAlipayPubServant(e, newAlipayPubSrv(ds))
			api.RegisterAlipayPrivServant(e, newAlipayPrivSrv(ds, client))
		},
		"Followship": func() {
			api.RegisterFollowshipServant(e, newFollowshipSrv(ds))
		},
		"Friendship": func() {
			api.RegisterFriendshipServant(e, newFriendshipSrv(ds))
		},
	})
}

// lazyInitial do some package lazy initialize for performance
func lazyInitial() {
	_onceInitial.Do(func() {
		_enablePhoneVerify = cfg.If("Sms")
		_disallowUserRegister = cfg.If("Web:DisallowUserRegister")
	})
}
