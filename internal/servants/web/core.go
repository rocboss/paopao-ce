// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
)

var (
	_ api.Core        = (*coreSrv)(nil)
	_ api.CoreBinding = (*coreBinding)(nil)
	_ api.CoreRender  = (*coreRender)(nil)
)

type coreSrv struct {
	api.UnimplementedCoreServant
	*base.DaoServant

	ts  core.TweetSearchService
	oss core.ObjectStorageService
}

type coreBinding struct {
	*api.UnimplementedCoreBinding
}

type coreRender struct {
	*api.UnimplementedCoreRender
}

func (s *coreSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func newCoreSrv(s *base.DaoServant, ts core.TweetSearchService, oss core.ObjectStorageService) api.Core {
	return &coreSrv{
		DaoServant: s,
		ts:         ts,
		oss:        oss,
	}
}

func newCoreBinding() api.CoreBinding {
	return &coreBinding{
		UnimplementedCoreBinding: &api.UnimplementedCoreBinding{
			BindAny: base.BindAny,
		},
	}
}

func newCoreRender() api.CoreRender {
	return &coreRender{
		UnimplementedCoreRender: &api.UnimplementedCoreRender{
			RenderAny: base.RenderAny,
		},
	}
}
