// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

func (s *sqlxSrv) deleteTags(tx *sqlx.Tx, tags []string) (err error) {
	var tagInfos []cs.TagInfo
	if err = tx.Stmtx(s.y.TagsFromNames).Select(&tagInfos, tags); err != nil {
		return
	}
	now := time.Now().Unix()
	for _, tag := range tagInfos {
		tag.QuoteNum--
		if _, err = tx.Stmtx(s.y.UpdateTagQuote).Exec(tag.QuoteNum, now, tag.ID); err != nil {
			return
		}
	}
	return
}

func (s *sqlxSrv) createTags(tx *sqlx.Tx, userId int64, tags []string) (res cs.TagInfoList, xerr error) {
	if len(tags) == 0 {
		return nil, nil
	}
	xerr = s.db.Withx(func(tx *sqlx.Tx) error {
		var upTags cs.TagInfoList
		if err := tx.InSelect(&upTags, s.y.TagsForIncr, tags); err != nil {
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
			if _, err := tx.InExec(s.y.IncrTagsById, now, ids); err != nil {
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
			res, err := s.y.InsertTag.Exec(userId, tag, now, now)
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
		if err := tx.InSelect(&newTags, s.y.TagsByIdB, ids); err != nil {
			return err
		}
		res = append(res, newTags...)
		return nil
	})
	return
}
