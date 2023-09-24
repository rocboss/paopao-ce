// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/internal/model/joint"
)

type GetIndexTrendsReq struct {
	SimpleInfo `json:"-" binding:"-"`
	joint.BasePageInfo
}

type GetIndexTrendsResp struct {
	joint.CachePageResp
}
