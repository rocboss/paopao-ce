// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

// TrendsManageServantA 动态信息管理服务
type TrendsManageServantA interface {
	GetIndexTrends(userId int64, limit int, offset int) ([]*cs.TrendsItem, int64, error)
}
