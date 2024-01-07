// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pga"
)

var (
	_ core.TopicService  = (*pgaTopicSrvA)(nil)
	_ core.TopicServantA = (*pgaTopicSrvA)(nil)
)

type pgaTopicSrvA struct {
	*topicSrvA
	p *pga.TopicA
}

func (s *pgaTopicSrvA) UpsertTags(userId int64, tags []string) (res cs.TagInfoList, xerr error) {
	if len(tags) == 0 {
		return nil, nil
	}
	xerr = s.db.Withx(func(tx *sqlx.Tx) error {
		var upTags cs.TagInfoList
		err := tx.InSelect(&upTags, s.q.TagsForIncr, tags)
		if err != nil {
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
		var (
			id  int64
			ids []int64
		)
		for _, tag := range tags {
			if err = s.p.InsertTag.Get(&id, userId, tag, now, now); err != nil {
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

func (s *pgaTopicSrvA) StickTopic(userId int64, topicId int64) (res int8, err error) {
	err = s.p.StickTopic.Get(&res, time.Now().Unix(), userId, topicId)
	return
}

func (s *pgaTopicSrvA) PinTopic(userId int64, topicId int64) (res int8, err error) {
	err = s.p.PinTopic.Get(&res, time.Now().Unix(), userId, topicId)
	return
}
