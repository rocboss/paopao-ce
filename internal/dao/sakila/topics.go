// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
)

var (
	_ core.TopicService = (*topicServant)(nil)
)

type topicServant struct {
	*sqlxServant
	stmtNewestTags     *sqlx.Stmt
	stmtHotTags        *sqlx.Stmt
	stmtTagsByKeywordA *sqlx.Stmt
	stmtTagsByKeywordB *sqlx.Stmt
	stmtInsertTag      *sqlx.Stmt
	stmtTagsByIdA      string
	stmtTagsByIdB      string
	stmtDecrTagsById   string
	stmtTagsForIncr    string
	stmtIncrTagsById   string
}

func (s *topicServant) UpsertTags(userId int64, tags []string) (res []*core.Tag, xerr error) {
	if len(tags) == 0 {
		return nil, nil
	}
	xerr = s.with(func(tx *sqlx.Tx) error {
		var upTags []*core.Tag
		if err := s.inSelect(tx, &upTags, s.stmtTagsForIncr, tags); err != nil {
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
			if _, err := s.inExec(tx, s.stmtIncrTagsById, now, ids); err != nil {
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
			res, err := s.stmtInsertTag.Exec(userId, tag, now, now)
			if err != nil {
				return err
			}
			id, err := res.LastInsertId()
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
		var newTags []*core.Tag
		if err := s.inSelect(tx, &newTags, s.stmtTagsByIdB, ids); err != nil {
			return err
		}
		res = append(res, newTags...)
		return nil
	})
	return
}

func (s *topicServant) DecrTagsById(ids []int64) error {
	return s.with(func(tx *sqlx.Tx) error {
		var ids []int64
		err := s.inSelect(tx, &ids, s.stmtTagsByIdA, ids)
		if err != nil {
			return err
		}
		_, err = s.inExec(tx, s.stmtDecrTagsById, time.Now().Unix(), ids)
		return err
	})
}

func (s *topicServant) GetTags(category core.TagCategory, offset int, limit int) (res []*core.Tag, err error) {
	switch category {
	case core.TagCategoryHot:
		err = s.stmtHotTags.Select(&res, offset, limit)
	case core.TagCategoryNew:
		err = s.stmtNewestTags.Select(&res, offset, limit)
	}
	return
}

func (s *topicServant) GetTagsByKeyword(keyword string) (res []*core.Tag, err error) {
	keyword = "%" + strings.Trim(keyword, " ") + "%"
	if keyword == "%%" {
		err = s.stmtTagsByKeywordA.Select(&res)
	} else {
		err = s.stmtTagsByKeywordB.Select(&res)
	}
	return
}

func newTopicService(db *sqlx.DB) core.TopicService {
	return &topicServant{
		sqlxServant:        newSqlxServant(db),
		stmtNewestTags:     c(`SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND quote_num > 0 ORDER BY id DESC OFFSET ? LIMIT ?`),
		stmtHotTags:        c(`SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND quote_num > 0 ORDER BY quote_num DESC OFFSET ? LIMIT ?`),
		stmtTagsByKeywordA: c(`SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC OFFSET 0 LIMIT 6`),
		stmtTagsByKeywordB: c(`SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC OFFSET 0 LIMIT 6`),
		stmtInsertTag:      c(`INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1)`),
		stmtTagsByIdA:      r(`SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0`),
		stmtTagsByIdB:      r(`SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?)`),
		stmtDecrTagsById:   r(`UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?)`),
		stmtTagsForIncr:    r(`SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?)`),
		stmtIncrTagsById:   r(`UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?)`),
	}
}
