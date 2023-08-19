// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"strings"
	"time"

	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
)

var (
	_ core.TopicService  = (*topicSrvA)(nil)
	_ core.TopicServantA = (*topicSrvA)(nil)
)

type topicInfo struct {
	TopicId int64
	IsTop   int8
}

type topicSrvA struct {
	*sqlxSrv
	q *cc.TopicA
}

func (s *topicSrvA) UpsertTags(userId int64, tags []string) (res cs.TagInfoList, xerr error) {
	if len(tags) == 0 {
		return nil, nil
	}
	xerr = s.db.Withx(func(tx *sqlx.Tx) error {
		var upTags cs.TagInfoList
		if err := tx.InSelect(&upTags, s.q.TagsForIncr, tags); err != nil {
			return err
		}
		now := time.Now().Unix()
		if len(upTags) > 0 {
			var ids []int64
			for _, t := range upTags {
				ids = append(ids, t.ID)
				t.QuoteNum++
				// prepare remain tags just delete updated tag
				// notice ensure tags slice is distinct elements
				for i, name := range tags {
					if name == t.Tag {
						lastIdx := len(tags) - 1
						tags[i] = tags[lastIdx]
						tags = tags[:lastIdx]
						break
					}
				}
			}
			if _, err := tx.InExec(s.q.IncrTagsById, now, ids); err != nil {
				return err
			}
			res = append(res, upTags...)
		}
		// process remain tags if tags is not empty
		if len(tags) == 0 {
			return nil
		}
		var ids []int64
		for _, tag := range tags {
			res, err := s.q.InsertTag.Exec(userId, tag, now, now)
			if err != nil {
				return err
			}
			id, err := res.LastInsertId()
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
		var newTags cs.TagInfoList
		if err := tx.InSelect(&newTags, s.q.TagsByIdB, ids); err != nil {
			return err
		}
		res = append(res, newTags...)
		return nil
	})
	return
}

func (s *topicSrvA) DecrTagsById(ids []int64) error {
	return s.db.Withx(func(tx *sqlx.Tx) error {
		var ids []int64
		err := tx.InSelect(&ids, s.q.TagsByIdA, ids)
		if err != nil {
			return err
		}
		_, err = tx.InExec(s.q.DecrTagsById, time.Now().Unix(), ids)
		return err
	})
}

func (s *topicSrvA) ListTags(typ cs.TagType, limit int, offset int) (res cs.TagList, err error) {
	switch typ {
	case cs.TagTypeHot:
		err = s.q.HotTags.Select(&res, limit, offset)
	case cs.TagTypeNew:
		err = s.q.NewestTags.Select(&res, limit, offset)
	}
	return
}

func (s *topicSrvA) TagsByKeyword(keyword string) (res cs.TagInfoList, err error) {
	keyword = "%" + strings.Trim(keyword, " ") + "%"
	if keyword == "%%" {
		err = s.q.TagsByKeywordA.Select(&res)
	} else {
		err = s.q.TagsByKeywordB.Select(&res)
	}
	return
}

func (s *topicSrvA) GetHotTags(userId int64, limit int, offset int) (res cs.TagList, err error) {
	if err = s.q.HotTags.Select(&res, limit, offset); err != nil {
		return
	}
	return s.tagsFormatA(userId, res)
}

func (s *topicSrvA) GetNewestTags(userId int64, limit int, offset int) (res cs.TagList, err error) {
	err = s.q.NewestTags.Select(&res, limit, offset)
	return
}

func (s *topicSrvA) GetFollowTags(userId int64, limit int, offset int) (res cs.TagList, err error) {
	if err = s.q.FollowTags.Select(&res, userId, limit, offset); err != nil {
		return
	}
	return s.tagsFormatA(userId, res)
}

func (s *topicSrvA) FollowTopic(userId int64, topicId int64) (err error) {
	exist := false
	if err = s.q.ExistTopicUser.Get(&exist, userId, topicId); err == nil {
		_, err = s.q.FollowTopic.Exec(userId, topicId, time.Now().Unix())
	}
	return
}

func (s *topicSrvA) UnfollowTopic(userId int64, topicId int64) error {
	_, err := s.q.UnfollowTopic.Exec(userId, topicId, time.Now().Unix())
	return err
}

func (s *topicSrvA) StickTopic(userId int64, topicId int64) (res int8, err error) {
	s.db.Withx(func(tx *sqlx.Tx) error {
		_, err = tx.Stmtx(s.q.StickTopic).Exec(time.Now().Unix(), userId, topicId)
		if err == nil {
			err = tx.Stmtx(s.q.TopicIsTop).Get(&res, userId, topicId)
		}
		return err
	})
	return
}

func (s *topicSrvA) tagsFormatA(userId int64, tags cs.TagList) (cs.TagList, error) {
	// 获取创建者User IDs
	tagIds := []int64{}
	for _, tag := range tags {
		tagIds = append(tagIds, tag.ID)
	}
	// 填充话题follow信息
	if userId > -1 {
		userTopics := []*topicInfo{}
		err := s.db.InSelect(&userTopics, s.q.TopicInfos, userId, tagIds)
		if err != nil {
			return nil, err
		}
		userTopicsMap := make(map[int64]*topicInfo, len(userTopics))
		for _, info := range userTopics {
			userTopicsMap[info.TopicId] = info
		}
		for _, tag := range tags {
			if info, exist := userTopicsMap[tag.ID]; exist {
				tag.IsFollowing, tag.IsTop = 1, info.IsTop
			}
		}
	}
	return tags, nil
}

func newTopicService(db *sqlx.DB) core.TopicService {
	return &topicSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTopicA),
	}
}

func newTopicServantA(db *sqlx.DB) core.TopicServantA {
	return &topicSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTopicA),
	}
}
