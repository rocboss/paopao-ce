//go:build constraint

package web

import (
	api "github.com/rocboss/paopao-ce/auto/api/v1"
)

var (
	_ api.Admin      = (*adminSrv)(nil)
	_ api.AlipayPub  = (*alipayPubSrv)(nil)
	_ api.AlipayPriv = (*alipayPrivSrv)(nil)
	_ api.Core       = (*coreSrv)(nil)
	_ api.Followship = (*followshipSrv)(nil)
	_ api.Friendship = (*friendshipSrv)(nil)
	_ api.Loose      = (*looseSrv)(nil)
	_ api.Pub        = (*pubSrv)(nil)
	_ api.Relax      = (*relaxSrv)(nil)
	_ api.Site       = (*siteSrv)(nil)
	_ api.Trends     = (*trendsSrv)(nil)
	_ api.Priv       = (*privSrv)(nil)
	_ api.PrivChain  = (*privChain)(nil)
)
