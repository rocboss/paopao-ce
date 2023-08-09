// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

type FollowUserReq struct {
	BaseInfo `json:"-" binding:"-"`
	UserId   int64 `json:"user_id" binding:"required"`
}

type UnfollowUserReq struct {
	BaseInfo `json:"-" binding:"-"`
	UserId   int64 `json:"user_id" binding:"required"`
}

type ListFollowsReq struct {
	BaseInfo `json:"-" binding:"-"`
	joint.BasePageInfo
}

type ListFollowsResp base.PageResp

type ListFollowingsReq struct {
	BaseInfo `form:"-" binding:"-"`
	joint.BasePageInfo
}

type ListFollowingsResp base.PageResp
