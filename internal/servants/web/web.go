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
	api.RegisterWebAdminServant(e, newWebAdminSrv(), newWebAdminBinding(), newWebAdminRender())
	api.RegisterWebCoreServant(e, newWebCoreSrv(), newWebCoreBinding(), newWebCoreRender())
	api.RegisterWebLooseServant(e, newWebLooseSrv(), newWebLooseBinding(), newWebLooseRender())
	api.RegisterWebPrivServant(e, newWebPrivSrv(), newWebPrivBinding(), newWebPrivRender())
	api.RegisterWebPubServant(e, newWebPubSrv(), newWebPubBinding(), newWebPubRender())

	cfg.In(cfg.Actions{
		"Alipay": func() {
			api.RegisterWebAlipayPubServant(e, newWebAlipayPubSrv(), newWebAlipayPubBinding(), newWebAlipayPubRender())
			api.RegisterWebAlipayPrivServant(e, newWebAlipayPrivSrv(), newWebAlipayPrivBinding(), newWebAlipayPrivRender())
		},
		"Followship": func() {
			api.RegisterWebFollowshipServant(e, newWebFollowshipSrv(), newWebFollowshipBinding(), newWebFollowshipRender())
		},
		"Friendship": func() {
			api.RegisterWebFriendshipServant(e, newWebFriendshipSrv(), newWebFriendshipBinding(), newWebFriendshipRender())
		},
	})
}
