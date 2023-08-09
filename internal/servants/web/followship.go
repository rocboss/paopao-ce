// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
)

var (
	_ api.Followship = (*followshipSrv)(nil)
)

type followshipSrv struct {
	api.UnimplementedFollowshipServant
	*base.DaoServant
}

func (s *followshipSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func newFollowshipSrv(s *base.DaoServant) api.Followship {
	return &followshipSrv{}
}
