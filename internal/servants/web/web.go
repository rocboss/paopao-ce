// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/cfg"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
)

// RouteWeb register web route
func RouteWeb(e *gin.Engine) {
	api.RegisterAdminServant(e, newAdminSrv(), newAdminBinding(), newAdminRender())
	api.RegisterCoreServant(e, newCoreSrv(), newCoreBinding(), newCoreRender())
	api.RegisterLooseServant(e, newLooseSrv(), newLooseBinding(), newLooseRender())
	api.RegisterPrivServant(e, newPrivSrv(), newPrivBinding(), newPrivRender())
	api.RegisterPubServant(e, newPubSrv(), newPubBinding(), newPubRender())

	cfg.In(cfg.Actions{
		"Alipay": func() {
			api.RegisterAlipayPubServant(e, newAlipayPubSrv(), newAlipayPubBinding(), newAlipayPubRender())
			api.RegisterAlipayPrivServant(e, newAlipayPrivSrv(), newAlipayPrivBinding(), newAlipayPrivRender())
		},
		"Followship": func() {
			api.RegisterFollowshipServant(e, newFollowshipSrv(), newFollowshipBinding(), newFollowshipRender())
		},
		"Friendship": func() {
			api.RegisterFriendshipServant(e, newFriendshipSrv(), newFriendshipBinding(), newFriendshipRender())
		},
	})
}
