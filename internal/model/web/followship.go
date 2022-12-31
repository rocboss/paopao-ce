// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import "github.com/rocboss/paopao-ce/internal/servants/base"

type AddFollowingReq struct {
	BaseInfo `json:"-" binding:"-"`
}

type DeleteFollowingReq struct {
	BaseInfo `json:"-" binding:"-"`
}

type ListFollowingsReq struct {
	BaseInfo `json:"-" binding:"-"`
}

type ListFollowingsResp base.PageResp

type ListFollowersReq struct {
	BaseInfo `form:"-" binding:"-"`
}

type ListFollowersResp base.PageResp
