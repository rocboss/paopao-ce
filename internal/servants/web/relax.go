// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/rueidis"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	_ api.Relax = (*relaxSrv)(nil)
)

type relaxSrv struct {
	api.UnimplementedRelaxServant
	*base.DaoServant
	wc core.WebCache
}

type relaxChain struct {
	api.UnimplementedRelaxChain
}

func (s *relaxChain) ChainGetUnreadMsgCount() gin.HandlersChain {
	return gin.HandlersChain{chain.OnlineUserMeasure()}
}

func (s *relaxSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtSurely()}
}

func (s *relaxSrv) GetUnreadMsgCount(req *web.GetUnreadMsgCountReq) (*web.GetUnreadMsgCountResp, error) {
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

func (s *relaxSrv) StreamUnreadMsgCount(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")

	uid, exists := c.Get("uid")
	if !exists {
		c.AbortWithStatus(401)
		return
	}
	uidInt, ok := uid.(int64)
	if !ok {
		c.AbortWithStatus(400)
		return
	}

	var lastCount int64 = -1
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data, xerr := s.wc.GetUnreadMsgCountResp(uidInt)
			count := int64(0)

			if xerr == nil && len(data) > 0 {
				var msgCount struct {
					Count int64 `json:"count"`
				}
				if err := json.Unmarshal(data, &msgCount); err == nil {
					count = msgCount.Count
				}
			}

			if count != lastCount {
				c.SSEvent("message", map[string]int64{"count": count})
				lastCount = count
			} else {
				c.SSEvent("", "")
			}
			c.Writer.Flush()
		case <-c.Writer.CloseNotify():
			return
		}
	}
}

func newRelaxSrv(s *base.DaoServant, wc core.WebCache) api.Relax {
	return &relaxSrv{
		DaoServant: s,
		wc:         wc,
	}
}

func newRelaxChain() api.RelaxChain {
	return &relaxChain{}
}
