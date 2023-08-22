// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package joint provider some common base type or logic for model define.

package joint

type BasePageInfo struct {
	Page     int `form:"-" binding:"-"`
	PageSize int `form:"-" binding:"-"`
}

func (r *BasePageInfo) SetPageInfo(page int, pageSize int) {
	r.Page, r.PageSize = page, pageSize
}
