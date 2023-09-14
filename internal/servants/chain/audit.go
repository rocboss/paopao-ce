// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package chain

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func AuditHook() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 此midleware后面是真正的http handlder，让handler先执行
		c.Next()
		// 审核hook 后处理逻辑
		var ami *web.AuditMetaInfo
		if val, ok := c.Get(web.AuditHookCtxKey); ok {
			if ami, ok = val.(*web.AuditMetaInfo); !ok {
				return
			}
		}
		OnAudiotHookEvent(ami)
	}
}
