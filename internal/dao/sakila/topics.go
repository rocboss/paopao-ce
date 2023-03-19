// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

var (
	_ core.TopicService = (*topicSrv)(nil)
)

type topicSrv struct {
	*sqlxSrv
	stmtNewestTags     *sqlx.Stmt
	stmtHotTags        *sqlx.Stmt
	stmtTagsByKeywordA *sqlx.Stmt
	stmtTagsByKeywordB *sqlx.Stmt
	stmtInsertTag      *sqlx.Stmt
	sqlTagsByIdA       string
	sqlTagsByIdB       string
	sqlDecrTagsById    string
	sqlTagsForIncr     string
	sqlIncrTagsById    string
}

func (s *topicSrv) UpsertTags(userId int64, tags []string) (res cs.TagInfoList, xerr error) {
	if len(tags) == 0 {
		return nil, nil
	}
	xerr = s.with(func(tx *sqlx.Tx) error {
		var upTags cs.TagInfoList
		if err := s.inSelect(tx, &upTags, s.sqlTagsForIncr, tags); err != nil {
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
			if _, err := s.inExec(tx, s.sqlIncrTagsById, now, ids); err != nil {
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
		var newTags cs.TagInfoList
		if err := s.inSelect(tx, &newTags, s.sqlTagsByIdB, ids); err != nil {
			return err
		}
		res = append(res, newTags...)
		return nil
	})
	return
}

func (s *topicSrv) DecrTagsById(ids []int64) error {
	return s.with(func(tx *sqlx.Tx) error {
		var ids []int64
		err := s.inSelect(tx, &ids, s.sqlTagsByIdA, ids)
		if err != nil {
			return err
		}
		_, err = s.inExec(tx, s.sqlDecrTagsById, time.Now().Unix(), ids)
		return err
	})
}

func (s *topicSrv) ListTags(typ cs.TagType, limit int, offset int) (res cs.TagList, err error) {
	switch typ {
	case cs.TagTypeHot:
		err = s.stmtHotTags.Select(&res, limit, offset)
	case cs.TagTypeNew:
		err = s.stmtNewestTags.Select(&res, limit, offset)
	}
	return
}

func (s *topicSrv) TagsByKeyword(keyword string) (res cs.TagInfoList, err error) {
	keyword = "%" + strings.Trim(keyword, " ") + "%"
	if keyword == "%%" {
		err = s.stmtTagsByKeywordA.Select(&res)
	} else {
		err = s.stmtTagsByKeywordB.Select(&res)
	}
	return
}

func newTopicService(db *sqlx.DB) core.TopicService {
	return &topicSrv{
		sqlxSrv:            newSqlxSrv(db),
		stmtNewestTags:     c(`SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.id DESC LIMIT ? OFFSET ?`),
		stmtHotTags:        c(`SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.quote_num DESC LIMIT ? OFFSET ?`),
		stmtTagsByKeywordA: c(`SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6`),
		stmtTagsByKeywordB: c(`SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6`),
		stmtInsertTag:      c(`INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1)`),
		sqlTagsByIdA:       t(`SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0`),
		sqlTagsByIdB:       t(`SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?)`),
		sqlDecrTagsById:    t(`UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?)`),
		sqlTagsForIncr:     t(`SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?)`),
		sqlIncrTagsById:    t(`UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?)`),
	}
}
