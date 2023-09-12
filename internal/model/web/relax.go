// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/joint"
)

type GetUnreadMsgCountReq struct {
	SimpleInfo `json:"-" binding:"-"`
}

type GetUnreadMsgCountResp struct {
	Count    int64           `json:"count"`
	JsonResp json.RawMessage `json:"-"`
}

func (r *GetUnreadMsgCountResp) Render(c *gin.Context) {
	if len(r.JsonResp) != 0 {
		c.JSON(http.StatusOK, r.JsonResp)
	} else {
		c.JSON(http.StatusOK, &joint.JsonResp{
			Code: 0,
			Msg:  "success",
			Data: r,
		})
	}
}
