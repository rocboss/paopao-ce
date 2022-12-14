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

func (s *topicServant) CreateTag(tag *core.Tag) (*core.Tag, error) {
	return createTag(s.db, tag)
}

func (s *topicServant) DeleteTag(tag *core.Tag) error {
	return deleteTag(s.db, tag)
}

func (s *topicServant) GetTags(conditions *core.ConditionsT, offset, limit int) ([]*core.Tag, error) {
	return (&dbr.Tag{}).List(s.db, conditions, offset, limit)
}

func (s *topicServant) GetTagsByKeyword(keyword string) ([]*core.Tag, error) {
	tag := &dbr.Tag{}

	keyword = "%" + strings.Trim(keyword, " ") + "%"
	if keyword == "%%" {
		return tag.List(s.db, &dbr.ConditionsT{
			"ORDER": "quote_num DESC",
		}, 0, 6)
	} else {
		return tag.List(s.db, &dbr.ConditionsT{
			"tag LIKE ?": keyword,
			"ORDER":      "quote_num DESC",
		}, 0, 6)
	}
}
