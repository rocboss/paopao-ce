// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

const (
	// 搜索查询类型
	TsQueryTypeDefault TsQueryType = "search"
	TsQueryTypeTag     TsQueryType = "tag"
)

type (
	// TsQueryType 搜索查询类型
	TsQueryType string

	// TsDocList 索引条陈列表
	TsDocList []TsDocItem
)

// TsQueryReq 搜索查询请求
type TsQueryReq struct {
	Query      string
	Visibility []TweetVisibleType
	Type       TsQueryType
}

// TsQueryResp 搜索查询响应
type TsQueryResp struct {
	Items TweetList
	Total int64
}

// TsDocItem 索引条陈
type TsDocItem struct {
	Post    *TweetInfo
	Content string
}
