// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package base

type Pager struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

type PageResp struct {
	List  any   `json:"list"`
	Pager Pager `json:"pager"`
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
