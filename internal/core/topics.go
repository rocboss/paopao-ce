// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

const (
	TagCategoryHot TagCategory = "hot"
	TagCategoryNew TagCategory = "new"
)

type TagCategory string

type Tag struct {
	ID       int64  `json:"id" db:"id"`
	UserID   int64  `json:"user_id" db:"user_id"`
	Tag      string `json:"tag"`
	QuoteNum int64  `json:"quote_num" db:"quote_num"`
}

type TagFormated struct {
	ID       int64         `json:"id"`
	UserID   int64         `json:"user_id"`
	User     *UserFormated `json:"user"`
	Tag      string        `json:"tag"`
	QuoteNum int64         `json:"quote_num"`
}

// TopicService 话题服务
type TopicService interface {
	UpsertTags(userId int64, tags []string) ([]*Tag, error)
	DecrTagsById(ids []int64) error
	GetTags(category TagCategory, offset int, limit int) ([]*Tag, error)
	GetTagsByKeyword(keyword string) ([]*Tag, error)
}

func (t *Tag) Format() *TagFormated {
	return &TagFormated{
		ID:       t.ID,
		UserID:   t.UserID,
		User:     &UserFormated{},
		Tag:      t.Tag,
		QuoteNum: t.QuoteNum,
	}
}
