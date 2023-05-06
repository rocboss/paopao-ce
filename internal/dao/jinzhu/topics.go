// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"errors"
	"strings"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.TopicService = (*topicServant)(nil)
)

type topicServant struct {
	db             *gorm.DB
	ums            core.UserManageService
	tnTopicUser    string
	tnDotTopicUser string
}

type topicInfo struct {
	TopicId int64
	IsTop   int8
}

func newTopicService(db *gorm.DB, ums core.UserManageService) core.TopicService {
	return &topicServant{
		db:             db,
		ums:            ums,
		tnTopicUser:    db.NamingStrategy.TableName("TopicUser"),
		tnDotTopicUser: db.NamingStrategy.TableName("TopicUser") + ".",
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

func (s *topicServant) GetHotTags(userId int64, limit int, offset int) ([]*core.TagFormated, error) {
	tags, err := (&dbr.Tag{}).List(s.db, &core.ConditionsT{
		"ORDER": "quote_num DESC",
	}, offset, limit)
	if err != nil {
		return nil, err
	}
	return s.tagsFormat(userId, nil, tags)
}

func (s *topicServant) GetNewestTags(userId int64, limit int, offset int) ([]*core.TagFormated, error) {
	tags, err := (&dbr.Tag{}).List(s.db, &core.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
	if err != nil {
		return nil, err
	}
	return s.tagsFormat(userId, nil, tags)
}

func (s *topicServant) GetFollowTags(userId int64, limit int, offset int) ([]*core.TagFormated, error) {
	if userId < 0 {
		return nil, nil
	}
	userTopics := []*topicInfo{}
	err := s.db.Model(&dbr.TopicUser{}).
		Where("user_id=?", userId).
		Order("is_top DESC").
		Limit(limit).
		Offset(offset).
		Find(&userTopics).Error
	if err != nil {
		return nil, err
	}
	userTopicsMap := make(map[int64]*topicInfo, len(userTopics))
	topicIds := make([]int64, 0, len(userTopics))
	topicIdsMap := make(map[int64]int, len(userTopics))
	for idx, info := range userTopics {
		userTopicsMap[info.TopicId] = info
		topicIds = append(topicIds, info.TopicId)
		topicIdsMap[info.TopicId] = idx
	}
	var tags []*core.Tag
	err = s.db.Model(&dbr.Tag{}).Where("quote_num > 0 and id in ?", topicIds).Order("quote_num DESC").Find(&tags).Error
	if err != nil {
		return nil, err
	}
	formtedTags, err := s.tagsFormat(-1, userTopicsMap, tags)
	if err != nil {
		return nil, err
	}
	// 置顶排序后处理
	// TODO: 垃圾办法，最好是topic_user join tag 一次查询，但是gorm的join真他喵的别扭，F*K
	res := make([]*core.TagFormated, len(topicIds), len(topicIds))
	for _, tag := range formtedTags {
		res[topicIdsMap[tag.ID]] = tag
	}
	return res, nil
}

func (s *topicServant) tagsFormat(userId int64, userTopicsMap map[int64]*topicInfo, tags []*core.Tag) ([]*core.TagFormated, error) {
	// 获取创建者User IDs
	userIds := []int64{}
	tagIds := []int64{}
	for _, tag := range tags {
		userIds = append(userIds, tag.UserID)
		tagIds = append(tagIds, tag.ID)
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
	// 填充话题follow信息
	if userId > -1 && len(userTopicsMap) <= 0 {
		userTopics := []*topicInfo{}
		err = s.db.Model(&dbr.TopicUser{}).Where("is_del=0 and user_id=? and topic_id in ?", userId, tagIds).Find(&userTopics).Error
		if err != nil {
			return nil, err
		}
		userTopicsMap = make(map[int64]*topicInfo, len(userTopics))
		for _, info := range userTopics {
			userTopicsMap[info.TopicId] = info
		}
	}
	if len(userTopicsMap) > 0 {
		for _, tag := range tagsFormated {
			if info, exist := userTopicsMap[tag.ID]; exist {
				tag.IsFollowing, tag.IsTop = 1, info.IsTop
			}
		}
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

func (s *topicServant) FollowTopic(userId int64, topicId int64) (err error) {
	return s.db.Create(&dbr.TopicUser{
		UserID:  userId,
		TopicID: topicId,
		IsTop:   0,
	}).Error
}

func (s *topicServant) UnfollowTopic(userId int64, topicId int64) error {
	return s.db.Exec("DELETE FROM "+s.tnTopicUser+" WHERE user_id=? AND topic_id=?", userId, topicId).Error
}

func (s *topicServant) StickTopic(userId int64, topicId int64) (status int8, err error) {
	db := s.db.Begin()
	defer db.Rollback()

	m := &dbr.TopicUser{}
	err = db.Model(m).
		Where("user_id=? and topic_id=?", userId, topicId).
		UpdateColumn("is_top", gorm.Expr("1-is_top")).Error
	if err != nil {
		return
	}
	status = -1
	err = db.Model(m).Where("user_id=? and topic_id=?", userId, topicId).Select("is_top").Scan(&status).Error
	if err != nil {
		return
	}
	if status < 0 {
		return -1, errors.New("topic not exist")
	}

	db.Commit()
	return
}
