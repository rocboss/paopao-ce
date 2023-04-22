// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"

	"github.com/rocboss/paopao-ce/pkg/types"
	"gorm.io/gorm"
)

type Comment struct {
	*Model
	PostID          int64  `json:"post_id"`
	UserID          int64  `json:"user_id"`
	IP              string `json:"ip"`
	IPLoc           string `json:"ip_loc"`
	ThumbsUpCount   int32  `json:"thumbs_up_count"`
	ThumbsDownCount int32  `json:"-"`
}

type CommentFormated struct {
	ID            int64                   `json:"id"`
	PostID        int64                   `json:"post_id"`
	UserID        int64                   `json:"user_id"`
	User          *UserFormated           `json:"user"`
	Contents      []*CommentContent       `json:"contents"`
	Replies       []*CommentReplyFormated `json:"replies"`
	IPLoc         string                  `json:"ip_loc"`
	ThumbsUpCount int32                   `json:"thumbs_up_count"`
	IsThumbsUp    int8                    `json:"is_thumbs_up"`
	IsThumbsDown  int8                    `json:"is_thumbs_down"`
	CreatedOn     int64                   `json:"created_on"`
	ModifiedOn    int64                   `json:"modified_on"`
}

func (c *Comment) Format() *CommentFormated {
	if c.Model == nil {
		return &CommentFormated{}
	}
	return &CommentFormated{
		ID:            c.Model.ID,
		PostID:        c.PostID,
		UserID:        c.UserID,
		User:          &UserFormated{},
		Contents:      []*CommentContent{},
		Replies:       []*CommentReplyFormated{},
		IPLoc:         c.IPLoc,
		ThumbsUpCount: c.ThumbsUpCount,
		IsThumbsUp:    types.No,
		IsThumbsDown:  types.No,
		CreatedOn:     c.CreatedOn,
		ModifiedOn:    c.ModifiedOn,
	}
}

func (c *Comment) Get(db *gorm.DB) (*Comment, error) {
	var comment Comment
	if c.Model != nil && c.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", c.ID, 0)
	} else {
		return nil, gorm.ErrRecordNotFound
	}

	err := db.First(&comment).Error
	if err != nil {
		return &comment, err
	}

	return &comment, nil
}

func (c *Comment) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*Comment, error) {
	var comments []*Comment
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if c.PostID > 0 {
		db = db.Where("id = ?", c.PostID)
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

func (c *Comment) Count(db *gorm.DB, conditions *ConditionsT) (int64, error) {
	var count int64
	if c.PostID > 0 {
		db = db.Where("post_id = ?", c.PostID)
	}
	for k, v := range *conditions {
		if k != "ORDER" {
			db = db.Where(k, v)
		}
	}
	if err := db.Model(c).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (c *Comment) Create(db *gorm.DB) (*Comment, error) {
	err := db.Create(&c).Error

	return c, err
}

func (c *Comment) Delete(db *gorm.DB) error {
	return db.Model(c).Where("id = ?", c.Model.ID).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}

func (c *Comment) CommentIdsByPostId(db *gorm.DB, postId int64) (ids []int64, err error) {
	err = db.Model(c).Where("post_id = ?", postId).Select("id").Find(&ids).Error
	return
}

func (c *Comment) DeleteByPostId(db *gorm.DB, postId int64) error {
	return db.Model(c).Where("post_id = ?", postId).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}
