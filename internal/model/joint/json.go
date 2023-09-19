// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package joint

import (
	stdJson "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/pkg/json"
)

type CachePageResp struct {
	Data     *PageResp
	JsonResp stdJson.RawMessage
}

func (r *CachePageResp) Render(c *gin.Context) {
	if len(r.JsonResp) != 0 {
		c.JSON(http.StatusOK, r.JsonResp)
	} else {
		c.JSON(http.StatusOK, &JsonResp{
			Code: 0,
			Msg:  "success",
			Data: r.Data,
		})
	}
}

func RespMarshal(data any) (stdJson.RawMessage, error) {
	return json.Marshal(data)
}
