// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"strings"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
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

func (s *topicServant) UpsertTags(userId int64, tags []string) (_ cs.TagInfoList, err error) {
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

func (s *topicServant) ListTags(typ cs.TagType, offset, limit int) (res cs.TagList, err error) {
	conditions := &core.ConditionsT{}
	switch typ {
	case cs.TagTypeHot:
		// 热门标签
		conditions = &core.ConditionsT{
			"ORDER": "quote_num DESC",
		}
	case cs.TagTypeNew:
		// 最新标签
		conditions = &core.ConditionsT{
			"ORDER": "id DESC",
		}
	}
	// TODO: 优化查询方式，直接返回[]*core.Tag, 目前保持先转换一下
	var (
		tags []*dbr.Tag
		item *cs.TagItem
	)
	if tags, err = (&dbr.Tag{}).List(s.db, conditions, offset, limit); err == nil {
		if len(tags) == 0 {
			return
		}
		tagMap := make(map[int64][]*cs.TagItem, len(tags))
		for _, tag := range tags {
			item = &cs.TagItem{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
			}
			tagMap[item.UserID] = append(tagMap[item.UserID], item)
			res = append(res, item)
		}
		ids := make([]int64, len(tagMap))
		for userId := range tagMap {
			ids = append(ids, userId)
		}
		userInfos, err := (&dbr.User{}).ListUserInfoById(s.db, ids)
		if err != nil {
			return nil, err
		}
		for _, userInfo := range userInfos {
			for _, item = range tagMap[userInfo.ID] {
				item.User = userInfo
			}
		}
	}
	return
}

func (s *topicServant) TagsByKeyword(keyword string) (res cs.TagInfoList, err error) {
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
			res = append(res, &cs.TagInfo{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
			})
		}
	}
	return
}
