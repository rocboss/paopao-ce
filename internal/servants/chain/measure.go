// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package chain

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

func OnlineUserMeasure() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 此midleware后面是真正的http handlder，让handler先执行
		c.Next()
		// 更新用户在线状态
		if uid, ok := base.UserIdFrom(c); ok {
			OnUserOnlineMetric(_ac, uid)
		}
	}
}
