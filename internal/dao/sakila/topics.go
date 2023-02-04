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
	stmtTagsByName     string
	stmtIncrTagsById   string
}

func (s *topicServant) UpsertTags(userId int64, tags []string) (res []*core.Tag, xerr error) {
	if len(tags) <= 0 {
		return nil, nil
	}
	xerr = s.with(func(tx *sqlx.Tx) error {
		query, args, err := in(s.db, s.stmtTagsByName, tags)
		var ts []*core.Tag
		if err = tx.Select(&ts, query, args...); err != nil {
			return err
		}
		var upTags []string
		if len(ts) > 0 {
			var ids []int64
			for _, t := range ts {
				ids = append(ids, t.ID)
				upTags = append(upTags, t.Tag)
				t.QuoteNum++
				// prepare remain tags just delete updated tag
				for i, name := range tags {
					if name == t.Tag {
						size := len(tags)
						tags[i] = tags[size-1]
						tags = tags[:size-1]
						break
					}
				}
			}
			if query, args, err = in(s.db, s.stmtIncrTagsById, ids); err != nil {
				return err
			}
			if _, err = tx.Exec(query, args...); err != nil {
				return err
			}
			res = append(res, ts...)
		}
		// process remain tags
		if len(tags) == 0 {
			return nil
		}
		var ids []int64
		now := time.Now().Unix()
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
		query, args, err = in(s.db, s.stmtTagsByIdB, ids)
		if err != nil {
			return err
		}

		if err = tx.Select(&newTags, query, args...); err != nil {
			return err
		}
		res = append(res, newTags...)
		return nil
	})
	return
}

func (s *topicServant) DecrTagsById(ids []int64) error {
	return s.with(func(tx *sqlx.Tx) error {
		query, args, err := in(s.db, s.stmtTagsByIdA, ids)
		if err != nil {
			return err
		}
		var ids []int64
		if err = tx.Select(&ids, query, args...); err != nil {
			return err
		}
		query, args, err = in(s.db, s.stmtDecrTagsById, time.Now().Unix(), ids)
		_, err = tx.Exec(query, args...)
		return err
	})
}

func (s *topicServant) GetTags(category core.TagCategory, offset int, limit int) (res []*core.Tag, err error) {
	switch category {
	case core.TagCategoryHot:
		err = s.stmtHotTags.Select(&res, offset, limit)
	case core.TagCategoryNew:
		err = s.stmtHotTags.Select(&res, offset, limit)
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
		stmtTagsByIdA:      r(`SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num >= 0`),
		stmtTagsByIdB:      r(`SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?)`),
		stmtDecrTagsById:   r(`UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?)`),
		stmtTagsByName:     r(`SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?) AND is_del = 0 AND quote_num >= 0`),
		stmtIncrTagsById:   r(`UPDATE @tag SET quote_num=quote_num+1 WHERE id IN (?)`),
	}
}
