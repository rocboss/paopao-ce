// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return conf.AppSetting.DefaultPageSize
	}
	if pageSize > conf.AppSetting.MaxPageSize {
		return conf.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(c *gin.Context) (offset, limit int) {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}

	limit = convert.StrTo(c.Query("page_size")).MustInt()
	if limit <= 0 {
		limit = conf.AppSetting.DefaultPageSize
	} else if limit > conf.AppSetting.MaxPageSize {
		limit = conf.AppSetting.MaxPageSize
	}
	offset = (page - 1) * limit
	return
}

func GetPageInfo(c *gin.Context) (page, pageSize int) {
	page = convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}
	pageSize = convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		pageSize = conf.AppSetting.DefaultPageSize
	} else if pageSize > conf.AppSetting.MaxPageSize {
		pageSize = conf.AppSetting.MaxPageSize
	}
	return
}
