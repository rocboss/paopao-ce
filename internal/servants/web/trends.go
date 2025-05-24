// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Trends = (*trendsSrv)(nil)
)

type trendsSrv struct {
	api.UnimplementedTrendsServant
	*base.DaoServant
	ac                core.AppCache
	indexTrendsExpire int64
	prefixTrends      string
}

func (s *trendsSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (s *trendsSrv) GetIndexTrends(req *web.GetIndexTrendsReq) (res *web.GetIndexTrendsResp, _ error) {
	limit, offset := req.PageSize, (req.Page-1)*req.PageSize
	// 尝试直接从缓存中获取数据
	key, ok := "", false
	if res, key, ok = s.trendsFromCache(req, limit, offset); ok {
		// logrus.Debugf("trendsSrv.GetIndexTrends from cache key:%s", key)
		return
	}
	trends, totalRows, err := s.Ds.GetIndexTrends(req.Uid, limit, offset)
	if err != nil {
		logrus.Errorf("Ds.GetIndexTrends err[1]: %s", err)
		return nil, web.ErrGetIndexTrendsFailed
	}
	resp := joint.PageRespFrom(trends, req.Page, req.PageSize, totalRows)
	// 缓存处理
	base.OnCacheRespEvent(s.ac, key, resp, s.indexTrendsExpire)
	return &web.GetIndexTrendsResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *trendsSrv) trendsFromCache(req *web.GetIndexTrendsReq, limit int, offset int) (res *web.GetIndexTrendsResp, key string, ok bool) {
	key = fmt.Sprintf("%s%d:%d:%d", s.prefixTrends, req.Uid, limit, offset)
	if data, err := s.ac.Get(key); err == nil {
		ok, res = true, &web.GetIndexTrendsResp{
			CachePageResp: joint.CachePageResp{
				JsonResp: data,
			},
		}
	}
	return
}

func newTrendsSrv(s *base.DaoServant) api.Trends {
	cs := conf.CacheSetting
	return &trendsSrv{
		DaoServant:        s,
		ac:                _ac,
		indexTrendsExpire: cs.IndexTrendsExpire,
		prefixTrends:      conf.PrefixIdxTrends,
	}
}
