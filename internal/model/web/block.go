// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

type BlockUserReq struct {
	BaseInfo     `json:"-" binding:"-"`
	UserId       int64  `json:"user_id" binding:"required"`
	Reason       string `json:"reason"`
}

type UnblockUserReq struct {
	BaseInfo `json:"-" binding:"-"`
	UserId   int64 `json:"user_id" binding:"required"`
}

type ListBlocksReq struct {
	BaseInfo `json:"-" binding:"-"`
	joint.BasePageInfo
}

type ListBlocksResp base.PageResp
