package model

import (
	"time"

	"gorm.io/gorm"
)

type CommentReply struct {
	*Model
	CommentID int64  `json:"comment_id"`
	UserID    int64  `json:"user_id"`
	AtUserID  int64  `json:"at_user_id"`
	Content   string `json:"content"`
	IP        string `json:"ip"`
	IPLoc     string `json:"ip_loc"`
}

type CommentReplyFormated struct {
	ID         int64         `json:"id"`
	CommentID  int64         `json:"comment_id"`
	UserID     int64         `json:"user_id"`
	User       *UserFormated `json:"user"`
	AtUserID   int64         `json:"at_user_id"`
	AtUser     *UserFormated `json:"at_user"`
	Content    string        `json:"content"`
	IPLoc      string        `json:"ip_loc"`
	CreatedOn  int64         `json:"created_on"`
	ModifiedOn int64         `json:"modified_on"`
}

func (c *CommentReply) Format() *CommentReplyFormated {
	if c.Model == nil {
		return &CommentReplyFormated{}
	}

	return &CommentReplyFormated{
		ID:         c.ID,
		CommentID:  c.CommentID,
		UserID:     c.UserID,
		User:       &UserFormated{},
		AtUserID:   c.AtUserID,
		AtUser:     &UserFormated{},
		Content:    c.Content,
		IPLoc:      c.IPLoc,
		CreatedOn:  c.CreatedOn,
		ModifiedOn: c.ModifiedOn,
	}
}

func (c *CommentReply) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*CommentReply, error) {
	var comments []*CommentReply
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

func (c *CommentReply) Create(db *gorm.DB) (*CommentReply, error) {
	err := db.Create(&c).Error

	return c, err
}

func (c *CommentReply) Get(db *gorm.DB) (*CommentReply, error) {
	var reply CommentReply
	if c.Model != nil && c.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", c.ID, 0)
	} else {
		return nil, gorm.ErrRecordNotFound
	}

	err := db.First(&reply).Error
	if err != nil {
		return &reply, err
	}

	return &reply, nil
}

func (c *CommentReply) Delete(db *gorm.DB) error {
	return db.Model(&CommentReply{}).Where("id = ? AND is_del = ?", c.Model.ID, 0).Updates(map[string]interface{}{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}
