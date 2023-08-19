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

func (s *sqlxSrv) createTags(tx *sqlx.Tx, userId int64, tags []string) (res cs.TagInfoList, err error) {
	// for _, name := range tags {
	// 	tag := &dbr.Tag{Tag: name}
	// 	if tag, err = tag.Get(db); err == nil {
	// 		// 更新
	// 		tag.QuoteNum++
	// 		if err = tag.Update(db); err != nil {
	// 			return
	// 		}
	// 	} else {
	// 		if tag, err = (&dbr.Tag{
	// 			UserID:   userId,
	// 			QuoteNum: 1,
	// 			Tag:      name,
	// 		}).Create(db); err != nil {
	// 			return
	// 		}
	// 	}
	// 	res = append(res, &cs.TagInfo{
	// 		ID:       tag.ID,
	// 		UserID:   tag.UserID,
	// 		Tag:      tag.Tag,
	// 		QuoteNum: tag.QuoteNum,
	// 	})
	// }
	return
}
