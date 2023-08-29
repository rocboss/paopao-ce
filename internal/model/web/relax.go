// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

type GetUnreadMsgCountReq struct {
	SimpleInfo `json:"-" binding:"-"`
}

type GetUnreadMsgCountResp struct {
	Count int64 `json:"count"`
}
