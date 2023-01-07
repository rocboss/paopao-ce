// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

func createTag(db *gorm.DB, tag *dbr.Tag) (*dbr.Tag, error) {
	t, err := tag.Get(db)
	if err != nil {
		tag.QuoteNum = 1
		return tag.Create(db)
	}

	// 更新
	t.QuoteNum++
	err = t.Update(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func deleteTag(db *gorm.DB, tag *dbr.Tag) error {
	tag, err := tag.Get(db)
	if err != nil {
		return err
	}
	tag.QuoteNum--
	return tag.Update(db)
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
