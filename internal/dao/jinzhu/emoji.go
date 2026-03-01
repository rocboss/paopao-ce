// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"errors"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.EmojiService = (*emojiSrv)(nil)
)

const (
	_emoji_       = "p_emoji"
	_user_emoji_  = "p_user_emoji"
)

type emojiSrv struct {
	db *gorm.DB
}

func newEmojiService(db *gorm.DB) core.EmojiService {
	return &emojiSrv{
		db: db,
	}
}

// UploadEmoji 上传表情包
func (s *emojiSrv) UploadEmoji(userID int64, name, url string, width, height int, size int64, emojiType string) (*dbr.Emoji, error) {
	emoji := &dbr.Emoji{
		Model: &dbr.Model{},
		UserID: userID,
		Name: name,
		URL: url,
		Width: width,
		Height: height,
		Size: size,
		Type: emojiType,
		CollectionCount: 0,
	}

	return emoji.Create(s.db)
}

// GetEmojiList 获取表情包列表
func (s *emojiSrv) GetEmojiList(offset, limit int) ([]*dbr.Emoji, error) {
	emoji := &dbr.Emoji{}
	conditions := &dbr.ConditionsT{
		"ORDER": "collection_count DESC, created_on DESC",
	}

	return emoji.List(s.db, conditions, offset, limit)
}

// GetUserEmojiList 获取用户上传的表情包列表
func (s *emojiSrv) GetUserEmojiList(userID int64, offset, limit int) ([]*dbr.Emoji, error) {
	emoji := &dbr.Emoji{}
	conditions := &dbr.ConditionsT{
		"user_id = ?": userID,
		"ORDER": "created_on DESC",
	}

	return emoji.List(s.db, conditions, offset, limit)
}

// GetUserCollectedEmojiList 获取用户收藏的表情包列表
func (s *emojiSrv) GetUserCollectedEmojiList(userID int64, offset, limit int) ([]*dbr.Emoji, error) {
	userEmoji := &dbr.UserEmoji{}
	return userEmoji.ListUserEmojis(s.db, userID, offset, limit)
}

// CollectEmoji 收藏表情包
func (s *emojiSrv) CollectEmoji(userID, emojiID int64) error {
	// 检查表情包是否存在
	emoji := &dbr.Emoji{
		Model: &dbr.Model{
			ID: emojiID,
		},
	}
	_, err := emoji.Get(s.db)
	if err != nil {
		return err
	}

	// 检查是否已经收藏
	userEmoji := &dbr.UserEmoji{
		UserID: userID,
		EmojiID: emojiID,
	}
	_, err = userEmoji.GetUserEmoji(s.db)
	if err == nil {
		return errors.New("emoji already collected")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 创建收藏记录
	userEmoji = &dbr.UserEmoji{
		Model: &dbr.Model{},
		UserID: userID,
		EmojiID: emojiID,
	}
	_, err = userEmoji.Create(s.db)
	if err != nil {
		return err
	}

	// 增加表情包收藏数
	return emoji.IncrementCollectionCount(s.db)
}

// UncollectEmoji 取消收藏表情包
func (s *emojiSrv) UncollectEmoji(userID, emojiID int64) error {
	// 检查是否已经收藏
	userEmoji := &dbr.UserEmoji{
		UserID: userID,
		EmojiID: emojiID,
	}
	_, err := userEmoji.GetUserEmoji(s.db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("emoji not collected")
		}
		return err
	}

	// 删除收藏记录
	err = userEmoji.Delete(s.db)
	if err != nil {
		return err
	}

	// 减少表情包收藏数
	emoji := &dbr.Emoji{
		Model: &dbr.Model{
			ID: emojiID,
		},
	}
	return emoji.DecrementCollectionCount(s.db)
}

// IsEmojiCollected 检查表情包是否被用户收藏
func (s *emojiSrv) IsEmojiCollected(userID, emojiID int64) (bool, error) {
	userEmoji := &dbr.UserEmoji{
		UserID: userID,
		EmojiID: emojiID,
	}
	_, err := userEmoji.GetUserEmoji(s.db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// GetEmojiByID 获取表情包详情
func (s *emojiSrv) GetEmojiByID(emojiID int64) (*dbr.Emoji, error) {
	emoji := &dbr.Emoji{
		Model: &dbr.Model{
			ID: emojiID,
		},
	}
	return emoji.Get(s.db)
}

// DeleteEmoji 删除表情包
func (s *emojiSrv) DeleteEmoji(userID, emojiID int64) error {
	// 检查表情包是否存在且属于该用户
	emoji := &dbr.Emoji{
		Model: &dbr.Model{
			ID: emojiID,
		},
	}
	emoji, err := emoji.Get(s.db)
	if err != nil {
		return err
	}

	if emoji.UserID != userID {
		return errors.New("permission denied")
	}

	// 删除表情包
	return emoji.Delete(s.db)
}
