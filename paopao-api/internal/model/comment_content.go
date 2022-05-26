package model

import "gorm.io/gorm"

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
