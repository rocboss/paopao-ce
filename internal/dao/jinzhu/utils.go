// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

func createTags(db *gorm.DB, userId int64, tags []string) (res cs.TagInfoList, err error) {
	for _, name := range tags {
		tag := &dbr.Tag{Tag: name}
		if tag, err = tag.Get(db); err == nil {
			// 更新
			tag.QuoteNum++
			if err = tag.Update(db); err != nil {
				return
			}
		} else {
			tag.UserID, tag.QuoteNum = userId, 1
			if tag, err = (&dbr.Tag{
				UserID:   userId,
				QuoteNum: 1,
			}).Create(db); err != nil {
				return
			}
		}
		res = append(res, &cs.TagInfo{
			ID:       tag.ID,
			UserID:   tag.UserID,
			Tag:      tag.Tag,
			QuoteNum: tag.QuoteNum,
		})
	}
	return
}

func decrTagsByIds(db *gorm.DB, ids []int64) (err error) {
	for _, id := range ids {
		tag := &dbr.Tag{Model: &dbr.Model{ID: id}}
		if tag, err = tag.Get(db); err == nil {
			tag.QuoteNum--
			if err = tag.Update(db); err != nil {
				return
			}
		} else {
			continue
		}
	}
	return nil
}

func deleteTags(db *gorm.DB, tags []string) error {
	allTags, err := (&dbr.Tag{}).TagsFrom(db, tags)
	if err != nil {
		return err
	}
	for _, tag := range allTags {
		tag.QuoteNum--
		if tag.QuoteNum < 0 {
			tag.QuoteNum = 0
		}
		// 宽松处理错误, 尽可能更新tag记录， 只记录最后一次错误
		if e := tag.Update(db); e != nil {
			err = e
		}
	}
	return err
}

// 根据IDs获取用户列表
func getUsersByIDs(db *gorm.DB, ids []int64) ([]*dbr.User, error) {
	user := &dbr.User{}
	return user.List(db, &dbr.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}
