// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

type ChangeUserStatusReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" form:"id" binding:"required"`
	Status   int   `json:"status" form:"status" binding:"required,oneof=1 2"`
}
