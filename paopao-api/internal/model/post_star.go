package model

import (
	"time"

	"gorm.io/gorm"
)

type PostStar struct {
	*Model
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
}

func (p *PostStar) Get(db *gorm.DB) (*PostStar, error) {
	var star PostStar
	if p.Model != nil && p.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", p.ID, 0)
	}
	if p.PostID > 0 {
		db = db.Where("post_id = ?", p.PostID)
	}
	if p.UserID > 0 {
		db = db.Where("user_id = ?", p.UserID)
	}

	err := db.First(&star).Error
	if err != nil {
		return &star, err
	}

	return &star, nil
}

func (p *PostStar) Create(db *gorm.DB) (*PostStar, error) {
	err := db.Create(&p).Error

	return p, err
}

func (p *PostStar) Delete(db *gorm.DB) error {
	return db.Model(&PostStar{}).Where("id = ? AND is_del = ?", p.Model.ID, 0).Updates(map[string]interface{}{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}

func (p *PostStar) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*PostStar, error) {
	var stars []*PostStar
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if p.UserID > 0 {
		db = db.Where("user_id = ?", p.UserID)
	}

	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&stars).Error; err != nil {
		return nil, err
	}

	return stars, nil
}

func (p *PostStar) Count(db *gorm.DB, conditions *ConditionsT) (int64, error) {
	var count int64
	if p.PostID > 0 {
		db = db.Where("post_id = ?", p.PostID)
	}
	if p.UserID > 0 {
		db = db.Where("user_id = ?", p.UserID)
	}
	for k, v := range *conditions {
		if k != "ORDER" {
			db = db.Where(k, v)
		}
	}
	if err := db.Model(p).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
