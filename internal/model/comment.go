package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	*Model
	PostID int64  `json:"post_id"`
	UserID int64  `json:"user_id"`
	IP     string `json:"ip"`
	IPLoc  string `json:"ip_loc"`
}

type CommentFormated struct {
	ID         int64                   `json:"id"`
	PostID     int64                   `json:"post_id"`
	UserID     int64                   `json:"user_id"`
	User       *UserFormated           `json:"user"`
	Contents   []*CommentContent       `json:"contents"`
	Replies    []*CommentReplyFormated `json:"replies"`
	IPLoc      string                  `json:"ip_loc"`
	CreatedOn  int64                   `json:"created_on"`
	ModifiedOn int64                   `json:"modified_on"`
}

func (c *Comment) Format() *CommentFormated {
	if c.Model == nil {
		return &CommentFormated{}
	}
	return &CommentFormated{
		ID:         c.Model.ID,
		PostID:     c.PostID,
		UserID:     c.UserID,
		User:       &UserFormated{},
		Contents:   []*CommentContent{},
		Replies:    []*CommentReplyFormated{},
		IPLoc:      c.IPLoc,
		CreatedOn:  c.CreatedOn,
		ModifiedOn: c.ModifiedOn,
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
	return db.Model(&Comment{}).Where("id = ? AND is_del = ?", c.Model.ID, 0).Updates(map[string]interface{}{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}
