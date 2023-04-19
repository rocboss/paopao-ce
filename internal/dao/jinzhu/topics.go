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
	db  *gorm.DB
	ums core.UserManageService
}

func newTopicService(db *gorm.DB, ums core.UserManageService) core.TopicService {
	return &topicServant{
		db:  db,
		ums: ums,
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

func (s *topicServant) GetHotTags(limit int, offset int) ([]*core.TagFormated, error) {
	tags, err := (&dbr.Tag{}).List(s.db, &core.ConditionsT{
		"ORDER": "quote_num DESC",
	}, offset, limit)
	if err != nil {
		return nil, err
	}
	return s.tagsFormat(tags)
}

func (s *topicServant) GetNewestTags(limit int, offset int) ([]*core.TagFormated, error) {
	tags, err := (&dbr.Tag{}).List(s.db, &core.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
	if err != nil {
		return nil, err
	}
	return s.tagsFormat(tags)
}

func (s *topicServant) GetFollowTags(limit int, offset int) ([]*core.TagFormated, error) {
	// TODO：具体逻辑稍后实现，先用热门标签替换
	tags, err := (&dbr.Tag{}).List(s.db, &core.ConditionsT{
		"ORDER": "quote_num DESC",
	}, offset, limit)
	if err != nil {
		return nil, err
	}
	return s.tagsFormat(tags)
}

func (s *topicServant) tagsFormat(tags []*core.Tag) ([]*core.TagFormated, error) {
	// 获取创建者User IDs
	userIds := []int64{}
	for _, tag := range tags {
		userIds = append(userIds, tag.UserID)
	}
	users, err := s.ums.GetUsersByIDs(userIds)
	if err != nil {
		return nil, err
	}
	tagsFormated := []*core.TagFormated{}
	for _, tag := range tags {
		tagFormated := tag.Format()
		for _, user := range users {
			if user.ID == tagFormated.UserID {
				tagFormated.User = user.Format()
			}
		}
		tagsFormated = append(tagsFormated, tagFormated)
	}
	return tagsFormated, nil
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
