// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/redis/rueidis"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Relax = (*relaxSrv)(nil)
)

type relaxSrv struct {
	api.UnimplementedRelaxServant
	*base.DaoServant
	wc core.WebCache
}

func (s *relaxSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtSurely()}
}

func (s *relaxSrv) GetUnreadMsgCount(req *web.GetUnreadMsgCountReq) (*web.GetUnreadMsgCountResp, mir.Error) {
	if data, xerr := s.wc.GetUnreadMsgCountResp(req.Uid); xerr == nil && len(data) > 0 {
		// logrus.Debugln("GetUnreadMsgCount get resp from cache")
		return &web.GetUnreadMsgCountResp{
			JsonResp: data,
		}, nil
	} else if !rueidis.IsRedisNil(xerr) {
		logrus.Warnf("GetUnreadMsgCount from cache occurs error: %s", xerr)
	}
	// 使用缓存机制特殊处理
	onCacheUnreadMsgEvent(req.Uid)
	return &web.GetUnreadMsgCountResp{}, nil
}

func newRelaxSrv(s *base.DaoServant, wc core.WebCache) api.Relax {
	return &relaxSrv{
		DaoServant: s,
		wc:         wc,
	}
}
