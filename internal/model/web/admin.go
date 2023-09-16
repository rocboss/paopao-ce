// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

type ChangeUserStatusReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" form:"id" binding:"required"`
	Status   int   `json:"status" form:"status" binding:"required,oneof=1 2"`
}

type SiteInfoReq struct {
	SimpleInfo `json:"-" binding:"-"`
}

type SiteInfoResp struct {
	RegisterUserCount int64 `json:"register_user_count"`
	OnlineUserCount   int   `json:"online_user_count"`
	HistoryMaxOnline  int   `json:"history_max_online"`
	ServerUpTime      int64 `json:"server_up_time"`
}
