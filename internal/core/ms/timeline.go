// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ms

type IndexTweetList struct {
	Tweets []*PostFormated
	Total  int64
}

// ContentType 内容类型
type ContentType int

const (
	ContentTypeTweet ContentType = iota + 1
	ContentTypeArticle
)

// IndexContentItem 首页内容项
type IndexContentItem struct {
	Type      ContentType    `json:"type"`
	Tweet     *PostFormated  `json:"tweet,omitempty"`
	Article   interface{}    `json:"article,omitempty"`
	CreatedOn int64          `json:"created_on"`
}

// IndexContentList 首页综合内容列表
type IndexContentList struct {
	Contents []*IndexContentItem `json:"contents"`
	Total    int64               `json:"total"`
}
