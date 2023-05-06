// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data any) {
	hostname, _ := os.Hostname()
	if data == nil {
		data = gin.H{
			"code":      0,
			"msg":       "success",
			"tracehost": hostname,
		}
	} else {
		data = gin.H{
			"code":      0,
			"msg":       "success",
			"data":      data,
			"tracehost": hostname,
		}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list any, totalRows int64) {
	r.ToResponse(gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *xerror.Error) {
	response := gin.H{"code": err.StatusCode(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(xerror.HttpStatusCode(err), response)
}
