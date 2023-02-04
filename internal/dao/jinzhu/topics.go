// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"strings"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.TopicService = (*topicServant)(nil)
)

type topicServant struct {
	db *gorm.DB
}

func newTopicService(db *gorm.DB) core.TopicService {
	return &topicServant{
		db: db,
	}
}

func (s *topicServant) UpsertTags(userId int64, tags []string) (_ []*core.Tag, err error) {
	db := s.db.Begin()
	defer func() {
		if err == nil {
			db.Commit()
		} else {
			db.Rollback()
		}
	}()
	return createTags(db, userId, tags)
}

func (s *topicServant) DecrTagsById(ids []int64) (err error) {
	db := s.db.Begin()
	defer func() {
		if err == nil {
			db.Commit()
		} else {
			db.Rollback()
		}
	}()
	return decrTagsByIds(db, ids)
}

func (s *topicServant) GetTags(category core.TagCategory, offset, limit int) (res []*core.Tag, err error) {
	conditions := &core.ConditionsT{}
	switch category {
	case core.TagCategoryHot:
		// 热门标签
		conditions = &core.ConditionsT{
			"ORDER": "quote_num DESC",
		}
	case core.TagCategoryNew:
		// 最新标签
		conditions = &core.ConditionsT{
			"ORDER": "id DESC",
		}
	}
	// TODO: 优化查询方式，直接返回[]*core.Tag, 目前保持先转换一下
	var tags []*dbr.Tag
	if tags, err = (&dbr.Tag{}).List(s.db, conditions, offset, limit); err == nil {
		for _, tag := range tags {
			res = append(res, &core.Tag{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
			})
		}
	}
	return
}

func (s *topicServant) GetTagsByKeyword(keyword string) (res []*core.Tag, err error) {
	keyword = "%" + strings.Trim(keyword, " ") + "%"
	tag := &dbr.Tag{}
	var tags []*dbr.Tag
	if keyword == "%%" {
		tags, err = tag.List(s.db, &dbr.ConditionsT{
			"ORDER": "quote_num DESC",
		}, 0, 6)
	} else {
		tags, err = tag.List(s.db, &dbr.ConditionsT{
			"tag LIKE ?": keyword,
			"ORDER":      "quote_num DESC",
		}, 0, 6)
	}
	if err == nil {
		for _, tag := range tags {
			res = append(res, &core.Tag{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
			})
		}
	}
	return
}
