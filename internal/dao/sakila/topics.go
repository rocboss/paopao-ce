// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"strings"
	"time"

	"github.com/alimy/yesql"
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/sirupsen/logrus"
)

var (
	_ core.TopicService = (*topicSrv)(nil)
)

type topicSrv struct {
	*sqlxSrv
	Scope              yesql.Scope `yesql:"topic"`
	StmtNewestTags     *sqlx.Stmt  `yesql:"newest_tags"`
	StmtHotTags        *sqlx.Stmt  `yesql:"hot_tags"`
	StmtTagsByKeywordA *sqlx.Stmt  `yesql:"tags_by_keyword_a"`
	StmtTagsByKeywordB *sqlx.Stmt  `yesql:"tags_by_keyword_b"`
	StmtInsertTag      *sqlx.Stmt  `yesql:"insert_tag"`
	SqlTagsByIdA       string      `yesql:"tags_by_id_a"`
	SqlTagsByIdB       string      `yesql:"tags_by_id_b"`
	SqlDecrTagsById    string      `yesql:"decr_tags_by_id"`
	SqlTagsForIncr     string      `yesql:"tags_for_incr"`
	SqlIncrTagsById    string      `yesql:"incr_tags_by_id"`
}

func (s *topicSrv) UpsertTags(userId int64, tags []string) (res cs.TagInfoList, xerr error) {
	if len(tags) == 0 {
		return nil, nil
	}
	xerr = s.with(func(tx *sqlx.Tx) error {
		var upTags cs.TagInfoList
		if err := s.inSelect(tx, &upTags, s.SqlTagsForIncr, tags); err != nil {
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
			if _, err := s.inExec(tx, s.SqlIncrTagsById, now, ids); err != nil {
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
			res, err := s.StmtInsertTag.Exec(userId, tag, now, now)
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
		if err := s.inSelect(tx, &newTags, s.SqlTagsByIdB, ids); err != nil {
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
		err := s.inSelect(tx, &ids, s.SqlTagsByIdA, ids)
		if err != nil {
			return err
		}
		_, err = s.inExec(tx, s.SqlDecrTagsById, time.Now().Unix(), ids)
		return err
	})
}

func (s *topicSrv) ListTags(typ cs.TagType, limit int, offset int) (res cs.TagList, err error) {
	switch typ {
	case cs.TagTypeHot:
		err = s.StmtHotTags.Select(&res, limit, offset)
	case cs.TagTypeNew:
		err = s.StmtNewestTags.Select(&res, limit, offset)
	}
	return
}

func (s *topicSrv) TagsByKeyword(keyword string) (res cs.TagInfoList, err error) {
	keyword = "%" + strings.Trim(keyword, " ") + "%"
	if keyword == "%%" {
		err = s.StmtTagsByKeywordA.Select(&res)
	} else {
		err = s.StmtTagsByKeywordB.Select(&res)
	}
	return
}

func newTopicService(db *sqlx.DB, query yesql.SQLQuery) core.TopicService {
	obj := &topicSrv{
		sqlxSrv: newSqlxSrv(db),
	}
	if err := yesql.Scan(obj, query); err != nil {
		logrus.Fatal(err)
	}
	return obj
}
