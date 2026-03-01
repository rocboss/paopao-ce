// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"gorm.io/gorm"
)

// Emoji 表情包模型
type Emoji struct {
	*Model
	UserID          int64  `json:"user_id"`
	Name            string `json:"name"`
	URL             string `json:"url"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	Size            int64  `json:"size"`
	Type            string `json:"type"`
	CollectionCount int64  `json:"collection_count"`
}

// UserEmoji 用户表情包收藏模型
type UserEmoji struct {
	*Model
	UserID  int64 `json:"user_id"`
	EmojiID int64 `json:"emoji_id"`
}

// EmojiFormated 格式化的表情包信息
type EmojiFormated struct {
	ID              int64  `json:"id"`
	UserID          int64  `json:"user_id"`
	Name            string `json:"name"`
	URL             string `json:"url"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	Size            int64  `json:"size"`
	Type            string `json:"type"`
	CollectionCount int64  `json:"collection_count"`
	IsCollected     bool   `json:"is_collected"`
	CreatedOn       int64  `json:"created_on"`
}

// Format 格式化表情包信息
func (e *Emoji) Format() *EmojiFormated {
	if e.Model != nil {
		return &EmojiFormated{
			ID:              e.ID,
			UserID:          e.UserID,
			Name:            e.Name,
			URL:             e.URL,
			Width:           e.Width,
			Height:          e.Height,
			Size:            e.Size,
			Type:            e.Type,
			CollectionCount: e.CollectionCount,
			IsCollected:     false,
			CreatedOn:       e.CreatedOn,
		}
	}
	return nil
}

// Get 获取单个表情包
func (e *Emoji) Get(db *gorm.DB) (*Emoji, error) {
	var emoji Emoji
	if e.Model != nil && e.Model.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", e.Model.ID, 0)
	} else {
		return nil, gorm.ErrRecordNotFound
	}

	err := db.First(&emoji).Error
	if err != nil {
		return nil, err
	}

	return &emoji, nil
}

// List 获取表情包列表
func (e *Emoji) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*Emoji, error) {
	var emojis []*Emoji
	var err error

	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}

	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&emojis).Error; err != nil {
		return nil, err
	}

	return emojis, nil
}

// Create 创建表情包
func (e *Emoji) Create(db *gorm.DB) (*Emoji, error) {
	err := db.Create(&e).Error
	return e, err
}

// Update 更新表情包
func (e *Emoji) Update(db *gorm.DB) error {
	return db.Model(&Emoji{}).Where("id = ? AND is_del = ?", e.Model.ID, 0).Save(e).Error
}

// Delete 删除表情包
func (e *Emoji) Delete(db *gorm.DB) error {
	return db.Delete(&e).Error
}

// IncrementCollectionCount 增加收藏数
func (e *Emoji) IncrementCollectionCount(db *gorm.DB) error {
	return db.Model(&Emoji{}).Where("id = ? AND is_del = ?", e.Model.ID, 0).Update("collection_count", gorm.Expr("collection_count + ?", 1)).Error
}

// DecrementCollectionCount 减少收藏数
func (e *Emoji) DecrementCollectionCount(db *gorm.DB) error {
	return db.Model(&Emoji{}).Where("id = ? AND is_del = ? AND collection_count > 0", e.Model.ID, 0).Update("collection_count", gorm.Expr("collection_count - ?", 1)).Error
}

// GetUserEmoji 获取用户的表情包收藏
func (ue *UserEmoji) GetUserEmoji(db *gorm.DB) (*UserEmoji, error) {
	var userEmoji UserEmoji
	err := db.Where("user_id = ? AND emoji_id = ? AND is_del = ?", ue.UserID, ue.EmojiID, 0).First(&userEmoji).Error
	if err != nil {
		return nil, err
	}
	return &userEmoji, nil
}

// ListUserEmojis 获取用户的表情包收藏列表
func (ue *UserEmoji) ListUserEmojis(db *gorm.DB, userID int64, offset, limit int) ([]*Emoji, error) {
	var emojis []*Emoji

	query := db.Table("p_emoji").Select("p_emoji.*").Joins("JOIN p_user_emoji ON p_emoji.id = p_user_emoji.emoji_id").Where("p_user_emoji.user_id = ? AND p_user_emoji.is_del = ? AND p_emoji.is_del = ?", userID, 0, 0)

	if offset >= 0 && limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	err := query.Order("p_user_emoji.created_on DESC").Find(&emojis).Error
	if err != nil {
		return nil, err
	}

	return emojis, nil
}

// Create 创建用户表情包收藏
func (ue *UserEmoji) Create(db *gorm.DB) (*UserEmoji, error) {
	err := db.Create(&ue).Error
	return ue, err
}

// Delete 删除用户表情包收藏
func (ue *UserEmoji) Delete(db *gorm.DB) error {
	return db.Where("user_id = ? AND emoji_id = ? AND is_del = ?", ue.UserID, ue.EmojiID, 0).Delete(&UserEmoji{}).Error
}
