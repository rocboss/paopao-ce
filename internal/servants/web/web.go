// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"sync"

	"github.com/alimy/tryst/cfg"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

var (
	_enablePhoneVerify    bool
	_disallowUserRegister bool
	_ds                   core.DataService
	_ac                   core.AppCache
	_wc                   core.WebCache
	_oss                  core.ObjectStorageService
	_onceInitial          sync.Once
)

// RouteWeb register web route
func RouteWeb(e *gin.Engine) {
	lazyInitial()
	ds := base.NewDaoServant()
	// aways register servants
	api.RegisterAdminServant(e, newAdminSrv(ds, _wc))
	api.RegisterCoreServant(e, newCoreSrv(ds, _oss, _wc))
	api.RegisterRelaxServant(e, newRelaxSrv(ds, _wc), newRelaxChain())
	api.RegisterLooseServant(e, newLooseSrv(ds, _ac))
	api.RegisterPrivServant(e, newPrivSrv(ds, _oss), newPrivChain())
	api.RegisterPubServant(e, newPubSrv(ds))
	api.RegisterFollowshipServant(e, newFollowshipSrv(ds))
	api.RegisterFriendshipServant(e, newFriendshipSrv(ds))
	// regster servants if needed by configure
	cfg.Be("Alipay", func() {
		client := conf.MustAlipayClient()
		api.RegisterAlipayPubServant(e, newAlipayPubSrv(ds))
		api.RegisterAlipayPrivServant(e, newAlipayPrivSrv(ds, client))
	})
}

// lazyInitial do some package lazy initialize for performance
func lazyInitial() {
	_onceInitial.Do(func() {
		_enablePhoneVerify = cfg.If("Sms")
		_disallowUserRegister = cfg.If("Web:DisallowUserRegister")
		_oss = dao.ObjectStorageService()
		_ds = dao.DataService()
		_ac = cache.NewAppCache()
		_wc = cache.NewWebCache()
	})
}
