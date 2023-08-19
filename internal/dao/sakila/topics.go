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
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.TopicService  = (*topicSrvA)(nil)
	_ core.TopicServantA = (*topicSrvA)(nil)
)

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

func (s *topicSrvA) GetHotTags(userId int64, limit int, offset int) (cs.TagList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *topicSrvA) GetNewestTags(userId int64, limit int, offset int) (cs.TagList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *topicSrvA) GetFollowTags(userId int64, limit int, offset int) (cs.TagList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *topicSrvA) FollowTopic(userId int64, topicId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *topicSrvA) UnfollowTopic(userId int64, topicId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *topicSrvA) StickTopic(userId int64, topicId int64) (int8, error) {
	// TODO
	return 0, debug.ErrNotImplemented
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
