// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"

	"gorm.io/gorm"
)

type CommentContent struct {
	*Model
	CommentID int64        `json:"comment_id"`
	UserID    int64        `json:"user_id"`
	Content   string       `json:"content"`
	Type      PostContentT `json:"type"`
	Sort      int64        `json:"sort"`
}

func (c *CommentContent) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*CommentContent, error) {
	var comments []*CommentContent
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if c.CommentID > 0 {
		db = db.Where("id = ?", c.CommentID)
	}

	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *CommentContent) Create(db *gorm.DB) (*CommentContent, error) {
	err := db.Create(&c).Error

	return c, err
}

func (c *CommentContent) MediaContentsByCommentId(db *gorm.DB, commentIds []int64) (contents []string, err error) {
	err = db.Model(c).Where("comment_id IN ? AND type = ?", commentIds, ContentTypeImage).Select("content").Find(&contents).Error
	return
}

func (c *CommentContent) DeleteByCommentIds(db *gorm.DB, commentIds []int64) error {
	return db.Model(c).Where("comment_id IN ?", commentIds).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}
