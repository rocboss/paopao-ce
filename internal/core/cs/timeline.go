// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

import (
	"github.com/rocboss/paopao-ce/internal/core/ms"
)

// TweetBox 推文列表盒子，包含其他一些关于推文列表的信息
type TweetBox struct {
	Tweets TweetList
	Total  int64
}

// ContentBox 综合内容列表盒子，包含动态和长文章
type ContentBox struct {
	Contents ContentList `json:"contents"`
	Total    int64       `json:"total"`
}

// ContentItem 内容项
type ContentItem struct {
	Type      ms.ContentType `json:"type"`
	Tweet     *TweetItem     `json:"tweet,omitempty"`
	Article   interface{}    `json:"article,omitempty"`
	CreatedOn int64          `json:"created_on"`
}

// ContentList 内容列表
type ContentList []*ContentItem
