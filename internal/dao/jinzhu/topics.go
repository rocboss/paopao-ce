// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"errors"
	"strings"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.TopicService  = (*topicSrv)(nil)
	_ core.TopicServantA = (*topicSrvA)(nil)
)

type topicSrv struct {
	db             *gorm.DB
	tnTopicUser    string
	tnDotTopicUser string
}

type topicSrvA struct {
	db *gorm.DB
}

type topicInfo struct {
	TopicId int64
	IsTop   int8
	IsPin   int8
}

func newTopicService(db *gorm.DB) core.TopicService {
	return &topicSrv{
		db:             db,
		tnTopicUser:    db.NamingStrategy.TableName("TopicUser"),
		tnDotTopicUser: db.NamingStrategy.TableName("TopicUser") + ".",
	}
}

func newTopicServantA(db *gorm.DB) core.TopicServantA {
	return &topicSrvA{
		db: db,
	}
}

func (s *topicSrv) UpsertTags(userId int64, tags []string) (_ cs.TagInfoList, err error) {
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

func (s *topicSrv) DecrTagsById(ids []int64) (err error) {
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

func (s *topicSrv) ListTags(typ cs.TagType, offset, limit int) (res cs.TagList, err error) {
	conditions := &ms.ConditionsT{}
	switch typ {
	case cs.TagTypeHot:
		// 热门标签
		conditions = &ms.ConditionsT{
			"ORDER": "quote_num DESC",
		}
	case cs.TagTypeNew:
		// 最新标签
		conditions = &ms.ConditionsT{
			"ORDER": "id DESC",
		}
	}
	return s.listTags(conditions, limit, offset)
}

func (s *topicSrv) GetHotTags(userId int64, limit int, offset int) (cs.TagList, error) {
	tags, err := s.listTags(&ms.ConditionsT{
		"ORDER": "quote_num DESC",
	}, limit, offset)
	if err != nil {
		return nil, err
	}
	return s.tagsFormatA(userId, tags)
}

func (s *topicSrv) GetNewestTags(userId int64, limit int, offset int) (cs.TagList, error) {
	tags, err := s.listTags(&ms.ConditionsT{
		"ORDER": "id DESC",
	}, limit, offset)
	if err != nil {
		return nil, err
	}
	return s.tagsFormatA(userId, tags)
}

func (s *topicSrv) GetFollowTags(userId int64, is_pin bool, limit int, offset int) (cs.TagList, error) {
	if userId < 0 {
		return nil, nil
	}
	userTopics := []*topicInfo{}
	db := s.db.Model(&dbr.TopicUser{}).Order("is_top DESC").Limit(limit).Offset(offset)
	if is_pin {
		db = db.Where("user_id=? AND is_pin=1", userId)
	} else {
		db = db.Where("user_id=?", userId)
	}
	err := db.Find(&userTopics).Error
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
	var tags cs.TagInfoList
	err = s.db.Model(&dbr.Tag{}).Where("quote_num > 0 and id in ?", topicIds).Order("quote_num DESC").Find(&tags).Error
	if err != nil {
		return nil, err
	}
	formtedTags, err := s.tagsFormatB(userTopicsMap, tags)
	if err != nil {
		return nil, err
	}
	// 置顶排序后处理
	// TODO: 垃圾办法，最好是topic_user join tag 一次查询，但是gorm的join真他喵的别扭，F*K
	res := make(cs.TagList, len(topicIds))
	for _, tag := range formtedTags {
		res[topicIdsMap[tag.ID]] = tag
	}
	return res, nil
}

func (s *topicSrv) listTags(conditions *ms.ConditionsT, limit int, offset int) (res cs.TagList, err error) {
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

func (s *topicSrv) tagsFormatA(userId int64, tags cs.TagList) (cs.TagList, error) {
	// 获取创建者User IDs
	tagIds := []int64{}
	for _, tag := range tags {
		tagIds = append(tagIds, tag.ID)
	}
	// 填充话题follow信息
	if userId > -1 {
		userTopics := []*topicInfo{}
		err := s.db.Model(&dbr.TopicUser{}).Where("is_del=0 and user_id=? and topic_id in ?", userId, tagIds).Find(&userTopics).Error
		if err != nil {
			return nil, err
		}
		userTopicsMap := make(map[int64]*topicInfo, len(userTopics))
		for _, info := range userTopics {
			userTopicsMap[info.TopicId] = info
		}
		for _, tag := range tags {
			if info, exist := userTopicsMap[tag.ID]; exist {
				tag.IsFollowing, tag.IsTop, tag.IsPin = 1, info.IsTop, info.IsPin
			}
		}
	}
	return tags, nil
}

func (s *topicSrv) tagsFormatB(userTopicsMap map[int64]*topicInfo, tags cs.TagInfoList) (cs.TagList, error) {
	// 获取创建者User IDs
	userIds := []int64{}
	for _, tag := range tags {
		userIds = append(userIds, tag.UserID)
	}
	users, err := (&dbr.User{}).ListUserInfoById(s.db, userIds)
	if err != nil {
		return nil, err
	}
	tagList := cs.TagList{}
	for _, tag := range tags {
		tagFormated := tag.Format()
		for _, user := range users {
			if user.ID == tagFormated.UserID {
				tagFormated.User = user
			}
		}
		tagList = append(tagList, tagFormated)
	}
	// 填充话题follow信息
	if len(userTopicsMap) > 0 {
		for _, tag := range tagList {
			if info, exist := userTopicsMap[tag.ID]; exist {
				tag.IsFollowing, tag.IsTop, tag.IsPin = 1, info.IsTop, info.IsPin
			}
		}
	}
	return tagList, nil
}

func (s *topicSrv) TagsByKeyword(keyword string) (res cs.TagInfoList, err error) {
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

func (s *topicSrvA) UpsertTags(userId int64, tags []string) (_ cs.TagInfoList, err error) {
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

func (s *topicSrvA) DecrTagsById(ids []int64) (err error) {
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

func (s *topicSrvA) ListTags(typ cs.TagType, offset, limit int) (res cs.TagList, err error) {
	conditions := &ms.ConditionsT{}
	switch typ {
	case cs.TagTypeHot:
		// 热门标签
		conditions = &ms.ConditionsT{
			"ORDER": "quote_num DESC",
		}
	case cs.TagTypeNew:
		// 最新标签
		conditions = &ms.ConditionsT{
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

func (s *topicSrvA) TagsByKeyword(keyword string) (res cs.TagInfoList, err error) {
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

func (s *topicSrv) FollowTopic(userId int64, topicId int64) (err error) {
	return s.db.Create(&dbr.TopicUser{
		UserID:  userId,
		TopicID: topicId,
		IsTop:   0,
	}).Error
}

func (s *topicSrv) UnfollowTopic(userId int64, topicId int64) error {
	return s.db.Exec("DELETE FROM "+s.tnTopicUser+" WHERE user_id=? AND topic_id=?", userId, topicId).Error
}

func (s *topicSrv) StickTopic(userId int64, topicId int64) (status int8, err error) {
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

func (s *topicSrv) PinTopic(userId int64, topicId int64) (status int8, err error) {
	db := s.db.Begin()
	defer db.Rollback()

	m := &dbr.TopicUser{}
	err = db.Model(m).
		Where("user_id=? and topic_id=?", userId, topicId).
		UpdateColumn("is_pin", gorm.Expr("1-is_pin")).Error
	if err != nil {
		return
	}
	status = -1
	err = db.Model(m).Where("user_id=? and topic_id=?", userId, topicId).Select("is_pin").Scan(&status).Error
	if err != nil {
		return
	}
	if status < 0 {
		return -1, errors.New("topic not exist")
	}

	db.Commit()
	return
}
