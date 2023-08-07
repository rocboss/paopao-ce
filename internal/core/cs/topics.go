// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

const (
	// 标签类型
	TagTypeHot       TagType = "hot"
	TagTypeNew       TagType = "new"
	TagTypeFollow    TagType = "follow"
	TagTypeHotExtral TagType = "hot_extral"
)

type (
	// TagType 标签类型
	TagType string

	// TagInfoList 标签信息列表
	TagInfoList []*TagInfo

	// TagList 标签列表
	TagList []*TagItem
)

// TagInfo 标签信息
type TagInfo struct {
	ID       int64  `json:"id" db:"id"`
	UserID   int64  `json:"user_id" db:"user_id"`
	Tag      string `json:"tag"`
	QuoteNum int64  `json:"quote_num" db:"quote_num"`
}

// TagItem 标签信息条陈
type TagItem struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	User        *UserInfo `json:"user" db:"u"`
	Tag         string    `json:"tag"`
	QuoteNum    int64     `json:"quote_num"`
	IsFollowing int8      `json:"is_following"`
	IsTop       int8      `json:"is_top"`
}

func (t *TagInfo) Format() *TagItem {
	return &TagItem{
		ID:          t.ID,
		UserID:      t.UserID,
		User:        &UserInfo{},
		Tag:         t.Tag,
		QuoteNum:    t.QuoteNum,
		IsFollowing: 0,
		IsTop:       0,
	}
}
