// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

import "github.com/rocboss/paopao-ce/pkg/types"

type TrendsItem struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	IsFresh  bool   `json:"is_fresh" gorm:"-"`
}

func DistinctTrends(items []*TrendsItem) []*TrendsItem {
	if len(items) == 0 {
		return items
	}
	res := make([]*TrendsItem, 0, len(items))
	set := make(map[string]types.Empty, len(items))
	for _, item := range items {
		if _, exist := set[item.Username]; exist {
			continue
		}
		res = append(res, item)
		set[item.Username] = types.Empty{}
	}
	return res
}
