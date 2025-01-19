// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package joint

import "github.com/rocboss/paopao-ce/internal/conf"

type Pager struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

type PageResp struct {
	List  any   `json:"list"`
	Pager Pager `json:"pager"`
}

type PageInfo struct {
	Limit  int `json:"-" form:"-" query:"limit" binding:"-"`
	Offset int `json:"-" form:"-" query:"offset" binding:"-"`
}

func PageRespFrom(list any, page int, pageSize int, totalRows int64) *PageResp {
	return &PageResp{
		List: list,
		Pager: Pager{
			Page:      page,
			PageSize:  pageSize,
			TotalRows: totalRows,
		},
	}
}

func (s *PageInfo) BuildPageInfo(page, size int) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		s.Limit = conf.AppSetting.DefaultPageSize
	} else if size > conf.AppSetting.MaxPageSize {
		s.Limit = conf.AppSetting.MaxPageSize
	}
	s.Offset = (page - 1) * s.Limit
}
