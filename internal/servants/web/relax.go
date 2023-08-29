// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

var (
	_ api.Relax      = (*relaxSrv)(nil)
	_ api.RelaxChain = (*relaxChain)(nil)
)

type relaxSrv struct {
	api.UnimplementedRelaxServant
	*base.DaoServant
}

type relaxChain struct {
	api.UnimplementedRelaxChain
}

func (s *relaxSrv) GetUnreadMsgCount(req *web.GetUnreadMsgCountReq) (*web.GetUnreadMsgCountResp, mir.Error) {
	count, err := s.Ds.GetUnreadCount(req.Uid)
	if err != nil {
		return nil, xerror.ServerError
	}
	return &web.GetUnreadMsgCountResp{
		Count: count,
	}, nil
}

func (*relaxChain) ChainGetUnreadMsgCount() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtSurely()}
}

func newRelaxSrv(s *base.DaoServant) api.Relax {
	return &relaxSrv{
		DaoServant: s,
	}
}

func newRelaxChain() api.RelaxChain {
	return &relaxChain{}
}
